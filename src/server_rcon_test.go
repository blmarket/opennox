package opennox

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/noxworld-dev/opennox-lib/console"
)

// mockChannel implements ssh.Channel for testing
type mockChannel struct {
	buf []byte
}

func (m *mockChannel) Read(data []byte) (int, error) { return 0, nil }
func (m *mockChannel) Write(data []byte) (int, error) {
	m.buf = append(m.buf, data...)
	return len(data), nil
}
func (m *mockChannel) Close() error      { return nil }
func (m *mockChannel) CloseWrite() error { return nil }
func (m *mockChannel) SendRequest(name string, wantReply bool, payload []byte) (bool, error) {
	return true, nil
}
func (m *mockChannel) Stderr() io.ReadWriter { return m }

func TestEncodePEM(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}
	data := encodePEM(key)
	if len(data) == 0 {
		t.Fatal("encodePEM returned empty data")
	}
	block, _ := pem.Decode(data)
	if block == nil {
		t.Fatal("failed to decode PEM")
	}
	if block.Type != "RSA PRIVATE KEY" {
		t.Errorf("unexpected block type: %s", block.Type)
	}
	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf("failed to parse key: %v", err)
	}
	if pk.N.Cmp(key.N) != 0 {
		t.Error("parsed key does not match original")
	}
}

func TestRconGenKeyReadKey(t *testing.T) {
	dir := t.TempDir()
	keyPath := filepath.Join(dir, "test.pem")

	rc := &RemoteConsole{}
	pk, err := rc.genKey(keyPath)
	if err != nil {
		t.Fatalf("genKey failed: %v", err)
	}
	if pk == nil {
		t.Fatal("genKey returned nil key")
	}

	// Check file exists and has correct permissions
	info, err := os.Stat(keyPath)
	if err != nil {
		t.Fatalf("key file not created: %v", err)
	}
	if info.Mode().Perm() != 0600 {
		t.Errorf("unexpected file mode: %v", info.Mode().Perm())
	}

	// Read back the key
	pk2, err := rc.readKey(keyPath)
	if err != nil {
		t.Fatalf("readKey failed: %v", err)
	}
	if pk2.N.Cmp(pk.N) != 0 {
		t.Error("read key does not match generated key")
	}

	// Test read non-existent file
	_, err = rc.readKey(filepath.Join(dir, "nonexistent.pem"))
	if err == nil {
		t.Error("expected error for non-existent file")
	}

	// Test read invalid PEM
	invalidPath := filepath.Join(dir, "invalid.pem")
	os.WriteFile(invalidPath, []byte("not a pem"), 0600)
	_, err = rc.readKey(invalidPath)
	if err == nil {
		t.Error("expected error for invalid PEM")
	}
}

func TestRconLoadKey(t *testing.T) {
	dir := t.TempDir()
	keyPath := filepath.Join(dir, "test.pem")

	rc := &RemoteConsole{}
	err := rc.loadKey(keyPath)
	if err != nil {
		t.Fatalf("loadKey failed for non-existent key (should generate): %v", err)
	}
	// If loadKey succeeds, host key should be added

	// Load existing key
	rc2 := &RemoteConsole{}
	err = rc2.loadKey(keyPath)
	if err != nil {
		t.Fatalf("loadKey failed for existing key: %v", err)
	}
}

func TestRconAuthPassword(t *testing.T) {
	rc := &RemoteConsole{
		opts: RconOptions{Pass: "secret123"},
	}

	// Correct password
	perm, err := rc.authPassword(nil, []byte("secret123"))
	if err != nil {
		t.Fatalf("auth with correct password failed: %v", err)
	}
	if perm == nil {
		t.Error("expected non-nil permissions")
	}

	// Wrong password
	_, err = rc.authPassword(nil, []byte("wrong"))
	if err == nil {
		t.Error("expected error for wrong password")
	}
	if err != errRconInvalidPass {
		t.Errorf("unexpected error: %v", err)
	}

	// Wrong length
	_, err = rc.authPassword(nil, []byte("short"))
	if err == nil {
		t.Error("expected error for wrong length password")
	}
}

func TestNewRemoteConsole(t *testing.T) {
	dir := t.TempDir()
	keyPath := filepath.Join(dir, "test.pem")

	// Test with password auth
	rc, err := NewRemoteConsole("127.0.0.1:0", nil, RconOptions{
		Pass:    "testpass",
		KeyPath: keyPath,
	})
	if err != nil {
		t.Fatalf("NewRemoteConsole failed: %v", err)
	}
	defer rc.Close()

	if rc.conf.PasswordCallback == nil {
		t.Error("expected password callback to be set")
	}
	if rc.conf.NoClientAuth {
		t.Error("expected NoClientAuth to be false when password is set")
	}

	// Test without password (no auth)
	rc2, err := NewRemoteConsole("127.0.0.1:0", nil, RconOptions{
		KeyPath: filepath.Join(dir, "test2.pem"),
	})
	if err != nil {
		t.Fatalf("NewRemoteConsole without password failed: %v", err)
	}
	defer rc2.Close()

	if !rc2.conf.NoClientAuth {
		t.Error("expected NoClientAuth to be true when no password")
	}

	// Test with invalid host
	_, err = NewRemoteConsole("invalid:host:port", nil, RconOptions{
		KeyPath: filepath.Join(dir, "test3.pem"),
	})
	if err == nil {
		t.Error("expected error for invalid host")
	}
}

func TestRconPrint(t *testing.T) {
	rc := &RemoteConsole{}
	rc.sessions.list = make(map[*rcShell]struct{})

	// Create a mock shell
	mch := &mockChannel{}
	sh := &rcShell{
		rc: rc,
		ch: mch,
		bw: bufio.NewWriter(mch),
	}
	rc.sessions.list[sh] = struct{}{}

	rc.Print(console.ColorRed, "test message")
	rc.Printf(console.ColorGreen, "formatted %s", "message")

	// Check that data was written
	if len(mch.buf) == 0 {
		t.Error("expected data to be written to channel")
	}
}

func TestRcShellExec(t *testing.T) {
	var executed []string
	execFn := func(ctx context.Context, cmd string) bool {
		executed = append(executed, cmd)
		return true
	}

	rc := &RemoteConsole{exec: execFn}
	sh := &rcShell{
		rc: rc,
		history: struct {
			list []string
			cur  int
		}{},
	}

	sh.Exec("test command")
	if len(executed) != 1 || executed[0] != "test command" {
		t.Errorf("unexpected executed commands: %v", executed)
	}

	// Test empty command
	sh.Exec("")
	if len(executed) != 1 {
		t.Error("empty command should not be executed")
	}

	// Test duplicate command not added to history twice
	sh.Exec("test command")
	if len(sh.history.list) != 1 {
		t.Errorf("duplicate command should not be added to history, got %d entries", len(sh.history.list))
	}

	sh.Exec("another command")
	if len(sh.history.list) != 2 {
		t.Errorf("expected 2 history entries, got %d", len(sh.history.list))
	}
}

func TestRcShellDoCmd(t *testing.T) {
	rc := &RemoteConsole{}
	mch := &mockChannel{}
	sh := &rcShell{
		rc:   rc,
		ch:   mch,
		bw:   bufio.NewWriter(mch),
		line: []rune("quit"),
	}

	// Test quit command
	result := sh.doCmd()
	if result != false {
		t.Error("quit command should return false")
	}

	// Test other command
	sh.line = []rune("test")
	sh.history = struct {
		list []string
		cur  int
	}{}
	rc.exec = func(ctx context.Context, cmd string) bool { return true }
	result = sh.doCmd()
	if result != true {
		t.Error("non-quit command should return true")
	}
	if len(sh.line) != 0 {
		t.Error("line should be cleared after doCmd")
	}
}

func TestRcShellAddRune(t *testing.T) {
	rc := &RemoteConsole{}
	mch := &mockChannel{}
	sh := &rcShell{
		rc: rc,
		ch: mch,
		bw: bufio.NewWriter(mch),
		br: bufio.NewReader(nil),
	}

	// Test regular character
	result := sh.addRune('a')
	if !result {
		t.Error("addRune should return true for regular char")
	}
	if string(sh.line) != "a" {
		t.Errorf("unexpected line content: %q", string(sh.line))
	}

	// Test backspace
	result = sh.addRune('\b')
	if !result {
		t.Error("addRune should return true for backspace")
	}
	if len(sh.line) != 0 {
		t.Error("line should be empty after backspace")
	}

	// Test enter
	sh.line = []rune("test")
	rc.exec = func(ctx context.Context, cmd string) bool { return true }
	result = sh.addRune('\r')
	if !result {
		t.Error("addRune should return true for enter with valid command")
	}

	// Test Ctrl+C
	sh.line = []rune("test")
	result = sh.addRune('\x03')
	if result {
		t.Error("addRune should return false for Ctrl+C")
	}

	// Test escape sequence
	// Use a reader that returns EOF immediately
	sh.br = bufio.NewReader(strings.NewReader(""))
	result = sh.addRune('\033')
	if result {
		t.Error("addRune should return false on EOF during escape sequence")
	}
}

func TestRcShellCursorMovement(t *testing.T) {
	rc := &RemoteConsole{}
	mch := &mockChannel{}
	sh := &rcShell{
		rc:     rc,
		ch:     mch,
		bw:     bufio.NewWriter(mch),
		line:   []rune("test"),
		cursor: 2,
	}

	// Test cursor left
	sh.cursorLeft()
	if sh.cursor != 1 {
		t.Errorf("cursor should be at 1 after left, got %d", sh.cursor)
	}

	// Test cursor left at beginning
	sh.cursor = 0
	sh.cursorLeft()
	if sh.cursor != 0 {
		t.Error("cursor should stay at 0 when at beginning")
	}

	// Test cursor right
	sh.cursor = 1
	sh.cursorRight()
	if sh.cursor != 2 {
		t.Errorf("cursor should be at 2 after right, got %d", sh.cursor)
	}

	// Test cursor right at end
	sh.cursor = len(sh.line)
	sh.cursorRight()
	if sh.cursor != len(sh.line) {
		t.Error("cursor should stay at end")
	}
}

func TestRcShellEraseOne(t *testing.T) {
	rc := &RemoteConsole{}
	mch := &mockChannel{}
	sh := &rcShell{
		rc:     rc,
		ch:     mch,
		bw:     bufio.NewWriter(mch),
		line:   []rune("test"),
		cursor: 4,
	}

	sh.eraseOne()
	if string(sh.line) != "tes" {
		t.Errorf("unexpected line after erase: %q", string(sh.line))
	}
	if sh.cursor != 3 {
		t.Errorf("cursor should be at 3, got %d", sh.cursor)
	}

	// Test erase at beginning
	sh.cursor = 0
	sh.eraseOne()
	if string(sh.line) != "tes" {
		t.Error("erase at beginning should do nothing")
	}

	// Test erase in middle
	sh.line = []rune("test")
	sh.cursor = 2
	sh.eraseOne()
	if string(sh.line) != "tst" {
		t.Errorf("unexpected line after erase in middle: %q", string(sh.line))
	}
}

func TestRcShellDelOne(t *testing.T) {
	rc := &RemoteConsole{}
	mch := &mockChannel{}
	sh := &rcShell{
		rc:     rc,
		ch:     mch,
		bw:     bufio.NewWriter(mch),
		line:   []rune("test"),
		cursor: 1,
	}

	sh.delOne()
	if string(sh.line) != "tst" {
		t.Errorf("unexpected line after del: %q", string(sh.line))
	}

	// Test del at end
	sh.cursor = len(sh.line)
	sh.delOne()
	if string(sh.line) != "tst" {
		t.Error("del at end should do nothing")
	}
}

func TestRcShellRecall(t *testing.T) {
	rc := &RemoteConsole{}
	mch := &mockChannel{}
	sh := &rcShell{
		rc:   rc,
		ch:   mch,
		bw:   bufio.NewWriter(mch),
		line: []rune("current"),
		history: struct {
			list []string
			cur  int
		}{
			list: []string{"cmd1", "cmd2", "cmd3"},
			cur:  1,
		},
	}

	sh.recall()
	if string(sh.line) != "cmd2" {
		t.Errorf("unexpected line after recall: %q", string(sh.line))
	}

	// Test recall with invalid index
	sh.history.cur = 10
	sh.line = []rune("unchanged")
	sh.recall()
	if string(sh.line) != "unchanged" {
		t.Error("recall with invalid index should do nothing")
	}
}

func TestRcShellClose(t *testing.T) {
	rc := &RemoteConsole{}
	rc.sessions.list = make(map[*rcShell]struct{})
	mch := &mockChannel{}
	sh := &rcShell{
		rc: rc,
		ch: mch,
		bw: bufio.NewWriter(mch),
	}
	rc.sessions.list[sh] = struct{}{}

	err := sh.Close()
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}
	if len(rc.sessions.list) != 0 {
		t.Error("shell should be removed from sessions list")
	}
}
