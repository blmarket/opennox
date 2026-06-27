package cnxz

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/noxtest"
	"github.com/stretchr/testify/require"
)

func TestSyntheticRoundTrip(t *testing.T) {
	pseudorandom := make([]byte, 64*1024)
	state := uint32(0x6d2b79f5)
	for i := range pseudorandom {
		state ^= state << 13
		state ^= state >> 17
		state ^= state << 5
		pseudorandom[i] = byte(state)
	}

	chunkBoundary := make([]byte, 500001)
	for i := range chunkBoundary {
		chunkBoundary[i] = byte(i*31 + i/251)
	}

	tests := []struct {
		name string
		data []byte
	}{
		{name: "single byte", data: []byte{0xa5}},
		{name: "short literal", data: []byte("OpenNox NXZ characterization")},
		{name: "repeated byte", data: bytes.Repeat([]byte{0x7f}, 64*1024)},
		{name: "repeated sequence", data: bytes.Repeat([]byte("ABRACADABRA"), 8192)},
		{name: "pseudorandom", data: pseudorandom},
		{name: "compressor chunk boundary", data: chunkBoundary},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			src := filepath.Join(dir, "source.bin")
			compressed := filepath.Join(dir, "source.nxz")
			decoded := filepath.Join(dir, "decoded.bin")

			require.NoError(t, os.WriteFile(src, tc.data, 0o600))
			require.NoError(t, CompressFile(src, compressed))
			require.NoError(t, DecompressFile(compressed, decoded))
			got, err := os.ReadFile(decoded)
			require.NoError(t, err)
			require.Equal(t, tc.data, got)
		})
	}
}

func TestDecompress(t *testing.T) {
	maps := noxtest.DataPath(t, "maps")
	files, err := os.ReadDir(maps)
	require.NoError(t, err)
	for _, fi := range files {
		mname := filepath.Join(maps, fi.Name(), fi.Name()+".map")
		zname := filepath.Join(maps, fi.Name(), fi.Name()+".nxz")
		if _, err = ifs.Stat(zname); err != nil {
			continue
		}
		t.Run(fi.Name(), func(t *testing.T) {
			mexp, mexpN := hashFile(t, mname)
			gotc, gotcN := decompressC(t, zname)
			require.Equal(t, mexpN, gotcN)
			require.Equal(t, mexp, gotc)
		})
	}
}

func TestCompress(t *testing.T) {
	maps := noxtest.DataPath(t, "maps")
	files, err := os.ReadDir(maps)
	require.NoError(t, err)
	for _, fi := range files {
		mname := filepath.Join(maps, fi.Name(), fi.Name()+".map")
		zname := filepath.Join(maps, fi.Name(), fi.Name()+".nxz")
		if _, err = ifs.Stat(zname); err != nil {
			continue
		}
		t.Run(fi.Name(), func(t *testing.T) {
			mexp, mexpN := hashFile(t, zname)
			gotc, gotcN := compressC(t, mname)
			require.Equal(t, mexpN, gotcN)
			require.Equal(t, mexp, gotc)
		})
	}
}

func decompressC(t testing.TB, path string) (string, int) {
	out, err := os.CreateTemp("", "nxzmap_*.map")
	require.NoError(t, err)
	defer func() {
		out.Close()
		_ = os.Remove(out.Name())
	}()
	err = DecompressFile(path, out.Name())
	require.NoError(t, err)
	return hashFile(t, out.Name())
}

func compressC(t testing.TB, path string) (string, int) {
	out, err := os.CreateTemp("", "nxzmap_*.nxz")
	require.NoError(t, err)
	defer func() {
		out.Close()
		_ = os.Remove(out.Name())
	}()
	err = CompressFile(path, out.Name())
	require.NoError(t, err)
	return hashFile(t, out.Name())
}

func hashReader(t testing.TB, r io.Reader) (string, int) {
	h := sha1.New()
	n, err := io.Copy(h, r)
	require.NoError(t, err)
	return hex.EncodeToString(h.Sum(nil)), int(n)
}

func hashFile(t testing.TB, path string) (string, int) {
	f, err := ifs.Open(path)
	require.NoError(t, err)
	defer f.Close()
	return hashReader(t, f)
}
