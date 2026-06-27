package cryptfile

import (
	"hash/crc32"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCRCUpdate(t *testing.T) {
	data := make([]byte, 256)
	for i := range data {
		data[i] = byte(i)
	}
	crc := crcUpdate(math.MaxUint32, data)
	scrc := crc32.Update(0, crcTable, data)
	require.Equal(t, crc, scrc)
}

func TestCryptXor(t *testing.T) {
	in := []byte{0x00, 0xFF, 0x55, 0xAA}
	want := []byte{0x7E, 0x81, 0x2B, 0xD4} // xor 126
	cryptXor(126, in)
	require.Equal(t, want, in)
	// double xor returns original
	cryptXor(126, in)
	require.Equal(t, []byte{0x00, 0xFF, 0x55, 0xAA}, in)
	// empty no panic
	cryptXor(126, nil)
	cryptXor(126, []byte{})
}

func TestCryptFileReadOnly(t *testing.T) {
	var cf *CryptFile
	require.False(t, cf.ReadOnly())
	cf = &CryptFile{readonly: true}
	require.True(t, cf.ReadOnly())
	cf = &CryptFile{readonly: false}
	require.False(t, cf.ReadOnly())
}

func TestSetXOR(t *testing.T) {
	cf := &CryptFile{}
	cf.SetXOR(true)
	require.True(t, cf.xor)
	cf.SetXOR(false)
	require.False(t, cf.xor)
}

func TestFlushCloseNil(t *testing.T) {
	var cf *CryptFile
	require.Equal(t, 0, cf.Flush())
	require.NoError(t, cf.Close())
	cf = &CryptFile{}
	// Flush on empty struct would panic because File is nil, so we test Close only which handles nil File
	require.NoError(t, cf.Close())
}

func TestCrcAdd(t *testing.T) {
	cf := &CryptFile{}
	// nil File should be safe and leave crcSum unchanged (default 0)
	cf.crcAdd([]byte{1, 2, 3})
	require.Equal(t, uint32(0), cf.crcSum)

	// with dummy File to test path, we can't easily construct binfile, but test nil receiver
	var nilCf *CryptFile
	nilCf.crcAdd([]byte{1})
}

func TestGlobal(t *testing.T) {
	require.Nil(t, Global())
	SetGlobal(nil)
	require.Nil(t, Global())
	cf := &CryptFile{readonly: true}
	SetGlobal(cf)
	require.Equal(t, cf, Global())
	require.NoError(t, Close())
	require.Nil(t, Global())
}

func TestOpenFileErrors(t *testing.T) {
	_, err := OpenFile("/nonexistent/path", ReadOnly, 0)
	require.Error(t, err)
	_, err = OpenFile("/nonexistent", WriteOnly, 0)
	require.Error(t, err)
}

func TestReadWriteUFunctions(t *testing.T) {
	// Test encoding logic via roundtrip in memory using a temp file
	tmp := filepath.Join(os.TempDir(), "crypt_test.bin")
	defer os.Remove(tmp)

	cf, err := OpenFile(tmp, WriteOnly, -1)
	require.NoError(t, err)
	require.False(t, cf.ReadOnly())
	require.NoError(t, cf.WriteU8(0xAB))
	require.NoError(t, cf.WriteU16(0x1234))
	require.NoError(t, cf.WriteU32(0x89ABCDEF))
	require.NoError(t, cf.WriteString8("hi"))
	require.NoError(t, cf.WriteString32("hello"))
	require.NoError(t, cf.Close())

	cf2, err := OpenFile(tmp, ReadOnly, -1)
	require.NoError(t, err)
	require.True(t, cf2.ReadOnly())
	v8, err := cf2.ReadU8()
	require.NoError(t, err)
	require.Equal(t, byte(0xAB), v8)
	v16, err := cf2.ReadU16()
	require.NoError(t, err)
	require.Equal(t, uint16(0x1234), v16)
	v32, err := cf2.ReadU32()
	require.NoError(t, err)
	require.Equal(t, uint32(0x89ABCDEF), v32)
	s8, err := cf2.ReadString8()
	require.NoError(t, err)
	require.Equal(t, "hi", s8)
	s32, err := cf2.ReadString32()
	require.NoError(t, err)
	require.Equal(t, "hello", s32)
	require.NoError(t, cf2.Close())
}

func TestReadWriteU(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "crypt_test2.bin")
	defer os.Remove(tmp)
	cf, err := OpenFile(tmp, WriteOnly, -1)
	require.NoError(t, err)
	v, err := cf.ReadWriteU8(5)
	require.NoError(t, err)
	require.Equal(t, byte(5), v)
	v16, err := cf.ReadWriteU16(0xBEEF)
	require.NoError(t, err)
	require.Equal(t, uint16(0xBEEF), v16)
	v32, err := cf.ReadWriteU32(0xDEADBEEF)
	require.NoError(t, err)
	require.Equal(t, uint32(0xDEADBEEF), v32)
	require.NoError(t, cf.Close())

	cf2, err := OpenFile(tmp, ReadOnly, -1)
	require.NoError(t, err)
	rv8, err := cf2.ReadWriteU8(0)
	require.NoError(t, err)
	require.Equal(t, byte(5), rv8)
	rv16, err := cf2.ReadWriteU16(0)
	require.NoError(t, err)
	require.Equal(t, uint16(0xBEEF), rv16)
	rv32, err := cf2.ReadWriteU32(0)
	require.NoError(t, err)
	require.Equal(t, uint32(0xDEADBEEF), rv32)
	require.NoError(t, cf2.Close())
}

func TestSectionStartEnd(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "crypt_test3.bin")
	defer os.Remove(tmp)
	cf, err := OpenFile(tmp, WriteOnly, -1)
	require.NoError(t, err)
	// SectionStart/End on write mode may depend on binfile internals; we test readonly no-op path for safety
	require.NoError(t, cf.Close())
	// readonly path should be no-op
	cf2, err := OpenFile(tmp, ReadOnly, -1)
	require.NoError(t, err)
	cf2.SectionStart()
	cf2.SectionEnd()
	require.NoError(t, cf2.Close())
}

func TestReadMaybeAlignAndOthers(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "crypt_test4.bin")
	defer os.Remove(tmp)
	cf, err := OpenFile(tmp, WriteOnly, -1)
	require.NoError(t, err)
	// ReadMaybeAlign on write mode returns nil
	require.NoError(t, cf.ReadMaybeAlign(make([]byte, 4)))
	// ReadWriteAlign
	cf.ReadWriteAlign()
	require.NoError(t, cf.Close())

	cf2, err := OpenFile(tmp, ReadOnly, -1)
	require.NoError(t, err)
	buf := make([]byte, 4)
	require.NoError(t, cf2.ReadMaybeAlign(buf))
	_, err = cf2.ReadAlignedU32()
	require.NoError(t, err)
	require.NoError(t, cf2.Close())
}
