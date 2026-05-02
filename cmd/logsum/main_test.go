package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

// TestCountLevels はログレベルを大文字小文字に関係なく集計できることを確認する。
func TestCountLevels(t *testing.T) {
	input := strings.NewReader(`
2026-05-02 INFO started
2026-05-02 ERROR failed
2026-05-02 info done
`)
	got, err := countLevels(input)
	if err != nil {
		t.Fatalf("countLevels() error = %v", err)
	}
	want := map[string]int{"INFO": 2, "ERROR": 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("countLevels() = %#v, want %#v", got, want)
	}
}

// TestCountLevelsInvalidLine は列数が足りないログ行をエラーにすることを確認する。
func TestCountLevelsInvalidLine(t *testing.T) {
	_, err := countLevels(strings.NewReader("broken"))
	if err == nil {
		t.Fatal("countLevels() error = nil, want error")
	}
}

// TestRunUsage は引数がない場合に使い方のエラーを返すことを確認する。
func TestRunUsage(t *testing.T) {
	err := run(nil, &bytes.Buffer{})
	if err == nil {
		t.Fatal("run() error = nil, want usage error")
	}
}
