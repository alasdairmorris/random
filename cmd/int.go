package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alasdairmorris/random/tools"
)

type intCommand struct {
	fs *flag.FlagSet

	length int
	signed bool
}

func newIntCommand() *intCommand {
	gc := &intCommand{
		fs: flag.NewFlagSet("int", flag.ExitOnError),
	}

	gc.fs.IntVar(&gc.length, "length", 10, "")
	gc.fs.IntVar(&gc.length, "l", 10, "")

	gc.fs.BoolVar(&gc.signed, "signed", false, "")
	gc.fs.BoolVar(&gc.signed, "s", false, "")

	gc.fs.Usage = func() {
		str := `usage: %s int [<options>]

    -l, --length <length>   specify number of digits
    -s, --signed            specify number should be signed (default: unsigned)

`

		fmt.Printf(str, filepath.Base(os.Args[0]))
	}

	return gc
}

func (i *intCommand) name() string {
	return i.fs.Name()
}

func (i *intCommand) init(args []string) error {
	return i.fs.Parse(args)
}

func (i *intCommand) usage() error {
	i.fs.Usage()
	return nil
}

func (i *intCommand) run() error {
	sign := ""
	if i.signed {
		sign = tools.RandomString([]rune("+-"), 1)
		if sign == "+" {
			sign = ""
		}
	}

	p1 := tools.RandomString([]rune("123456789"), 1)
	p2 := tools.RandomString([]rune("0123456789"), i.length-1)

	fmt.Println(sign + p1 + p2)

	return nil
}
