#  Welcome to **Eradicate (ERA)**! 

Ever find yourself cringing when a debugging or test command sneaks its way into your production code? 😱 Don’t worry, ERA's got your back!

Imagine ERA as your digital "magic wand" 🪄 for code cleanup. It’s like having a superpower that sweeps through your codebase, banishing those pesky output command that slipped through the cracks and ensuring your production environment remains pristine. 🚀

## 🛠️ **How It Works:**
Just run ERA, and watch it work its magic! It’ll zap away any unwanted output command from your code , leaving your codebase cleaner than ever. 🌟

## 🚀 **Why ERA Rocks:**
- **Instant Cleanup:** ERA swiftly removes those unintended commands with a flick of your wrist (or a click of your mouse).
- **Stress-Free:** No more second-guessing or frantic searches. ERA makes sure your production environment stays squeaky clean.
- **Developer’s Best Friend:** ERA is designed to solve the problem every developer faces—keeping your code spotless and professional.

So, why wait? Give ERA a spin and make your codebase sparkle! ✨

# Eradicate(ERA) Installation Guide

This guide will help you install the latest version of Eradicate(ERA) on your system. The installation script supports Linux, macOS, and Windows.

## Prerequisites

- Bash shell (for Linux and macOS)
- PowerShell (for Windows)
- `curl` command-line tool
- Internet connection
- `sudo` access (for Linux and macOS)

## Installation Steps

1. Download the installation script:
   ```
   curl -O https://raw.githubusercontent.com/wadmit/era/master/scripts/install.sh
   ```

2. Make the script executable (Linux and macOS only):
   ```
   chmod +x install.sh
   ```

3. Run the installation script:
   - On Linux and macOS:
     ```
     ./install.sh
     ```
   - On Windows (using PowerShell):
     ```
     bash install.sh
     ```

4. The script will automatically detect your operating system, fetch the latest version of ERA, download the appropriate binary, and install it in the correct location.

5. After successful installation, you should be able to run ERA by typing `era` in your terminal or command prompt.


## Usage Guide

This section provides a step-by-step guide on how to use Eradicate (era) effectively in your project.

### 1. Initializing Your Project

To set up Eradicate in your project, run the following command:

```bash
erd init
```

- **What It Does**: This command initializes Eradicate in your project workspace. It creates a default configuration file named `erd.yaml` in the current directory.
- **After Initialization**: You will see the `erd.yaml` file generated. This file contains default settings that you can customize according to your project's needs.

### 2. Cleaning Your Project

To remove console output and other specified patterns from your project, use the command:

```bash
erd clean
```

- **Functionality**: This command searches for the `erd.yaml` file in your project directory and applies the cleaning rules defined in it. It will scan through your codebase and remove any matching patterns, such as console logs or other specified outputs.
- **Default Configuration**: 
  - If the `erd.yaml` file is not found, Eradicate will automatically use a default configuration to perform the cleaning.
  - The default settings will target common console output functions for various programming languages.

### 3. Removing Specific Files or Folders

If you need to clean up specific files or folders, you can use:

```bash
erd remove -f <filelocation/filename> -d <directory>
```

- **Parameters**:
  - `-f <filelocation/filename>`: This flag specifies the path to a single file that you want to clean. It verifies the existence of the file before proceeding with the removal of specified patterns.
  - `-d <directory>`: This flag allows you to specify a directory from which you want to remove outputs. Eradicate will check all files within the specified directory and apply the cleaning rules.
  
- **Functionality**:
  - The command first checks if the `-f` flag is present. If it is, the tool verifies the provided file location and removes specified patterns (like console logs) from that file.
  - If the `-d` flag is provided, it will clean all supported files in the specified directory. Files that are not supported or specified in the ignore list will be skipped.

### Example Scenarios

#### Example 1: Initializing Eradicate

After navigating to your project directory, run:

```bash
erd init
```

This creates an `erd.yaml` file. You can open this file to adjust the cleaning rules according to your preferences.

#### Example 2: Cleaning the Entire Project

To clean your project based on the rules in `erd.yaml`, simply run:

```bash
erd clean
```

This command will traverse your project files and remove unwanted console outputs based on your configuration.

#### Example 3: Cleaning a Specific File

If you want to remove console outputs only from a specific file, you can do:

```bash
erd remove -f src/main.js
```

This will check the file `main.js` in the `src` directory and remove specified outputs.

#### Example 4: Cleaning a Specific Directory

To clean all files in a directory, use:

```bash
erd remove -d src/components
```

This command will apply the cleaning rules to all relevant files in the `components` directory.

### Additional Notes

- Ensure that you have the necessary permissions to read and write files in the directories you specify.
- Review your `erd.yaml` configuration before running the `clean` command to make sure it meets your requirements.
- For further details on the configuration options, refer to the [Configuration Section](#configuration).


## Configuration

The configuration for Eradicate (era) is defined in the `erd.yaml` file generated during the initialization process. This configuration controls how the tool operates, what files to ignore, and where to store reports.

### Configuration Structure

The configuration is structured as follows:

### Configuration Fields

- **Root**: Specifies the root directory for the project. It defaults to the current directory (`"."`). This is where Eradicate will start searching for files to process.

- **ReportPath**: Indicates the directory where the reports will be stored. By default, it is set to `"era-reports"`. After cleaning or processing files, the reports detailing the changes made will be saved in this location.

- **IgnoreKeyword**: An array of keywords that, if found in the code, will instruct Eradicate to ignore the associated lines or sections. By default, it includes `["erd:ignore", "erd:ignoreAll"]`. You can customize this list to add any keywords relevant to your project that you want to ignore.
  Example
  ```python
  print("hello world") #erd:ignore
  ```
  Above code will ignore that print statement from being removed.It will go similar with other languages like Javascript, Java, Ruby,Php etc.

- **IgnoreFileExtensions**: A list of file extensions that Eradicate will skip during the cleaning process. This is useful for excluding binary files or archives that should not be processed. The default extensions are:
  - `.exe`
  - `.dll`
  - `.so`
  - `.dylib`
  - `.zip`
  - `.tar`
  - `.gz`
  - `.rar`

- **IgnoreDirs**: An array of directory names that will be ignored by Eradicate. Common directories like `node_modules`, version control directories (`.git`, `.svn`, etc.), and IDE configurations are included by default. You can customize this list to add any directories that should not be processed.

- **IgnoreFiles**: A list of specific files to ignore during processing. This is an empty array by default, allowing you to add specific files that you want to exclude from cleaning.

- **ListenType**: Defines the type of listening method for Eradicate. Currently, it supports only the `"command"` type, which means the tool will execute commands based on user input from the command line. This is where you can specify if you want to add additional types in future versions.

### Example `erd.yaml` Configuration

Here's an example of how your `erd.yaml` file might look after customization:

```yaml
root: "."
reportPath: "era-reports"
ignoreKeyword:
  - "erd:ignore"
  - "erd:ignoreAll"
ignoreFileExtensions:
  - ".exe"
  - ".dll"
  - ".so"
ignoreDirs:
  - "node_modules"
  - ".git"
ignoreFiles: []
listenType: "command"
```

### Notes

- Ensure that your `erd.yaml` file is saved in the root directory of your project for Eradicate to function properly.
- Customize the configuration fields according to your project's needs to optimize the cleaning process.


## Troubleshooting

If you encounter any issues during installation, please check the following:

- Ensure you have an active internet connection.
- Verify that you have the necessary permissions to install software on your system.
- For Linux and macOS users, make sure you have `sudo` access and write permissions to `/usr/local/bin/`.
- For Windows users, ensure you're running PowerShell as an administrator.

If problems persist, please open an issue on our GitHub repository with details about the error you're experiencing.

## Uninstallation

To uninstall ERA:

- On Linux and macOS:
  ```
  sudo rm /usr/local/bin/era
  ```
- On Windows:
  ```
  Remove-Item $env:USERPROFILE\era.exe
  ```

## Updating ERA

To update ERA to the latest version, simply run the installation script again. It will automatically fetch and install the most recent release.

## Support

If you need help or have any questions, please open an issue on our GitHub repository or contact our support team.

Thank you for using ERA!
