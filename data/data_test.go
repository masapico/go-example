package data

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestEncodeDecodeEvent は Event を JSON 化して元に戻せることを確認する。
func TestEncodeDecodeEvent(t *testing.T) {
	at := time.Date(2026, 5, 2, 12, 0, 0, 0, time.UTC)
	event := Event{Name: "learn go", At: at}

	encoded, err := EncodeEvent(event)
	if err != nil {
		t.Fatalf("EncodeEvent() error = %v", err)
	}

	got, err := DecodeEvent(encoded)
	if err != nil {
		t.Fatalf("DecodeEvent() error = %v", err)
	}
	if got.Name != event.Name || !got.At.Equal(event.At) {
		t.Fatalf("DecodeEvent() = %#v, want %#v", got, event)
	}
}

// TestDecodeEventRequiresName は name が必須項目として検証されることを確認する。
func TestDecodeEventRequiresName(t *testing.T) {
	_, err := DecodeEvent([]byte(`{"name":"","at":"2026-05-02T12:00:00Z"}`))
	if err == nil {
		t.Fatal("DecodeEvent() error = nil, want error")
	}
}

// TestLoadEvent はファイルから Event を読み込めることを確認する。
func TestLoadEvent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "event.json")
	if err := os.WriteFile(path, []byte(`{"name":"learn go","at":"2026-05-02T12:00:00Z"}`), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	got, err := LoadEvent(path)
	if err != nil {
		t.Fatalf("LoadEvent() error = %v", err)
	}
	if got.Name != "learn go" {
		t.Fatalf("LoadEvent().Name = %q, want learn go", got.Name)
	}
}
