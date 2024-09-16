# ERA Installation Guide

This guide will help you install the latest version of ERA(Eradicate) on your system. The installation script supports Linux, macOS, and Windows.

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