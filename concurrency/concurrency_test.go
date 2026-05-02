package concurrency

import (
	"context"
	"reflect"
	"sort"
	"testing"
)

// TestSquareAll は並行処理後も入力と同じ順序で二乗結果が返ることを確認する。
func TestSquareAll(t *testing.T) {
	got, err := SquareAll(context.Background(), []int{1, 2, 3, 4})
	if err != nil {
		t.Fatalf("SquareAll() error = %v", err)
	}
	want := []int{1, 4, 9, 16}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SquareAll() = %#v, want %#v", got, want)
	}
}

// TestSquareAllCanceled はキャンセル済み context でエラーになることを確認する。
func TestSquareAllCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := SquareAll(ctx, []int{1})
	if err == nil {
		t.Fatal("SquareAll() error = nil, want cancellation error")
	}
}

// TestMerge は複数チャネルの値が 1 つの出力チャネルから受け取れることを確認する。
func TestMerge(t *testing.T) {
	ctx := context.Background()
	left := make(chan int, 2)
	right := make(chan int, 2)
	left <- 1
	left <- 2
	right <- 3
	right <- 4
	close(left)
	close(right)

	var got []int
	for value := range Merge(ctx, left, right) {
		got = append(got, value)
	}
	sort.Ints(got)

	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Merge() = %#v, want %#v", got, want)
	}
}
