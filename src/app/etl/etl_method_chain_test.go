package etl

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSeq_Filter_Map_Items_Reduce_BasicFlow(t *testing.T) {
	// 入力データを準備
	s := NewSeq([]int{1, 2, 3, 4, 5})

	// 検証1: Filter→Map→Items のチェーンが意図した順序で適用されること
	//  - 奇数のみ残す（1,3,5）
	//  - それらを平方する（1,9,25）
	//  - Items() が最終結果のスライスを返す
	got := s.
		Filter(func(v int) bool { return v%2 == 1 }).
		Map(func(v int) int { return v * v }).
		Items()

	want := []int{1, 9, 25}
	// 検証: 期待値と完全一致（順序・要素が同じ）
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Items() = %v, want %v", got, want)
	}

	// 検証2: Reduce が全要素に対して累積関数を適用し、初期値から正しい累計を返す
	//  - 1+2+3+4+5 = 15
	sum := s.Reduce(0, func(acc, v int) int { return acc + v })
	if sum != 15 {
		t.Fatalf("Reduce(sum) = %d, want 15", sum)
	}
}

func TestSeq_MapTo_TypeChange(t *testing.T) {
	// 検証: MapTo で「要素型が変わる」変換ができること（int -> string）
	s := NewSeq([]int{10, 20, 30})

	strSeq := MapTo[int, string](s, func(v int) string { return strconv.Itoa(v) })
	got := strSeq.Items()
	want := []string{"10", "20", "30"}

	// 検証: 変換後の要素列が期待どおりである（要素の順序も維持される）
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("MapTo Items() = %v, want %v", got, want)
	}
}

func TestSeq_Empty_Reduce_ReturnsAccumulator(t *testing.T) {
	// 検証: 空シーケンスに対する Reduce は、関数が一度も呼ばれずに初期値をそのまま返す
	empty := NewSeq([]int{})
	callCount := 0
	got := empty.Reduce(42, func(acc, v int) int {
		callCount++
		return acc + v
	})
	if got != 42 {
		t.Fatalf("Reduce on empty = %d, want 42", got)
	}
	if callCount != 0 {
		t.Fatalf("Reduce function should not be called, but was called %d times", callCount)
	}
}

func TestSeq_Filter_NoMatch_ReturnEmpty(t *testing.T) {
	// 検証: Filter で条件に一致する要素がない場合は空スライスになる
	s := NewSeq([]int{2, 4, 6})
	got := s.
		Filter(func(v int) bool { return v%2 == 1 }).
		Items()
	if len(got) != 0 {
		t.Fatalf("Filter with no match = %v, want empty", got)
	}
}
