package algos

import (
	"golang.org/x/exp/constraints"
)

const (
	// Threshold for switching to insertion sort

	// parallel limit
	ParallelLimit = 1000
)

func LomutoPartition[T constraints.Ordered](arr []T, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSort[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	low, high := 0, len(arr)-1
	pivotIndex := LomutoPartition(arr, low, high)
	QuickSort(arr[low:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
	return arr
}

func QuickSortParallel[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	low, high := 0, len(arr)-1
	pivotIndex := LomutoPartition(arr, low, high)
	if high-low < ParallelLimit {
		QuickSortParallel(arr[low:pivotIndex])
		QuickSortParallel(arr[pivotIndex+1:])
	} else {
		var done = make(chan struct{}, 2)
		go func() {
			QuickSortParallel(arr[low:pivotIndex])
			done <- struct{}{}
		}()
		go func() {
			QuickSortParallel(arr[pivotIndex+1:])
			done <- struct{}{}
		}()
		for i := 0; i < 2; i++ {
			<-done
		}
	}
	return arr
}

func QuickSortHoare[T constraints.Ordered](arr []T, low, high int) []T {
	if low >= high {
		return arr
	}
	pivot := arr[low]
	i, j := low, high
	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}
		for arr[j] > pivot && j > low {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// Swap the pivot element with the element at index j
	arr[low], arr[j] = arr[j], arr[low]

	QuickSortHoare(arr, low, j-1)
	QuickSortHoare(arr, j+1, high)
	return arr
}

func QuickV2[T constraints.Ordered](arr []T) []T {
	arr = QuickSortHoare(arr, 0, len(arr)-1)
	return arr
}
