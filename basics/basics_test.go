package basics

import (
	"reflect"
	"testing"
)

// TestSum はスライス内の整数が正しく合計されることを確認する。
func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3, 4})
	if got != 10 {
		t.Fatalf("Sum() = %d, want 10", got)
	}
}

// TestCountWords は空文字を除外しながら単語数を数えることを確認する。
func TestCountWords(t *testing.T) {
	got := CountWords([]string{"go", "test", "go", ""})
	want := map[string]int{"go": 2, "test": 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("CountWords() = %#v, want %#v", got, want)
	}
}

// TestTopWords は出現回数順と同数時の文字列順で並ぶことを確認する。
func TestTopWords(t *testing.T) {
	counts := map[string]int{"go": 3, "test": 2, "api": 2}
	got := TopWords(counts, 2)
	want := []string{"go", "api"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TopWords() = %#v, want %#v", got, want)
	}
}
