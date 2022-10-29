# random

A command-line tool for generating random numbers and strings.

## Installation

### Linux (32-bit)

To install the binary in e.g. $HOME/bin

```
curl -L -o - https://github.com/alasdairmorris/random/releases/latest/download/random-linux-386.tar.gz | tar -zxvf - -C $HOME/bin
```

### Linux (64-bit)

To install the binary in e.g. $HOME/bin

```
curl -L -o - https://github.com/alasdairmorris/random/releases/latest/download/random-linux-amd64.tar.gz | tar -zxvf - -C $HOME/bin
```

### Mac OS X (Intel)

To install the binary in e.g. $HOME/bin

```
curl -L -o - https://github.com/alasdairmorris/random/releases/latest/download/random-darwin-amd64.tar.gz | tar -zxvf - -C $HOME/bin
```

### Mac OS X (Apple Silicon)

To install the binary in e.g. $HOME/bin

```
curl -L -o - https://github.com/alasdairmorris/random/releases/latest/download/random-darwin-arm64.tar.gz | tar -zxvf - -C $HOME/bin
```

### Build From Source

If you have Go installed and prefer to build the app yourself, you can do:

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
