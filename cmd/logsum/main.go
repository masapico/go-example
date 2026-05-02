package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	// main では入出力と終了コードだけを扱い、実際の処理は run に分ける。
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run はコマンドライン引数を検証し、ログレベルごとの件数を出力します。
func run(args []string, out io.Writer) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: logsum <path>")
	}

	file, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("open log file: %w", err)
	}
	defer file.Close()

	counts, err := countLevels(file)
	if err != nil {
		return err
	}

	// map のままだと出力順が不定なので、キーをソートしてから出力する。
	levels := make([]string, 0, len(counts))
	for level := range counts {
		levels = append(levels, level)
	}
	sort.Strings(levels)

	for _, level := range levels {
		fmt.Fprintf(out, "%s %d\n", level, counts[level])
	}
	return nil
}

// countLevels はログを読み取り、2列目のログレベルごとの件数を数えます。
func countLevels(r io.Reader) (map[string]int, error) {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// 空行はログとして扱わず読み飛ばす。
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			return nil, fmt.Errorf("invalid log line: %q", line)
		}
		// INFO と info を同じログレベルとして数えるため大文字にそろえる。
		counts[strings.ToUpper(fields[1])]++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan log: %w", err)
	}
	return counts, nil
}
