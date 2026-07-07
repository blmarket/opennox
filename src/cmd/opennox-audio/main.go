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
