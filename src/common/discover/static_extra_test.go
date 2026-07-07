package discover

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStaticIPs(t *testing.T) {
	// Test with non-existent file
	servers, err := staticIPs("/non/existent/file.txt")
	if err != nil {
		t.Errorf("staticIPs with non-existent file should not return error, got %v", err)
	}
	if len(servers) != 0 {
		t.Errorf("staticIPs with non-existent file should return empty list, got %d servers", len(servers))
	}

	// Create a temporary file with test data
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "game_ip.txt")

	content := `# Test game IP file
127.0.0.1:18590
192.168.1.1
# Comment line
10.0.0.1:12345
invalid_ip
`
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	servers, err = staticIPs(tmpFile)
	if err == nil {
		t.Error("staticIPs with invalid IP should return an error")
	}
	// Should have 3 valid servers (127.0.0.1:18590, 192.168.1.1 with default port, 10.0.0.1:12345)
	if len(servers) != 3 {
		t.Errorf("staticIPs should return 3 valid servers, got %d", len(servers))
	}

	// Check first server
	if len(servers) > 0 {
		if servers[0].IP.String() != "127.0.0.1" {
			t.Errorf("First server IP = %q, want %q", servers[0].IP.String(), "127.0.0.1")
		}
		if servers[0].Game.Port != 18590 {
			t.Errorf("First server port = %d, want %d", servers[0].Game.Port, 18590)
		}
	}

	// Test with empty file
	emptyFile := filepath.Join(tmpDir, "empty.txt")
	if err := os.WriteFile(emptyFile, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create empty file: %v", err)
	}
	servers, err = staticIPs(emptyFile)
	if err != nil {
		t.Errorf("staticIPs with empty file should not return error, got %v", err)
	}
	if len(servers) != 0 {
		t.Errorf("staticIPs with empty file should return empty list, got %d servers", len(servers))
	}

	// Test with only comments
	commentFile := filepath.Join(tmpDir, "comments.txt")
	if err := os.WriteFile(commentFile, []byte("# Only comments\n# Another comment\n"), 0644); err != nil {
		t.Fatalf("Failed to create comment file: %v", err)
	}
	servers, err = staticIPs(commentFile)
	if err != nil {
		t.Errorf("staticIPs with only comments should not return error, got %v", err)
	}
	if len(servers) != 0 {
		t.Errorf("staticIPs with only comments should return empty list, got %d servers", len(servers))
	}
}

func TestStaticIPsInvalidFile(t *testing.T) {
	// Test with a directory instead of a file
	tmpDir := t.TempDir()
	servers, err := staticIPs(tmpDir)
	if err == nil {
		t.Error("staticIPs with directory should return an error")
	}
	if len(servers) != 0 {
		t.Errorf("staticIPs with directory should return empty list, got %d servers", len(servers))
	}
}
