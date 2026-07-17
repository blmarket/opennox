package legacy

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"
)

const clientDecodeCorpusChild = "NOX_CLIENT_DECODE_CORPUS_CHILD"
const clientDecodeCorpusOpcode = "NOX_CLIENT_DECODE_CORPUS_OPCODE"

func TestLegacyClientDecoderCorpus(t *testing.T) {
	if os.Getenv(clientDecodeCorpusChild) == "1" {
		op, err := strconv.Atoi(os.Getenv(clientDecodeCorpusOpcode))
		if err != nil || op < 0 || op > 255 {
			t.Fatalf("invalid opcode: %q", os.Getenv(clientDecodeCorpusOpcode))
		}
		Nox_client_isConnected = func() bool { return false }
		_ = C_clientDecodeZeroPacket(byte(op))
		return
	}

	// These packet kinds immediately enter client subsystems that are not
	// initialized in the isolated decoder process. The remaining corpus
	// reaches all stateless and disconnected-client paths safely.
	stateful := map[int]bool{
		86: true, 87: true, 89: true, 106: true, 147: true, 152: true,
		164: true, 169: true, 176: true, 195: true, 215: true, 226: true,
	}
	var failures []string
	for op := 0; op <= 255; op++ {
		if stateful[op] {
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestLegacyClientDecoderCorpus$")
		cmd.Env = append(os.Environ(),
			clientDecodeCorpusChild+"=1",
			clientDecodeCorpusOpcode+"="+strconv.Itoa(op),
		)
		out, err := cmd.CombinedOutput()
		cancel()
		if err != nil {
			failures = append(failures, fmt.Sprintf("%d: %v: %s", op, err, strings.TrimSpace(string(out))))
		}
	}
	if len(failures) != 0 {
		t.Fatalf("zero-filled decoder corpus failures:\n%s", strings.Join(failures, "\n"))
	}
}
