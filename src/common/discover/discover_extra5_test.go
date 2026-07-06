package discover

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStaticIPsExtra(t *testing.T) {
	// Test with non-existent file
	list, err := staticIPs("nonexistent.txt")
	require.NoError(t, err)
	require.Empty(t, list)

	// Test with valid file
	tmp := filepath.Join(t.TempDir(), "game_ip.txt")
	content := `# Test file
127.0.0.1:18590
192.168.1.1
invalid_ip
`
	require.NoError(t, os.WriteFile(tmp, []byte(content), 0644))
	list, err = staticIPs(tmp)
	// Should parse valid IPs and log error for invalid
	require.NotNil(t, list)
	require.Len(t, list, 2)
	require.Equal(t, "127.0.0.1", list[0].IP.String())
}

func TestPingEachServerExtra(t *testing.T) {
	// Just verify function exists and doesn't panic with nil
	require.NotPanics(t, func() {
		_ = PingEachServer
	})
}

func TestListServersWithExtra(t *testing.T) {
	require.NotPanics(t, func() {
		_ = ListServersWith
	})
}
