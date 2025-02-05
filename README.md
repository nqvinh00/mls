# mls

## Overview

`mls` is a command-line tool like `ls`, designed to list and manage files with enhanced features such as colorized output, file type icons, and sorting options. It is cross-platform and supports both Unix and Windows systems.

## Features

- Colorized output with customizable color codes.
- File type icons for quick identification.
- Sorting options by size, date, extension, and type.
- Ability to show hidden files.
- Option to display files as a list or in columns.
- Link resolution for symlinked files.

## Installation

To install `mls`, clone the repository and build the binary using Go:

```sh
git clone https://github.com/nqvinh00/mls.git
cd mls
go build
```

Or:

```
go install github.com/nqvinh00/mls
```

Ensure that the `mls` binary is in your PATH for easy access from the command line.

## Usage

Run `mls` with various options to tailor the output to your needs:

```sh
mls [options] [path...]
```

### Options

- `-a, --all`: Show all files including hidden ones.
- `-l, --list`: Show files in a detailed list format.
- `-t, --tree`: Display files as a tree structure.
- `-s, --sort [type]`: Sort files by extension (`x`), type (`t`), size (`s`), or date (`d`).
- `-C, --no-color`: Disable color output.
- `-I, --no-icon`: Disable icon output.
- `-L, --no-link`: Disable link resolution for symlinks.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your improvements.
