package errors

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ErrEmptyInput は入力が空だったことを判定しやすくする番兵エラーです。
var ErrEmptyInput = errors.New("empty input")

// ParsePositiveInt は文字列を正の整数に変換します。
func ParsePositiveInt(input string) (int, error) {
	// 前後の空白を許容してから、空文字かどうかを確認する。
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, ErrEmptyInput
	}

	n, err := strconv.Atoi(input)
	if err != nil {
		// 変換エラーを包んで、どの処理で失敗したかを補足する。
		return 0, fmt.Errorf("parse positive int: %w", err)
	}
	if n <= 0 {
		return 0, fmt.Errorf("parse positive int: must be positive: %d", n)
	}
	return n, nil
}

// ReadAllAndClose はすべてのデータを読み取り、最後に必ず Close します。
func ReadAllAndClose(r io.ReadCloser) (_ string, err error) {
	defer func() {
		// 読み取りが成功していた場合だけ、Close のエラーを返す。
		if closeErr := r.Close(); err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	data, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("read all: %w", err)
	}
	return string(data), nil
}
