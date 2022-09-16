package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
)

func main() {

	usage := `A simple command-line tool to generate random integers, floats and strings.

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

`

	opts, _ := docopt.ParseArgs(usage, nil, "https://github.com/alasdairmorris/random v1.1.0")

	switch {
	case opts["int"]:
		length, _ := opts.Int("--length")
		signed, _ := opts.Bool("--signed")
		str, err := intCommand(length, signed)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(str)
	case opts["float"]:
		whole, _ := opts.Int("--whole")
		fractional, _ := opts.Int("--fractional")
		signed, _ := opts.Bool("--signed")
		str, err := floatCommand(whole, fractional, signed)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(str)
	case opts["str"]:
		length, _ := opts.Int("--length")
		charset, _ := opts.String("--charset")
		uppercase, _ := opts.Bool("--uppercase")
		lowercase, _ := opts.Bool("--lowercase")
		str, err := strCommand(length, charset, uppercase, lowercase)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(str)
	}

}

func intCommand(length int, signed bool) (string, error) {

	if length < 1 {
		return "", errors.New("Invalid value for 'length' parameter")
	}

	sign := ""
	if signed {
		sign = randomString([]rune("+-"), 1)
		if sign == "+" {
			sign = ""
		}
	}

	p1 := randomString([]rune("123456789"), 1)
	p2 := randomString([]rune("0123456789"), length-1)

	return sign + p1 + p2, nil
}

func floatCommand(whole int, fractional int, signed bool) (string, error) {
	if whole < 1 {
		return "", errors.New("Invalid value for 'whole' parameter")
	}

	if fractional < 1 {
		return "", errors.New("Invalid value for 'fractional' parameter")
	}

	sign := ""
	if signed {
		sign = randomString([]rune("+-"), 1)
		if sign == "+" {
			sign = ""
		}
	}

	p1 := randomString([]rune("123456789"), 1)
	p2 := randomString([]rune("0123456789"), whole-1)
	p3 := randomString([]rune("0123456789"), fractional)

	return sign + p1 + p2 + "." + p3, nil
}

func strCommand(length int, charset string, uppercase bool, lowercase bool) (string, error) {

	if length < 1 {
		return "", errors.New("Invalid value for 'length' parameter")
	}

	var runes = []rune(charset)
	retval := randomString(runes, length)

	if uppercase {
		retval = strings.ToUpper(retval)
	} else if lowercase {
		retval = strings.ToLower(retval)
	}

	return retval, nil
}
