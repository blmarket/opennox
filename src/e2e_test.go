package opennox

import (
	"os"
	"path/filepath"
	"testing"
)

func TestImageDiff(t *testing.T) {
	pix1 := []byte{10, 20, 30, 255, 0, 0, 0, 0}
	pix2 := []byte{10, 20, 30, 255, 5, 5, 5, 128}
	out := imageDiff(pix1, pix2)
	if len(out) != len(pix1) {
		t.Fatalf("expected len %d, got %d", len(pix1), len(out))
	}
	// First pixel identical, diff should be 0 for RGB, 255 for alpha (0xff - 0)
	if out[0] != 0 || out[1] != 0 || out[2] != 0 {
		t.Errorf("expected zero diff for identical RGB, got %v", out[:3])
	}
	if out[3] != 255 {
		t.Errorf("expected 255 for alpha channel diff, got %d", out[3])
	}
	// Second pixel diff: 5*10=50 for RGB, alpha: 128*10=1280 clamped to 255, then 0xff-255=0
	if out[4] != 50 || out[5] != 50 || out[6] != 50 {
		t.Errorf("expected 50 for RGB diff, got %v", out[4:7])
	}
	if out[7] != 0 {
		t.Errorf("expected 0 for alpha diff, got %d", out[7])
	}
}

func TestE2eAngToPos(t *testing.T) {
	// Angle 0 should point up (negative Y)
	p := e2eAngToPos(0, 100)
	if p.X != 512 {
		t.Errorf("expected X=512 for angle 0, got %d", p.X)
	}
	if p.Y >= 384 {
		t.Errorf("expected Y < 384 for angle 0 (up), got %d", p.Y)
	}
	// Angle 0.25 should point right (approximately, Y near center)
	p2 := e2eAngToPos(0.25, 100)
	if p2.X <= 512 {
		t.Errorf("expected X > 512 for angle 0.25 (right), got %d", p2.X)
	}
	// Y should be near center (384) for right direction, allow some tolerance
	if p2.Y < 300 || p2.Y > 400 {
		t.Errorf("expected Y near 384 for angle 0.25, got %d", p2.Y)
	}
	// Distance 0 should be center
	p3 := e2eAngToPos(0.5, 0)
	if p3.X != 512 || p3.Y != 384 {
		t.Errorf("expected center (512,384) for dist 0, got %v", p3)
	}
}

func TestE2eHash(t *testing.T) {
	h := e2eHash()
	if h == nil {
		t.Fatal("expected non-nil hash")
	}
	h.Write([]byte("test"))
	sum := h.Sum(nil)
	if len(sum) != 32 {
		t.Errorf("expected 32 byte sum for blake2b-256, got %d", len(sum))
	}
}

func TestE2eAbsPath(t *testing.T) {
	// Absolute path should be returned as-is (or cleaned)
	abs := "/tmp/testpath"
	out := e2eAbsPath(abs)
	if out != abs {
		t.Errorf("expected %s, got %s", abs, out)
	}
	// Relative path that exists should be resolved
	tmpDir := t.TempDir()
	relFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(relFile, []byte("data"), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	// Change to tmpDir and test relative path
	origWd, _ := os.Getwd()
	defer os.Chdir(origWd)
	os.Chdir(tmpDir)
	out2 := e2eAbsPath("test.txt")
	if !filepath.IsAbs(out2) {
		t.Errorf("expected absolute path, got %s", out2)
	}
}

func TestE2eHashDir(t *testing.T) {
	tmpDir := t.TempDir()
	// Create a couple of files
	if err := os.WriteFile(filepath.Join(tmpDir, "a.txt"), []byte("hello"), 0644); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	subDir := filepath.Join(tmpDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create subdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(subDir, "b.txt"), []byte("world"), 0644); err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	hashes := e2eHashDir(tmpDir)
	if len(hashes) != 2 {
		t.Errorf("expected 2 hashes, got %d: %v", len(hashes), hashes)
	}
	if _, ok := hashes["a.txt"]; !ok {
		t.Error("expected hash for a.txt")
	}
	if _, ok := hashes["sub/b.txt"]; !ok {
		t.Error("expected hash for sub/b.txt")
	}
	// Empty dir should return empty map
	emptyDir := t.TempDir()
	emptyHashes := e2eHashDir(emptyDir)
	if len(emptyHashes) != 0 {
		t.Errorf("expected empty hashes for empty dir, got %v", emptyHashes)
	}
}

func TestE2eScenarioMethods(t *testing.T) {
	var sc e2eScenario
	// Test basic methods that don't require game state
	sc.Slow(10)
	if len(sc.steps) != 1 {
		t.Errorf("expected 1 step after Slow, got %d", len(sc.steps))
	}
	sc.Wait(5, "testwait")
	sc.Move(100, 200, "testmove")
	sc.ClickLeft(10, 20, "testclick")
	sc.ClickSlowLeft(30, 40, "testclickslow")
	sc.Key(1, "testkey")
	sc.WalkFor(0.5, 10, "testwalk")
	sc.RunFor(0.25, 10, "testrun")
	sc.Melee(0, "testmelee")
	sc.Save("testsave", map[string]string{"a": "b"})
	sc.Screen("testscreen")
	if len(sc.steps) == 0 {
		t.Error("expected steps to be added")
	}
}
