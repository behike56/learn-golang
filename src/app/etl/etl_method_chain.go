package etl

//
// ===== メソッドチェーン実装（ジェネリクス使用, Go 1.18+） =====
//

type Seq[T any] struct {
	items []T
}

func NewSeq[T any](items []T) Seq[T] {
	return Seq[T]{items: items}
}

func (s Seq[T]) Items() []T {
	return s.items
}

func (s Seq[T]) Filter(pred func(T) bool) Seq[T] {
	var r []T
	for _, v := range s.items {
		if pred(v) {
			r = append(r, v)
		}
	}
	return Seq[T]{items: r}
}

func (s Seq[T]) Map(f func(T) T) Seq[T] {
	r := make([]T, 0, len(s.items))
	for _, v := range s.items {
		r = append(r, f(v))
	}
	return Seq[T]{items: r}
}

func (s Seq[T]) Reduce(acc T, f func(T, T) T) T {
	for _, v := range s.items {
		acc = f(acc, v)
	}
	return acc
}

// ★ 型変換したい場合はスタンドアロン関数で
func MapTo[T any, U any](s Seq[T], f func(T) U) Seq[U] {
	r := make([]U, 0, len(s.items))
	for _, v := range s.items {
		r = append(r, f(v))
	}
	return Seq[U]{items: r}
}
