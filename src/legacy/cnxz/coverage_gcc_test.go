//go:build ccover

package cnxz

import (
	"os"
	"testing"

	"github.com/noxworld-dev/opennox/v1/internal/ccover"
)

func TestMain(m *testing.M) {
	code := m.Run()
	ccover.Dump()
	os.Exit(code)
}
