package logic_test

import (
	"reflect"
	"testing"
	"tyranno/backend/utils/logic"
)

func TestSplitStrings(t *testing.T) {
	tests := []struct {
		input  string
		output [][]string
	}{
		{"++ xxx yyy zzz --", [][]string{{"xxx", "yyy", "zzz"}}},
		{"++ aaa bbb ccc -- ++ xxx yyy zzz --", [][]string{{"aaa", "bbb", "ccc"}, {"xxx", "yyy", "zzz"}}},
		{"++ a --", [][]string{{"a"}}},
		{"++ one two three -- ++ four five --", [][]string{{"one", "two", "three"}, {"four", "five"}}},
	}

	for _, test := range tests {
		result, err := logic.SplitStrings(test.input)
		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
			continue
		}

		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("入力: %s, 期待される出力: %v, 実際の出力: %v", test.input, test.output, result)
		}
	}
}
