#  Welcome to **Eradicate (ERA)**! 

Ever find yourself cringing when a debugging or test command sneaks its way into your production code? 😱 Don’t worry, ERA's got your back!

Imagine ERA as your digital "magic wand" 🪄 for code cleanup. It’s like having a superpower that sweeps through your codebase, banishing those pesky command lines that slipped through the cracks and ensuring your production environment remains pristine. 🚀

## 🛠️ **How It Works:**
Just run ERA, and watch it work its magic! It’ll zap away any unwanted command lines from your output, leaving your code cleaner than ever. 🌟

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