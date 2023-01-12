package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func randomSlice[N Number]() []N {
	res := []N{}
	min, max := 0, 1000000
	for i := 0; i < rand.Intn(max-min)+min; i++ {
		// good enough ¯\_(ツ)_/¯
		res = append(res, N(rand.Int()))
	}
	return res
}

func TestMaxHeap(t *testing.T) {
	// Randomly generated tests for each type
	t.Run("int", test[int])
	t.Run("int8", test[int8])
	t.Run("int16", test[int16])
	t.Run("int32", test[int32])
	t.Run("int64", test[int64])
	t.Run("uint", test[uint])
	t.Run("uint8", test[uint8])
	t.Run("uint16", test[uint16])
	t.Run("uint32", test[uint32])
	t.Run("uint64", test[uint64])
	t.Run("float32", test[float32])
	t.Run("float64", test[float64])
}

func test[N Number](t *testing.T) {
	init := randomSlice[N]()
	push := randomSlice[N]()

	// Build a sorted list to compare to
	want := append(init, push...)
	sort.Slice(want, func(i, j int) bool {
		return want[i] < want[j]
	})

	h := Heapify(init)
	h.Push(push...)

	got := make([]N, 0, len(want))
	for h.Size() > 0 {
		got = append(got, h.Pop())
	}
	reflect.DeepEqual(want, got)
}
