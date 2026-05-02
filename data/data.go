package data

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Event は JSON に保存するイベント情報を表します。
type Event struct {
	// 構造体タグで JSON のフィールド名を指定する。
	Name string    `json:"name"`
	At   time.Time `json:"at"`
}

// EncodeEvent は Event を読みやすいインデント付き JSON に変換します。
func EncodeEvent(event Event) ([]byte, error) {
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		// %w で元のエラーを包むと、呼び出し側で errors.Is/As が使える。
		return nil, fmt.Errorf("encode event: %w", err)
	}
	return data, nil
}

// DecodeEvent は JSON から Event を復元し、必須項目を検証します。
func DecodeEvent(data []byte) (Event, error) {
	var event Event
	if err := json.Unmarshal(data, &event); err != nil {
		return Event{}, fmt.Errorf("decode event: %w", err)
	}
	// TrimSpace で空白だけの名前も未入力として扱う。
	if strings.TrimSpace(event.Name) == "" {
		return Event{}, fmt.Errorf("decode event: name is required")
	}
	return event, nil
}

// LoadEvent はファイルから JSON を読み込み、Event として返します。
func LoadEvent(path string) (Event, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Event{}, fmt.Errorf("load event: %w", err)
	}
	return DecodeEvent(data)
}
