package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

type runner interface {
	init([]string) error
	run() error
	name() string
	usage() error
}

func printUsage() {
	str := `A simple command-line tool to generate random integers, floats and strings.

USAGE
    %s <command> [flags]

COMMANDS
    int:    Output a random integer
    float:  Output a random float
    str:    Output a random string

    help <command>
            Output usage info for the given command

`

	fmt.Printf(str, filepath.Base(os.Args[0]))
}

func Root(args []string) error {
	if len(args) < 1 {
		printUsage()
		return nil
	}

	cmds := []runner{
		newIntCommand(),
		newStrCommand(),
		newFloatCommand(),
	}

	if os.Args[1] == "help" && len(args) > 1 {
		subcommand := os.Args[2]
		for _, cmd := range cmds {
			if cmd.name() == subcommand {
				cmd.init(os.Args[2:])
				return cmd.usage()
			}
		}
	} else {
		subcommand := os.Args[1]
		for _, cmd := range cmds {
			if cmd.name() == subcommand {
				cmd.init(os.Args[2:])
				return cmd.run()
			}
		}
	}

	printUsage()
	return nil
}
