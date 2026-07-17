package legacy

/*
#include <stdio.h>

int nox_script_readWriteYyy_542380(FILE* f1, FILE* f2, int relocate);
int nox_script_readWriteJjj_5418C0(FILE* f1, FILE* f2, FILE* out);
int nox_script_readWriteVvv_541E40(FILE* f1, FILE* f2, FILE* out);
int nox_script_readWriteIii_541D80(FILE* f1, FILE* out);
int nox_script_readWriteXxx_541A50(FILE* f1, FILE* f2, FILE* out);
size_t nox_script_readWriteWww_5417C0(FILE* f1, FILE* f2, FILE* out);
*/
import "C"

import (
	"io"
	"os"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
)

func scriptFileHandle(data []byte) (*os.File, *FILE, error) {
	f, err := os.CreateTemp("", "opennox-script-coverage-*")
	if err != nil {
		return nil, nil, err
	}
	if len(data) != 0 {
		if _, err := f.Write(data); err != nil {
			f.Close()
			os.Remove(f.Name())
			return nil, nil, err
		}
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		f.Close()
		os.Remove(f.Name())
		return nil, nil, err
	}
	return f, NewFileHandle(binfile.NewFile(f)), nil
}

func scriptFileOutput(f *os.File) ([]byte, error) {
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	return io.ReadAll(f)
}

func closeScriptFile(f *os.File, h *FILE) {
	Nox_fs_close(h)
	os.Remove(f.Name())
}

func C_scriptReadWriteOpcodes(input []byte, relocate bool) (int, []byte, error) {
	in, hin, err := scriptFileHandle(input)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(in, hin)
	out, hout, err := scriptFileHandle(nil)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(out, hout)
	var rel C.int
	if relocate {
		rel = 1
	}
	ret := int(C.nox_script_readWriteYyy_542380((*C.FILE)(hin), (*C.FILE)(hout), rel))
	got, err := scriptFileOutput(out)
	return ret, got, err
}

func C_scriptMergeStrings(first, second []byte) (int, []byte, error) {
	f1, h1, err := scriptFileHandle(first)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(f1, h1)
	f2, h2, err := scriptFileHandle(second)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(f2, h2)
	out, hout, err := scriptFileHandle(nil)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(out, hout)
	ret := int(C.nox_script_readWriteJjj_5418C0((*C.FILE)(h1), (*C.FILE)(h2), (*C.FILE)(hout)))
	got, err := scriptFileOutput(out)
	return ret, got, err
}

func C_scriptCopyIntSection(input []byte) (int, []byte, error) {
	in, hin, err := scriptFileHandle(input)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(in, hin)
	out, hout, err := scriptFileHandle(nil)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(out, hout)
	ret := int(C.nox_script_readWriteIii_541D80((*C.FILE)(hin), (*C.FILE)(hout)))
	got, err := scriptFileOutput(out)
	return ret, got, err
}

func C_scriptMergeVariableSection(first, second []byte) (int, []byte, error) {
	f1, h1, err := scriptFileHandle(first)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(f1, h1)
	f2, h2, err := scriptFileHandle(second)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(f2, h2)
	out, hout, err := scriptFileHandle(nil)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(out, hout)
	ret := int(C.nox_script_readWriteVvv_541E40((*C.FILE)(h1), (*C.FILE)(h2), (*C.FILE)(hout)))
	got, err := scriptFileOutput(out)
	return ret, got, err
}

func C_scriptMergeCodeSections(first, second []byte) (int, []byte, error) {
	f1, h1, err := scriptFileHandle(first)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(f1, h1)
	f2, h2, err := scriptFileHandle(second)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(f2, h2)
	out, hout, err := scriptFileHandle(nil)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(out, hout)
	ret := int(C.nox_script_readWriteXxx_541A50((*C.FILE)(h1), (*C.FILE)(h2), (*C.FILE)(hout)))
	got, err := scriptFileOutput(out)
	return ret, got, err
}

func C_scriptMergeContainers(first, second []byte) (int, []byte, error) {
	f1, h1, err := scriptFileHandle(first)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(f1, h1)
	f2, h2, err := scriptFileHandle(second)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(f2, h2)
	out, hout, err := scriptFileHandle(nil)
	if err != nil {
		return 0, nil, err
	}
	defer closeScriptFile(out, hout)
	ret := int(C.nox_script_readWriteWww_5417C0((*C.FILE)(h1), (*C.FILE)(h2), (*C.FILE)(hout)))
	got, err := scriptFileOutput(out)
	return ret, got, err
}
