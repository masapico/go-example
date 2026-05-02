package errors

import (
	stderrors "errors"
	"io"
	"strings"
	"testing"
)

// TestParsePositiveInt は正常系と代表的なエラーケースを表形式で確認する。
func TestParsePositiveInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "valid", input: " 42 ", want: 42},
		{name: "empty", input: "", wantErr: true},
		{name: "negative", input: "-1", wantErr: true},
		{name: "not number", input: "abc", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePositiveInt(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatal("ParsePositiveInt() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("ParsePositiveInt() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("ParsePositiveInt() = %d, want %d", got, tt.want)
			}
		})
	}
}

// TestParsePositiveIntEmptySentinel は空入力で番兵エラーを返すことを確認する。
func TestParsePositiveIntEmptySentinel(t *testing.T) {
	_, err := ParsePositiveInt(" ")
	if !stderrors.Is(err, ErrEmptyInput) {
		t.Fatalf("error = %v, want ErrEmptyInput", err)
	}
}

// TestReadAllAndClose は読み取った内容が文字列として返ることを確認する。
func TestReadAllAndClose(t *testing.T) {
	got, err := ReadAllAndClose(io.NopCloser(strings.NewReader("hello")))
	if err != nil {
		t.Fatalf("ReadAllAndClose() error = %v", err)
	}
	if got != "hello" {
		t.Fatalf("ReadAllAndClose() = %q, want hello", got)
	}
}
