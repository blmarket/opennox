package opennox

/*
#include <stdbool.h>
#include <stdio.h>
*/
import "C"
import (
	"bufio"
	"io"
	"os"
	"sync"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/datapath"
	"github.com/noxworld-dev/opennox-lib/ifs"

	"github.com/noxworld-dev/opennox/v1/common/alloc/handles"
)

var files struct {
	sync.RWMutex
	byHandle map[unsafe.Pointer]*File
}

type File struct {
	h unsafe.Pointer
	*os.File
	buf  *bufio.Reader
	err  error
	text bool
	bin  *Binfile
}

func (f *File) enableBuffer() {
	if f.buf != nil {
		return
	}
	f.buf = bufio.NewReader(f.File)
}

func (f *File) Seek(off int64, whence int) (int64, error) {
	if f.buf != nil {
		if whence == io.SeekCurrent {
			off -= int64(f.buf.Buffered())
		}
		f.buf = nil
	}
	n, err := f.File.Seek(off, whence)
	f.err = err
	return n, err
}

func (f *File) Read(p []byte) (int, error) {
	if f.buf != nil {
		n, err := f.buf.Read(p)
		f.err = err
		return n, err
	}
	n, err := f.File.Read(p)
	f.err = err
	return n, err
}

func (f *File) Write(p []byte) (int, error) {
	if f.buf != nil {
		panic("TODO: write on a buffered file")
	}
	n, err := f.File.Write(p)
	f.err = err
	return n, err
}

func (f *File) WriteString(p string) (int, error) {
	if f.buf != nil {
		panic("TODO: write on a buffered file")
	}
	n, err := f.File.WriteString(p)
	f.err = err
	return n, err
}

func (f *File) Close() error {
	if f.buf != nil {
		f.buf = nil
	}
	if f.bin != nil {
		if err := f.bin.close(); err != nil {
			_ = f.File.Close()
			return err
		}
	}
	return f.File.Close()
}

//export nox_fs_root
func nox_fs_root() *C.char {
	return internCStr(datapath.Data())
}

//export nox_fs_normalize
func nox_fs_normalize(path *C.char) *C.char {
	out := ifs.Normalize(GoString(path))
	return CString(out)
}

//export nox_fs_remove
func nox_fs_remove(path *C.char) C.bool {
	return ifs.Remove(GoString(path)) == nil
}

//export nox_fs_mkdir
func nox_fs_mkdir(path *C.char) C.bool {
	return ifs.Mkdir(GoString(path)) == nil
}

//export nox_fs_set_workdir
func nox_fs_set_workdir(path *C.char) C.bool {
	return ifs.Chdir(GoString(path)) == nil
}

//export nox_fs_copy
func nox_fs_copy(src, dst *C.char) C.bool {
	return ifs.Copy(GoString(src), GoString(dst)) == nil
}

//export nox_fs_move
func nox_fs_move(src, dst *C.char) C.bool {
	return ifs.Rename(GoString(src), GoString(dst)) == nil
}

func convWhence(mode C.int) int {
	var whence int
	switch mode {
	case C.SEEK_SET:
		whence = io.SeekStart
	case C.SEEK_CUR:
		whence = io.SeekCurrent
	case C.SEEK_END:
		whence = io.SeekEnd
	default:
		panic("unsupported seek mode")
	}
	return whence
}

//export nox_fs_fseek
func nox_fs_fseek(f *C.FILE, off C.long, mode C.int) C.int {
	fp := fileByHandle(f)
	_, err := fp.Seek(int64(off), convWhence(mode))
	if err != nil {
		return -1
	}
	return 0
}

//export nox_fs_ftell
func nox_fs_ftell(f *C.FILE) C.long {
	fp := fileByHandle(f)
	off, err := fp.Seek(0, io.SeekCurrent)
	if err != nil {
		e := int64(-1)
		return C.long(e)
	}
	return C.long(off)
}

//export nox_fs_fsize
func nox_fs_fsize(f *C.FILE) C.long {
	fp := fileByHandle(f)
	cur, err := fp.Seek(0, io.SeekCurrent)
	if err != nil {
		e := int64(-1)
		return C.long(e)
	}
	size, err := fp.Seek(0, io.SeekEnd)
	if err != nil {
		e := int64(-1)
		return C.long(e)
	}
	_, err = fp.Seek(cur, io.SeekStart)
	if err != nil {
		e := int64(-1)
		return C.long(e)
	}
	return C.long(size)
}

//export nox_fs_fread
func nox_fs_fread(f *C.FILE, dst unsafe.Pointer, sz C.int) C.int {
	fp := fileByHandle(f)
	n, _ := fp.Read(unsafe.Slice((*byte)(dst), int(sz)))
	return C.int(n)
}

//export nox_fs_fwrite
func nox_fs_fwrite(f *C.FILE, dst unsafe.Pointer, sz C.int) C.int {
	fp := fileByHandle(f)
	n, _ := fp.Write(unsafe.Slice((*byte)(dst), int(sz)))
	return C.int(n)
}

//export nox_fs_fgets
func nox_fs_fgets(f *C.FILE, dst *C.char, sz C.int) C.bool {
	fp := fileByHandle(f)
	fp.enableBuffer()
	var (
		out []byte
		end bool
	)
	for {
		b, err := fp.buf.ReadByte()
		fp.err = err
		if err == io.EOF {
			end = true
			break
		} else if err != nil {
			return false
		}
		out = append(out, b)
		if b == '\n' {
			break
		}
	}
	if n := len(out); n >= 2 && out[n-2] == '\r' && out[n-1] == '\n' {
		out[n-2] = '\n'
		out = out[:n-1]
	}
	StrCopy(dst, int(sz), string(out))
	return C.bool(!end)
}

//export nox_fs_fputs
func nox_fs_fputs(f *C.FILE, str *C.char) C.int {
	fp := fileByHandle(f)
	n, err := fp.WriteString(GoString(str))
	if err != nil {
		return -1
	}
	return C.int(n)
}

//export nox_fs_feof
func nox_fs_feof(f *C.FILE) C.bool {
	fp := fileByHandle(f)
	return fp.err == io.EOF
}

func fileByHandle(f *C.FILE) *File {
	h := unsafe.Pointer(f)
	handles.AssertValidPtr(h)
	files.RLock()
	fp := files.byHandle[h]
	files.RUnlock()
	return fp
}

//export nox_fs_close
func nox_fs_close(f *C.FILE) {
	if f == nil {
		return
	}
	h := unsafe.Pointer(f)
	handles.AssertValidPtr(h)
	files.Lock()
	defer files.Unlock()
	fp := files.byHandle[h]
	if fp != nil {
		_ = fp.Close()
		delete(files.byHandle, h)
	}
}

func newFileHandle(f *File) *C.FILE {
	if f.h != nil {
		return (*C.FILE)(f.h)
	}
	f.h = handles.NewPtr()
	files.Lock()
	defer files.Unlock()
	if files.byHandle == nil {
		files.byHandle = make(map[unsafe.Pointer]*File)
	}
	files.byHandle[f.h] = f
	return (*C.FILE)(f.h)
}

//export nox_fs_access
func nox_fs_access(path *C.char, mode C.int) C.int {
	_, err := ifs.Stat(GoString(path))
	if os.IsNotExist(err) {
		return -1
	} else if err != nil {
		return -2
	}
	return 0
}

//export nox_fs_open
func nox_fs_open(path *C.char) *C.FILE {
	f, err := ifs.Open(GoString(path))
	if err != nil {
		return nil
	}
	return newFileHandle(&File{File: f})
}

//export nox_fs_open_text
func nox_fs_open_text(path *C.char) *C.FILE {
	f, err := ifs.Open(GoString(path))
	if err != nil {
		return nil
	}
	return newFileHandle(&File{File: f, text: true})
}

//export nox_fs_create
func nox_fs_create(path *C.char) *C.FILE {
	f, err := ifs.Create(GoString(path))
	if err != nil {
		return nil
	}
	return newFileHandle(&File{File: f})
}

//export nox_fs_create_text
func nox_fs_create_text(path *C.char) *C.FILE {
	f, err := ifs.Create(GoString(path))
	if err != nil {
		return nil
	}
	return newFileHandle(&File{File: f, text: true})
}

//export nox_fs_open_rw
func nox_fs_open_rw(path *C.char) *C.FILE {
	f, err := ifs.OpenFile(GoString(path), os.O_RDWR)
	if err != nil {
		return nil
	}
	return newFileHandle(&File{File: f})
}
