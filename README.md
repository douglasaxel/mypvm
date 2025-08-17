# MyPVM - PHP Version Manager

MyPVM is a command-line tool designed to help developers manage multiple PHP versions on their system. It allows you to download, install, and switch between different PHP versions easily.

## Features

- List available PHP versions online
- List locally installed PHP versions
- Install specific PHP versions
- Switch between installed PHP versions
- Cross-platform support (Windows, Linux, macOS)

## Installation

### Prerequisites

- Go 1.18 or higher

### Building from Source

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/mypvm.git
   cd mypvm
   ```

2. Build the application:
   ```
   go build -o mypvm
   ```

3. Add the executable to your PATH or move it to a directory in your PATH.

## Usage

```
mypvm [command] [arguments]
```

### Available Commands

- `list` - Lists PHP versions available online
- `local` - Lists PHP versions installed locally
- `install [version]` - Installs a specific PHP version
- `remove [version]` - Removes a specific PHP version
- `use [version]` - Selects a specific PHP version

### Examples

List all available PHP versions:
```
mypvm list
```

Install a specific PHP version:
```
mypvm install 8.3.0
```

Switch to a specific PHP version:
```
mypvm use 8.3.0
```

## Roadmap

### Implemented Features
- ✅ List available PHP versions online
- ✅ List locally installed PHP versions
- ✅ Install specific PHP versions
- ✅ Switch between installed PHP versions
- ✅ Cross-platform support (Windows, Linux, macOS)
- ✅ Progress bar for downloads

### Planned Features
- ⬜ Remove specific PHP versions
- ⬜ Update installed PHP versions
- ⬜ Install PHP extensions
- ⬜ Configure PHP settings
- ⬜ Create isolated PHP environments for projects
- ⬜ Integration with common web servers (Apache, Nginx)
- ⬜ Support for PHP development tools (Composer, PHPUnit)
- ⬜ Interactive mode for version selection
- ⬜ Auto-update functionality

## How It Works

MyPVM downloads PHP binaries from the official PHP website and manages them in a dedicated directory (`~/.mypvm`). When you switch between versions, MyPVM creates symbolic links to the selected version.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the Creative Commons CC0 1.0 Universal License - see the [LICENSE](LICENSE) file for details.
