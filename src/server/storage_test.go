package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEscapeKey(t *testing.T) {
	require.Equal(t, "simple", escapeKey("simple"))
	require.Equal(t, "hello%20world", escapeKey("hello world"))
	require.Equal(t, "a%2Fb", escapeKey("a/b"))
}

func TestTempStorage(t *testing.T) {
	s := &tempStorage{kv: make(map[string][]byte)}

	// Get non-existent key
	val, err := s.Get("nonexistent")
	require.NoError(t, err)
	require.Nil(t, val)

	// Set and get
	err = s.Set("key1", []byte("value1"))
	require.NoError(t, err)

	val, err = s.Get("key1")
	require.NoError(t, err)
	require.Equal(t, []byte("value1"), val)

	// Set empty value (should delete)
	err = s.Set("key1", []byte{})
	require.NoError(t, err)

	val, err = s.Get("key1")
	require.NoError(t, err)
	require.Nil(t, val)
}

func TestNewStore(t *testing.T) {
	s := newStore(t.TempDir())
	require.NotNil(t, s)
	require.NotNil(t, s.temp)
	require.NotNil(t, s.persist)
}

func TestStorageSession(t *testing.T) {
	s := newStore(t.TempDir())

	// Default session
	sess := s.Session("")
	require.NotNil(t, sess)

	sess2 := s.Session("default")
	require.NotNil(t, sess2)

	// Named session
	sess3 := s.Session("mysession")
	require.NotNil(t, sess3)

	// Same session should be returned
	sess4 := s.Session("mysession")
	require.Equal(t, sess3, sess4)
}

func TestStoragePersistent(t *testing.T) {
	s := newStore(t.TempDir())

	// Default persistent
	p := s.Persistent("")
	require.NotNil(t, p)

	p2 := s.Persistent("default")
	require.NotNil(t, p2)

	// Named persistent
	p3 := s.Persistent("mypersist")
	require.NotNil(t, p3)

	// Same persistent should be returned
	p4 := s.Persistent("mypersist")
	require.Equal(t, p3, p4)
}

func TestFileStorageKeyPath(t *testing.T) {
	fs := &fileStorage{dir: "/tmp/test", cache: make(map[string][]byte)}
	path := fs.keyPath("mykey")
	require.Contains(t, path, "mykey.json")
}
