package main

import (
	"fmt"
	"github.com/shuaixinjie/interview/algorithm"
)

func main() {
	fmt.Println("hello,world!")

	var aa = []int{9, 8, 3, 4, 5, 3, 2, 1}
	algorithm.QuickSort(aa)

	fmt.Println(aa)
}
