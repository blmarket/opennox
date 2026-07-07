package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/noxworld-dev/opennox/v1"
)

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	if len(args) == 0 {
		args = []string{"opennox-audio"}
	}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	if err := flags.Parse(args[1:]); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if flags.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "unexpected arguments: %v\n", flags.Args())
		return 1
	}
	if err := opennox.RunAudioTest(); err != nil && err != flag.ErrHelp {
		if code, ok := err.(opennox.ErrExit); ok {
			return int(code)
		} else {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	}
	return 0
}
