package main

import (
	"fmt"
)

func main() {
	maxHeap := Heapify([]int{1, 2, 3, 4, 5, 6})
	maxHeap.Push(7, 8, 9)
	for maxHeap.Size() > 0 {
		fmt.Println(maxHeap.Pop())
	}
}
