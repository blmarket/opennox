package legacy

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"

	"github.com/noxworld-dev/noxscript/ns/asm"
	"github.com/stretchr/testify/require"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

func scriptPut32(buf *bytes.Buffer, values ...uint32) {
	for _, v := range values {
		_ = binary.Write(buf, binary.LittleEndian, v)
	}
}

func scriptOpcodeCorpus() []byte {
	var buf bytes.Buffer
	for op := uint32(0); op <= 72; op++ {
		scriptPut32(&buf, op)
		switch op {
		case 0, 1, 2:
			scriptPut32(&buf, 1, 10+op)
		case 3:
			scriptPut32(&buf, 1, 2)
		case 4:
			scriptPut32(&buf, 4)
		case 5:
			scriptPut32(&buf, math.Float32bits(1.25))
		case 6:
			scriptPut32(&buf, 6)
		case 19, 20, 21:
			scriptPut32(&buf, op+100)
		case 69:
			scriptPut32(&buf, 126)
		case 70:
			scriptPut32(&buf, 20)
		}
	}
	return buf.Bytes()
}

func scriptVariableSection(seed uint32, count int) []byte {
	var b bytes.Buffer
	put := func(values ...uint32) { scriptPut32(&b, values...) }
	put(seed)
	put(4)
	b.Write([]byte{byte(seed), 2, 3, 4})
	put(seed+10, seed+11, seed+12)
	put(uint32(count))
	put(seed + 13)
	put(seed+14, seed+15, seed+16, seed+17)
	put(4)
	b.Write([]byte{5, 6, 7, byte(seed)})
	put(seed+18, seed+19, seed+20)
	put(uint32(count))
	put(seed + 21)
	for i := 0; i < count; i++ {
		put(seed + 30 + uint32(i))
	}
	put(seed+40, uint32(count))
	b.Write(scriptOpcodeCorpus())
	return b.Bytes()
}

func scriptFunctionRecord(seed uint32, name string) []byte {
	var b bytes.Buffer
	scriptPut32(&b, seed)
	scriptPut32(&b, uint32(len(name)))
	b.WriteString(name)
	scriptPut32(&b, seed+1, seed+2)
	// Integer subsection: header, count, trailer, and count values.
	scriptPut32(&b, seed+3, 2, seed+4, seed+5, seed+6)
	scriptPut32(&b, seed+7, seed+8)
	b.Write(scriptOpcodeCorpus())
	return b.Bytes()
}

func scriptCodeSection(seed uint32, name string) []byte {
	var b bytes.Buffer
	b.WriteString("CODE")
	scriptPut32(&b, 3)
	b.Write(scriptVariableSection(seed, 5))
	b.Write(scriptFunctionRecord(seed+100, name))
	return b.Bytes()
}

func scriptStringSection(values ...string) []byte {
	var b bytes.Buffer
	b.WriteString("STRG")
	scriptPut32(&b, uint32(len(values)))
	for _, value := range values {
		scriptPut32(&b, uint32(len(value)))
		b.WriteString(value)
	}
	return b.Bytes()
}

func scriptContainer(seed uint32, name string, strings ...string) []byte {
	var b bytes.Buffer
	b.WriteString("SCRIPT03")
	b.Write(scriptStringSection(strings...))
	b.Write(scriptCodeSection(seed, name))
	b.WriteString("DONE")
	return b.Bytes()
}

func TestLegacyScriptFileTransforms(t *testing.T) {
	handles.Init()
	oldMore := Nox_script_shouldReadMoreXxx
	oldEvenMore := Nox_script_shouldReadEvenMoreXxx
	Nox_script_shouldReadMoreXxx = func(fi asm.Builtin) bool {
		return fi == asm.BuiltinSetDialog
	}
	Nox_script_shouldReadEvenMoreXxx = func(fi asm.Builtin) bool {
		return fi == asm.BuiltinSetDialog
	}
	t.Cleanup(func() {
		Nox_script_shouldReadMoreXxx = oldMore
		Nox_script_shouldReadEvenMoreXxx = oldEvenMore
	})

	t.Run("all bytecode opcodes", func(t *testing.T) {
		input := scriptOpcodeCorpus()
	ret, copied, err := C_scriptReadWriteOpcodes(input, false)
	require.NoError(t, err)
	require.Equal(t, 1, ret)
		require.Equal(t, input, copied)

		ret, relocated, err := C_scriptReadWriteOpcodes(input, true)
		require.NoError(t, err)
		require.Equal(t, 1, ret)
		require.Len(t, relocated, len(input))
		require.NotEqual(t, input, relocated)
	})

	t.Run("string table merge", func(t *testing.T) {
		var a, b bytes.Buffer
		a.WriteString("STRG")
		scriptPut32(&a, 2, 3)
		a.WriteString("one")
		scriptPut32(&a, 3)
		a.WriteString("two")
		b.WriteString("STRG")
		scriptPut32(&b, 1, 5)
		b.WriteString("three")
		ret, merged, err := C_scriptMergeStrings(a.Bytes(), b.Bytes())
		require.NoError(t, err)
		require.Equal(t, 1, ret)
		require.Equal(t, "STRG", string(merged[:4]))
		require.Equal(t, uint32(3), binary.LittleEndian.Uint32(merged[4:8]))
		require.Contains(t, string(merged), "one")
		require.Contains(t, string(merged), "three")
	})

	t.Run("integer section copy", func(t *testing.T) {
		var in bytes.Buffer
		scriptPut32(&in, 7, 3, 9, 11, 12, 13)
		ret, copied, err := C_scriptCopyIntSection(in.Bytes())
		require.NoError(t, err)
		require.Equal(t, 3, ret)
		require.Equal(t, in.Bytes(), copied)

		var empty bytes.Buffer
		scriptPut32(&empty, 7, 0, 9)
		ret, copied, err = C_scriptCopyIntSection(empty.Bytes())
		require.NoError(t, err)
		require.Zero(t, ret)
		require.Equal(t, empty.Bytes(), copied)
	})

	t.Run("variable section merge", func(t *testing.T) {
		ret, merged, err := C_scriptMergeVariableSection(
			scriptVariableSection(10, 5),
			scriptVariableSection(70, 6),
		)
		require.NoError(t, err)
		require.Equal(t, 1, ret)
		require.Greater(t, len(merged), 300)
		require.Equal(t, uint32(70), binary.LittleEndian.Uint32(merged[:4]))
	})

	t.Run("code section merge", func(t *testing.T) {
		ret, merged, err := C_scriptMergeCodeSections(
			scriptCodeSection(10, "FirstFunction"),
			scriptCodeSection(70, "SecondFunction"),
		)
		require.NoError(t, err)
		require.Equal(t, 1, ret)
		require.Equal(t, "CODE", string(merged[:4]))
		require.Equal(t, uint32(4), binary.LittleEndian.Uint32(merged[4:8]))
		require.Contains(t, string(merged), "FirstFunction")
		require.Contains(t, string(merged), "SecondFunction")
	})

	t.Run("complete container merge", func(t *testing.T) {
	ret, merged, err := C_scriptMergeContainers(
			scriptContainer(10, "FirstFunction", "alpha", "beta"),
			scriptContainer(70, "SecondFunction", "gamma"),
	)
	require.NoError(t, err)
	require.Equal(t, 4, ret)
		require.True(t, bytes.HasPrefix(merged, []byte("SCRIPT03STRG")))
		require.True(t, bytes.HasSuffix(merged, []byte("DONE")))
		require.Contains(t, string(merged), "alpha")
		require.Contains(t, string(merged), "gamma")
	})
}
