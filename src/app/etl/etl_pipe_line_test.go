package etl

import (
	"reflect"
	"strconv"
	"testing"
)

// テスト補助: チャネルをすべて読み切ってスライス化する
// 検証観点: 出力チャネルが close されること、全要素を取り切れること
func collect[T any](ch <-chan T) []T {
	var out []T
	for v := range ch {
		out = append(out, v)
	}
	return out
}

func TestPipeline_FilterThenMap_SquaredOdds(t *testing.T) {
	// 入力の生成（1..5）
	in := gen(1, 2, 3, 4, 5)

	// 検証1: filter が条件（奇数のみ）に合致する要素だけを通す
	odds := filter(in, func(n int) bool { return n%2 == 1 })

	// 検証2: mapChan が通過要素に関数（平方）を適用する
	squared := mapChan[int, int](odds, func(n int) int { return n * n })

	// 検証3: チャネルのパイプラインが正しく連結され、期待値が得られる
	got := collect(squared)
	want := []int{1, 9, 25}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("pipeline result = %v, want %v", got, want)
	}
}

func TestMapChan_IntToString(t *testing.T) {
	// 検証: mapChan によって要素型の変換（int -> string）ができる
	in := gen(7, 8, 9)
	strs := mapChan[int, string](in, func(n int) string { return strconv.Itoa(n) })

	// 検証: 変換された要素が順序どおりに出力される
	got := collect(strs)
	want := []string{"7", "8", "9"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("mapChan int->string = %v, want %v", got, want)
	}
}

func TestGen_NoArgs_Empty(t *testing.T) {
	// 検証: 引数なしの gen はすぐに close され、要素は 0 件
	in := gen()
	got := collect(in)
	if len(got) != 0 {
		t.Fatalf("gen() = %v, want empty", got)
	}
}

func TestFilter_NoPass_AllFilteredOut(t *testing.T) {
	// 検証: filter の述語に一致しない場合、出力は空になる
	in := gen(2, 4, 6)
	out := filter(in, func(n int) bool { return n%2 == 1 })
	got := collect(out)
	if len(got) != 0 {
		t.Fatalf("filter with no pass = %v, want empty", got)
	}
}
