# random

A command-line tool for generating random numbers and strings.

## Installation

`random` will run on most Linux and Mac OS X systems.

To install it, just find the appropriate one-liner below - based on the destination O/S and architecture - and copy-paste it into your terminal.

Feel free to change the install dir - `$HOME/bin` in the examples below - to be something more appropriate for your needs.

### Linux (32-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/random/releases/latest/download/random-linux-386.tar.gz | tar -zxf - -C $HOME/bin
```

### Linux (64-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/random/releases/latest/download/random-linux-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Intel)

```
curl -s -L -o - https://github.com/alasdairmorris/random/releases/latest/download/random-darwin-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o - https://github.com/alasdairmorris/random/releases/latest/download/random-darwin-arm64.tar.gz | tar -zxf - -C $HOME/bin
```

### Build From Source

If you have Go installed and would prefer to build the app yourself, you can do:

```
go install github.com/alasdairmorris/random@latest
```

## Usage

```
A simple command-line tool to generate random integers, floats and strings.

Usage:
  random str [-l <length>] [-c <charset>] [-U | -L]
  random int [-l <length>] [-s]
  random float [-w <whole>] [-f <fractional>] [-s]
  random -h | --help
  random --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.

Shared Options:
  -l, --length=<length>  Length [default: 10]
  -s, --signed           Result can be positive or negative [default: positive only]

String Options:
  -c, --charset=<chars>  Character set [default: abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ]
  -U, --uppercase        Convert to uppercase
  -L, --lowercase        Convert to lowercase

Float Options:
  -w, --whole=<whole>    specify number of digits for the whole part (i.e. before the decimal point) [default: 8]
  -f, --fractional=<fr>  specify number of digits for the fractional part (i.e. after the decimal point) [default: 4]
```

## Examples

```
$ random int
8679141860
```

```
$ random int -l 6
588895
```

```
$ random str
qduXgGUVPZ
```

```
$ random str --charset "ABCDEF0123456789" -l 6   ## generate a random colour!
A9D0EC
```
