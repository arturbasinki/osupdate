# osupdate

**osupdate** is a command-line tool written in Go for managing and automating Linux system updates. It automatically detects your Linux distribution and runs the appropriate package manager commands to keep your system up-to-date.

## Prerequisites

- Go 1.16 or higher (for building from source)
- sudo privileges on your system

## Installation

### Option 1: Download Pre-built Binary

1. Download the latest binary from the [Releases page](https://github.com/arturbasinki/osupdate/releases)
2. Extract the binary:
   ```bash
   tar -xzf osupdate-linux-amd64.tar.gz
   ```
3. Move it to a directory in your PATH:
   ```bash
   sudo mv osupdate /usr/local/bin/
   ```
4. Make it executable:
   ```bash
   sudo chmod +x /usr/local/bin/osupdate
   ```

### Option 2: Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/arturbasinki/osupdate.git
   cd osupdate
   ```
2. Build and install:
   ```bash
   go install
   ```
   or
   ```bash
   go build -o osupdate
   sudo mv osupdate /usr/local/bin/
   ```

## Usage

Simply run the command:
```bash
osupdate
```

The tool will:
1. Detect your Linux distribution
2. Update package lists
3. Upgrade all installed packages
4. Display progress information during the update process

## Supported Distributions

- Debian
- Ubuntu
- Arch Linux
- Manjaro
- openSUSE
- Fedora
- CentOS
- RHEL

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
