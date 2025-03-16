package algos

import (
	"golang.org/x/exp/constraints"
	"sync"
)

func MergeSort[T constraints.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) >> 1
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge[T constraints.Ordered](left, right []T) []T {
	arr := make([]T, len(left)+len(right))

	l, r := 0, 0
	for i := 0; i < len(arr); i++ {
		if l >= len(left) {
			arr[i] = right[r]
			r++
		} else if r >= len(right) {
			arr[i] = left[l]
			l++
		} else if left[l] < right[r] {
			arr[i] = left[l]
			l++
		} else {
			arr[i] = right[r]
			r++
		}
	}

	return arr
}

func MergeSortParallel[T constraints.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) >> 1
	var left, right []T

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		left = MergeSortParallel(arr[:mid])
	}()

	go func() {
		defer wg.Done()
		right = MergeSortParallel(arr[mid:])
	}()

	wg.Wait()

	return merge(left, right)
}

// refer: https://github.com/teivah/golang-parallel-mergesort, https://teivah.medium.com/parallel-merge-sort-in-go-fe14c1bc006
var maxSize = 1 << 11

func MergeSortParallelV2[T constraints.Ordered](arr []T) []T {
	arr_len := len(arr)
	if arr_len < 2 {
		return arr
	}

	mid := arr_len >> 1
	if arr_len <= maxSize {
		return MergeSort(arr)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	left := make([]T, mid)
	right := make([]T, arr_len-mid)

	go func() {
		defer wg.Done()
		left = MergeSortParallelV2(arr[:mid])
	}()

	go func() {
		defer wg.Done()
		right = MergeSortParallelV2(arr[mid:])
	}()

	wg.Wait()

	return merge(left, right)
}

func MergeSortParallelV3[T constraints.Ordered](arr []T) []T {
	// creates lesser go routines, reducing GC load (less memory allocation)
	arr_len := len(arr)
	if arr_len < 2 {
		return arr
	}

	mid := arr_len >> 1
	if arr_len <= maxSize {
		return MergeSort(arr)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	var left, right []T

	go func() {
		defer wg.Done()
		left = MergeSortParallelV3(arr[:mid])
	}()

	right = MergeSortParallelV3(arr[mid:])

	wg.Wait()

	return merge(left, right)
}

// refer: https://www.codershaven.com/posts/exploring-sorting-algorithms-in-go/
func MergeSortIterative[T constraints.Ordered](arr []T) []T {
	length := len(arr)
	buffer := make([]T, length)

	// segment size
	for size := 1; size < length; size <<= 1 {
		for segment := 0; segment < length; segment += size << 1 {
			end := segment + size << 1
			if end > length {
				end = length
			}

			l, r := segment, segment + size
			for i := segment; i < end; i++ {
				if l >= segment + size {
					buffer[i] = arr[r]
					r++
				} else if r >= end {
					buffer[i] = arr[l]
					l++
				} else if arr[l] < arr[r] {
					buffer[i] = arr[l]
					l++
				} else {
					buffer[i] = arr[r]
					r++
				}
			}
		}
		buffer, arr = arr, buffer
	}

	return arr
}
