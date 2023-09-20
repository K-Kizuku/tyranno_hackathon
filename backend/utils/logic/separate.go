package logic

import "strings"

func SplitStrings(s string) ([][]string, error) {
	// "++" と "--" で文字列を分割します
	groups := strings.Split(s, "-- ++")

	var result [][]string

	for _, group := range groups {
		// 各グループの前後の "++" と "--" を取り除きます
		trimmed := strings.TrimPrefix(strings.TrimSuffix(group, "--"), "++")

		// 文字列をスペースで分割します
		parts := strings.Split(strings.TrimSpace(trimmed), " ")
		result = append(result, parts)
	}

	return result, nil
}
