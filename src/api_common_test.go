package opennox

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestWriteJSONResp(t *testing.T) {
	w := httptest.NewRecorder()
	v := map[string]any{"foo": "bar", "num": 42}
	writeJSONResp(w, v)

	resp := w.Result()
	if ct := resp.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", ct)
	}
	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if out["foo"] != "bar" {
		t.Errorf("foo = %v, want bar", out["foo"])
	}
	// json numbers decode as float64
	if out["num"] != float64(42) {
		t.Errorf("num = %v, want 42", out["num"])
	}
}
