package etl

//
// ===== パイプライン実装 =====
//

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func filter(in <-chan int, pred func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if pred(v) {
				out <- v
			}
		}
	}()
	return out
}

func mapChan[T any, U any](in <-chan T, f func(T) U) <-chan U {
	out := make(chan U)
	go func() {
		defer close(out)
		for v := range in {
			out <- f(v)
		}
	}()
	return out
}
