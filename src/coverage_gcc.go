//go:build ccover

package opennox

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/noxworld-dev/opennox/v1/internal/ccover"
)

func init() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-ch
		ccover.Dump()
		if sig == os.Interrupt {
			os.Exit(130)
		}
		os.Exit(143)
	}()
}

func flushCoverage() {
	ccover.Dump()
}
