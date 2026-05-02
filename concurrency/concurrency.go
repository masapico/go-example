package concurrency

import (
	"context"
	"sync"
)

// SquareAll は各整数の二乗を並行に計算し、入力と同じ順序で返します。
func SquareAll(ctx context.Context, nums []int) ([]int, error) {
	out := make([]int, len(nums))
	// 最初のエラーだけ受け取れればよいので、バッファを 1 にする。
	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	for i, n := range nums {
		// goroutine がループ変数を共有しないよう、各反復で値を固定する。
		i, n := i, n
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				// 複数 goroutine から送信されてもブロックしないようにする。
				select {
				case errCh <- ctx.Err():
				default:
				}
			default:
				out[i] = n * n
			}
		}()
	}

	wg.Wait()
	// キャンセルが発生していればエラー、なければ計算結果を返す。
	select {
	case err := <-errCh:
		return nil, err
	default:
		return out, nil
	}
}

// Merge は複数の入力チャネルを 1 つの出力チャネルにまとめます。
func Merge[T any](ctx context.Context, inputs ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup

	for _, input := range inputs {
		// goroutine ごとに読み取るチャネルを固定する。
		input := input
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case value, ok := <-input:
					if !ok {
						return
					}
					// 送信待ちの間にもキャンセルを受け付ける。
					select {
					case <-ctx.Done():
						return
					case out <- value:
					}
				}
			}
		}()
	}

	go func() {
		// すべての入力読み取りが終わったら、出力チャネルを閉じる。
		wg.Wait()
		close(out)
	}()

	return out
}
