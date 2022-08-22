package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alasdairmorris/random/tools"
)

type strCommand struct {
	fs *flag.FlagSet

	length    int
	uppercase bool
	lowercase bool
	charset   string
}

func newStrCommand() *strCommand {
	gc := &strCommand{
		fs: flag.NewFlagSet("str", flag.ExitOnError),
	}

	gc.fs.IntVar(&gc.length, "length", 10, "")
	gc.fs.IntVar(&gc.length, "l", 10, "")

	gc.fs.BoolVar(&gc.uppercase, "uppercase", false, "")
	gc.fs.BoolVar(&gc.uppercase, "U", false, "")

	gc.fs.BoolVar(&gc.lowercase, "lowercase", false, "")
	gc.fs.BoolVar(&gc.lowercase, "L", false, "")

	gc.fs.StringVar(&gc.charset, "charset", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	gc.fs.StringVar(&gc.charset, "c", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

	gc.fs.Usage = func() {
		str := `usage: %s str [<options>]

    -l, --length <length>   specify number of letters
    -L, --lowercase         specify string should be all lowercase (default: mixed case)
    -U, --uppercase         specify string should be all uppercase (default: mixed case)
    -c, --charset           specify characters to use when generating string

`

		fmt.Printf(str, filepath.Base(os.Args[0]))
	}

	return gc
}

func (s *strCommand) name() string {
	return s.fs.Name()
}

func (s *strCommand) init(args []string) error {
	return s.fs.Parse(args)
}

func (s *strCommand) usage() error {
	s.fs.Usage()
	return nil
}

func (s *strCommand) run() error {
	var charset = []rune(s.charset)

	retval := tools.RandomString(charset, s.length)

	if s.uppercase {
		retval = strings.ToUpper(retval)
	} else if s.lowercase {
		retval = strings.ToLower(retval)
	}

	fmt.Println(retval)

	return nil
}
