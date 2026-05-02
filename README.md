# go-example

Go の基礎固め用リポジトリです。他言語経験がある人が Go らしい書き方を短いサンプルとテストで確認できるように、単元ごとに小さなパッケージを置いています。

## 進め方

1. 各ディレクトリの `.go` と `_test.go` を読む
2. `go test ./...` を実行する
3. テストケースを1つ追加する
4. 実装を少し変更して、失敗と修正を確認する

## 単元

- `basics`: 変数、関数、制御構文、スライス、マップ
- `types`: 構造体、メソッド、ポインタ、インターフェース
- `errors`: エラー処理、`defer`、標準ライブラリ
- `data`: ファイル I/O、JSON、日時、文字列処理
- `concurrency`: goroutine、channel、context、sync
- `cmd/logsum`: 小課題。ログ行を集計する CLI

## よく使うコマンド

```powershell
go test ./...
go test -race ./...
go run ./cmd/logsum ./testdata/logs.txt
```

`go test -race ./...` は cgo と C コンパイラが必要です。Windows で `go: -race requires cgo` と出る場合は、まず通常の `go test ./...` で進めてください。

## 学習メモ

- Go は例外ではなく `error` を値として返します。
- インターフェースは実装側に宣言を書かず、必要な振る舞いだけで受け取ります。
- 並行処理は goroutine だけでなく、終了条件とキャンセルを必ず考えます。
- テストは標準の `testing` パッケージだけでもかなり書けます。
