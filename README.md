# Would You Rather - Terminal Game

## Overview

Would You Rather is a command-line game that presents users with difficult hypothetical choices between two options. The game tracks and displays statistics about how other players might have answered the same question, adding a competitive and social element to the experience. and it's the first terminal application that supports would you rather 

## Features

- 200+ built-in "Would You Rather" questions
- Persistent question database stored in `~/.questions/questions.txt`
- Simple command-line interface
- Randomized question selection
- Answer statistics simulation
- Cross-platform support (macOS, Linux, FreeBSD)
- Easy customization by editing the question file

## Installation

### macOS
```bash
curl -fsSL https://github.com/funterminal/wouldyourather/blob/c4bad09a188309e414bb2a4ef302f489cc577a42/wouldyourather-darwin-amd64 -o wouldyourather
chmod +x wouldyourather
sudo mv wouldyourather /usr/local/bin/
```

### Linux
```bash
curl -fsSL https://github.com/funterminal/wouldyourather/raw/c4bad09a188309e414bb2a4ef302f489cc577a42/wouldyourather-linux-amd64 -o wouldyourather
chmod +x wouldyourather
sudo mv wouldyourather /usr/local/bin/
```

### FreeBSD
```bash
curl -fsSL https://github.com/funterminal/wouldyourather/raw/c4bad09a188309e414bb2a4ef302f489cc577a42/wouldyourather-freebsd-amd64 -o wouldyourather
chmod +x wouldyourather
sudo mv wouldyourather /usr/local/bin/
```

### Manual Installation (from source)
```bash
curl -fsSL https://raw.githubusercontent.com/funterminal/wouldyourather/refs/heads/main/wouldyourather.go -o wouldyourather.go
go build wouldyourather.go
sudo mv wouldyourather /usr/local/bin/
```

## Usage

1. First, initialize the question database:
```bash
wouldyourather init
```

2. Then, start asking questions:
```bash
wouldyourather ask
```

3. When presented with a question, type `1` or `2` to select your preferred option.

## Customization

To add your own questions, edit the file at `~/.questions/questions.txt`. Each question should be on its own line in the format:
```
Option 1 or Option 2
```

The game will automatically include any valid questions from this file along with the built-in questions.

## Examples

Sample game session:
```
$ wouldyourather ask
Would you rather:
1) Be able to fly
2) Be invisible
> 1
You chose option 1.
Stats: 1) 63%  2) 37%

Would you rather:
1) Live without music
2) Live without television
> 2
You chose option 2.
Stats: 1) 28%  2) 72%
```

## Technical Details

- Questions are stored in a simple text file format
- The game uses Go's built-in randomization functions
- Statistics are simulated for demonstration purposes
- The code is written in pure Go with no external dependencies

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome. Please fork the repository and submit a pull request with your changes.

## Support

For issues or feature requests, please open an issue on the GitHub repository.
