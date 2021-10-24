package main

import (
	bubble "algorithms/internal/sorting/bubble_sort"
	merge "algorithms/internal/sorting/merge_sort"
	selection "algorithms/internal/sorting/selection_sort"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	data := generateDataSet(100000)
	data2 := make([]int, 100000)
	data3 := make([]int, 100000)
	copy(data2, data)
	copy(data3, data)

	t1 := time.Now()
	r := selection.SelctionSort(data)
	show(r)
	fmt.Printf("\nt1 take: %v\n", time.Since(t1))
	fmt.Println()

	t2 := time.Now()
	r2 := bubble.BubbleSort(data2)
	show(r2)
	fmt.Printf("\nt2 take: %v\n", time.Since(t2))

	fmt.Println()
	t3 := time.Now()
	r3 := merge.MergeSort(data3)
	show(r3)
	fmt.Printf("\nt3 take: %v\n", time.Since(t3))
}

func show(data []int) {
	for _, _ = range data {
		//fmt.Printf("%v ", val)
	}
}

func generateDataSet(size int) []int {
	var r []int
	for i := 1; i <= size; i++ {
		r = append(r, rand.Intn(100))
	}
	return r
}
