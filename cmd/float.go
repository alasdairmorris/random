package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alasdairmorris/random/tools"
)

type floatCommand struct {
	fs *flag.FlagSet

	whole      int
	fractional int
	signed     bool
}

func newFloatCommand() *floatCommand {
	gc := &floatCommand{
		fs: flag.NewFlagSet("float", flag.ExitOnError),
	}

	gc.fs.IntVar(&gc.whole, "whole", 8, "")
	gc.fs.IntVar(&gc.whole, "w", 8, "")

	gc.fs.IntVar(&gc.fractional, "fractional", 4, "")
	gc.fs.IntVar(&gc.fractional, "f", 4, "")

	gc.fs.BoolVar(&gc.signed, "signed", false, "")
	gc.fs.BoolVar(&gc.signed, "s", false, "")

	gc.fs.Usage = func() {
		str := `usage: %s float [<options>]

    -w, --whole <length>       specify number of digits for whole part (i.e. before d.p.)
    -f, --fractional <length>  specify number of digits for fractional part (i.e. after d.p.)
    -s, --signed               specify number should be signed (default: unsigned)

`

		fmt.Printf(str, filepath.Base(os.Args[0]))
	}

	return gc
}

func (f *floatCommand) name() string {
	return f.fs.Name()
}

func (f *floatCommand) init(args []string) error {
	return f.fs.Parse(args)
}

func (f *floatCommand) usage() error {
	f.fs.Usage()
	return nil
}

func (f *floatCommand) run() error {

	if f.whole < 1 {
		return fmt.Errorf("Invalid whole length: %d", f.whole)
	}

	if f.fractional < 1 {
		return fmt.Errorf("Invalid fractional length: %d", f.fractional)
	}

	sign := ""
	if f.signed {
		sign = tools.RandomString([]rune("+-"), 1)
		if sign == "+" {
			sign = ""
		}
	}

	p1 := tools.RandomString([]rune("123456789"), 1)
	p2 := tools.RandomString([]rune("0123456789"), f.whole-1)
	p3 := tools.RandomString([]rune("0123456789"), f.fractional)

	fmt.Println(sign + p1 + p2 + "." + p3)

	return nil
}
