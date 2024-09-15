package transform

import (
	"github.com/wadmit/eradicate/internal/fileio"
	"github.com/wadmit/eradicate/internal/types"
	"github.com/wadmit/eradicate/internal/utils"
)

// Transformer defines the methods for transforming files.
type Transformer interface {
	Transform(root string, cfg *types.Config, fileTransformChan chan<- Transform)
	Detect(root string, cfg *types.Config) <-chan string
}

// Transform holds the file transformation data.
type Transform struct {
	FilePath     string
	Extension    string
	ContentLines []string
}

// FileTransformer is a concrete implementation of Transformer.
type FileTransformer struct{}

// Transform reads files and sends transformed data to the channel.
func (t *FileTransformer) Transform(root string, cfg *types.Config, fileTransformChan chan<- Transform) {
	defer close(fileTransformChan) // Ensure channel is closed when done

	fileChan := t.Detect(root, cfg)

	for filePath := range fileChan {
		transformed, err := TransformFile(filePath)
		if err != nil {
			continue
		}
		fileTransformChan <- transformed
	}
}

func TransformFile(filePath string) (Transform, error) {
	fileReader := fileio.NewFileReader(filePath)
	contentLines, err := fileReader.ReadLines()
	if err != nil {
		return Transform{}, err
	}

	transformed := Transform{
		FilePath:     filePath,
		Extension:    utils.GetFileExt(filePath),
		ContentLines: contentLines,
	}

	return transformed, nil
}

// Detect identifies files based on config and returns a channel of file paths.
func (t *FileTransformer) Detect(root string, cfg *types.Config) <-chan string {
	fileChan := make(chan string)
	go func() {
		defer close(fileChan) // Ensure channel is closed when done
		utils.FileWalker(root, fileChan, cfg)
	}()
	return fileChan
}
