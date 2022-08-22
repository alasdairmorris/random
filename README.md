# random

A command-line tool for generating random numbers and strings.

## Installation

If you have Go installed, you can do:

```
go install github.com/alasdairmorris/random@latest
```

Otherwise, just download the relevant binary from [the release page](https://github.com/alasdairmorris/random/releases).

## Usage

```
> random
A simple command-line tool to generate random integers, floats and strings.

USAGE
    random <command> [flags]

COMMANDS
    int:    Output a random integer
    float:  Output a random float
    str:    Output a random string

    help <command>
            Output usage info for the given command
```

## Examples

```
> random int
8679141860
```

```
> random int -l 6
588895
```

```
> random str
qduXgGUVPZ
```

```
> random str --charset "ABCDEF0123456789" -l 6   ## generate a random colour!
A9D0EC
```

```
> random help int
usage: random int [<options>]

    -l, --length <length>   specify number of digits
    -s, --signed            specify number should be signed (default: unsigned)

```
