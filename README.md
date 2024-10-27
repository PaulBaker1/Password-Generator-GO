# Password Generator GO

A Go-based desktop application for secure and customizable password generation, built using the [Fyne](https://fyne.io/) GUI framework. This app offers a range of password generation options and customization features to meet user needs for strong and memorable passwords.

---

## Table of Contents

- [Demo](#demo)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Customization](#customization)
- [Contributing](#contributing)
- [License](#license)
- [Contact Information](#contact-information)

---

## Demo

*Placeholder for demo GIF.*  
Replace this with an actual demo of the application:

![Demo of the app](path/to/demo.gif)

---

## Features

- **Customizable Password Length**: Set the desired length of the password.
- **Character Options**: Toggle inclusion of symbols, numbers, uppercase letters, and lowercase letters.
- **Enhanced Security Options**:
  - **No Similar Characters**: Exclude similar-looking characters (e.g., `i`, `l`, `1`, `O`) to improve readability.
  - **No Duplicate Characters**: Ensure each character in the password is unique.
  - **No Sequential Characters**: Prevent sequences like `abc` or `123` for added security.
- **Editable Password Display**: Allows users to modify the generated password before copying.
- **Auto-Generate on Start**: Automatically generates a password when the application is launched.

---

## Installation

### Prerequisites

- **Go**: Ensure [Go is installed](https://golang.org/doc/install) (v1.16+ recommended).
- **Fyne**: GUI library used in this project. Installation is included in the steps below.

### Steps

1. **Clone the repository**:
   ```bash
   git clone https://github.com/username/password-generator-gui.git
   cd password-generator-gui
   ```

2. **Install dependencies**:
   ```bash
   go get fyne.io/fyne/v2
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```

---

## Usage

### Interface Overview

- **Password Length Slider**: Set the desired password length from the provided range.
- **Character Options**: Select which types of characters to include:
  - **Symbols**: `!@#$%^&*()-_=+[]{}|;:,.<>/?`
  - **Numbers**: `0123456789`
  - **Uppercase**: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
  - **Lowercase**: `abcdefghijklmnopqrstuvwxyz`
- **Additional Options**:
  - **Begin with Letter**: Ensures the password starts with an alphabetic character.
  - **No Similar Characters**: Excludes similar characters (`iIl1Lo0O`).
  - **No Duplicate Characters**: Ensures each character is unique.
  - **No Sequential Characters**: Prevents sequences like `abc` or `123`.

### Steps to Generate a Password

1. Launch the app by running `go run main.go`.
2. Adjust the password length and character options as needed.
3. Click **Generate** to create a password.
4. Edit the generated password directly in the output field if needed.
5. Copy the password as needed.

---

## Customization

To customize the application’s behavior or default settings, modify the `PasswordOptions` struct in `model/password.go`.

### Changing Character Sets

1. **Symbols**: To change the symbols used in passwords, update the `buildCharacterSet` function in `model/password.go`.
2. **Default Settings**: Adjust fields like `DefaultLength`, `IncludeSymbols`, `IncludeNumbers`, etc., within the `PasswordOptions` struct.

### Adding New Features

If you’d like to add additional features, consider modifying the `GeneratePassword` function in `model/password.go`. Add options to the `PasswordOptions` struct as necessary, following the structure of existing options.

---

## Contributing

We welcome contributions! To contribute to the project, please follow these steps:

1. **Fork the repository** on GitHub.
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/your-username/password-generator-gui.git
   ```
3. **Create a new branch** for your feature or bug fix:
   ```bash
   git checkout -b feature-or-bugfix-name
   ```
4. **Make your changes** and **commit** them with descriptive messages.
5. **Push to your fork**:
   ```bash
   git push origin feature-or-bugfix-name
   ```
6. **Submit a pull request** to the main repository for review.

Please ensure your code follows the project’s coding and documentation standards.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contact Information

For questions, feedback, or collaboration requests, please contact the project maintainer:

[GitHub Profile](https://github.com/PaulBaker1)

---

This `README.md` file provides a comprehensive guide to the project, making it easier for users and contributors to understand, install, and use the password generator. Remember to replace any placeholders like `your-username` or paths to demo files with actual content specific to your project.
