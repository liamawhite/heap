package main

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// O(n + log(n)) => O(n)
func Heapify[N Number](input []N) *MaxHeap[N] {
	// Create the basic struct
	heap := &MaxHeap[N]{elems: make([]N, len(input)+1)}

	// Copy the slice and put the 0-th element at the end (this is why its len + 1)
	copy(heap.elems, input)
	heap.elems[len(heap.elems)-1], heap.elems[0] = heap.elems[0], heap.elems[len(heap.elems)-1]

	// Go through each node in reverse bubbling down
	// We can skip the first half because they are all the leaf nodes and will be handled by their parents
	// We don't need to do 0 because its not actually part of out heap
	for i := (heap.Size() + 1) / 2; i > 0; i-- {
		curr := i
		// While children are in bounds (i.e. exist) AND bigger than the current elem.
		for left(curr) <= heap.Size() &&
			right(curr) <= heap.Size() &&
			(heap.elems[left(curr)] > heap.elems[curr] || heap.elems[right(curr)] > heap.elems[curr]) {
			// Swap with the larger child
			if heap.elems[left(curr)] > heap.elems[right(curr)] {
				heap.elems[left(curr)], heap.elems[curr] = heap.elems[curr], heap.elems[left(curr)]
				curr = left(curr)
				continue
			}
			heap.elems[right(curr)], heap.elems[curr] = heap.elems[curr], heap.elems[right(curr)]
			curr = right(curr)
		}
	}
	return heap
}

type MaxHeap[N Number] struct {
	elems []N
}

// O(log n)
func (h *MaxHeap[N]) Pop() N {
	// Swap root with last element
	h.elems[len(h.elems)-1], h.elems[1] = h.elems[1], h.elems[len(h.elems)-1]

	// Pop it off
	var res N
	res, h.elems = h.elems[h.Size()], h.elems[:h.Size()]

	// Bubble root down
	curr := 1
	// While children are in bounds (i.e. exist) AND bigger than the current elem.
	for (left(curr) <= h.Size() && h.elems[left(curr)] > h.elems[curr]) ||
		(right(curr) <= h.Size() && h.elems[right(curr)] > h.elems[curr]) {

		// If no right child use left, has to be at least a left child so we dont need to check reverse.
		if right(curr) > h.Size() {
			h.elems[left(curr)], h.elems[curr] = h.elems[curr], h.elems[left(curr)]
			curr = left(curr)
			continue
		}

		// Swap with the larger child
		if h.elems[left(curr)] > h.elems[right(curr)] {
			h.elems[left(curr)], h.elems[curr] = h.elems[curr], h.elems[left(curr)]
			curr = left(curr)
			continue
		}
		h.elems[right(curr)], h.elems[curr] = h.elems[curr], h.elems[right(curr)]
		curr = right(curr)
	}
	return res
}

// O(log n * elements)
func (h *MaxHeap[N]) Push(elements ...N) {
	for _, element := range elements {

		// Push to the end
		h.elems = append(h.elems, element)

		// Bubble up!
		curr := h.Size()
		for parent(curr) > 0 && h.elems[parent(curr)] < h.elems[curr] {
			h.elems[parent(curr)], h.elems[curr] = h.elems[curr], h.elems[parent(curr)]
			curr = parent(curr)
		}
	}
}

func (h *MaxHeap[N]) Size() int {
	return len(h.elems) - 1
}

func left(parent int) int {
	return parent * 2
}

func right(parent int) int {
	return left(parent) + 1
}

func parent(child int) int {
	return child / 2
}
