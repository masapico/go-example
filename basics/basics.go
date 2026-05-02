package basics

import "sort"

// Sum は整数スライスの合計値を返します。
func Sum(nums []int) int {
	total := 0
	// range ではインデックスと値を取り出せるが、ここでは値だけを使う。
	for _, n := range nums {
		total += n
	}
	return total
}

// CountWords は単語ごとの出現回数を map にまとめます。
func CountWords(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		// 空文字は単語として数えない。
		if word == "" {
			continue
		}
		// map のゼロ値を利用して、初登場の単語も 0 + 1 として扱える。
		counts[word]++
	}
	return counts
}

// TopWords は出現回数が多い順に、最大 limit 件の単語を返します。
func TopWords(counts map[string]int, limit int) []string {
	if limit <= 0 {
		return nil
	}

	// map は順序を持たないため、キーをスライスに取り出してから並べ替える。
	words := make([]string, 0, len(counts))
	for word := range counts {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		// 回数が同じ場合は文字列順にして、結果が毎回同じになるようにする。
		if counts[words[i]] == counts[words[j]] {
			return words[i] < words[j]
		}
		return counts[words[i]] > counts[words[j]]
	})

	// limit が単語数を超えてもスライス範囲外にならないように調整する。
	if limit > len(words) {
		limit = len(words)
	}
	return words[:limit]
}
