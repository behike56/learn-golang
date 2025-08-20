package main

import (
	"fmt"

	"github.com/behike56/learn-golang/app/etl"
)

func main() {
	sum := etl.NewSeq([]int{1, 2, 3, 4, 5}).
		Filter(func(n int) bool { return n%2 == 0 }).
		Map(func(n int) int { return n * 2 }).
		Reduce(0, func(acc, v int) int { return acc + v })
	fmt.Println(sum)
}
