package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/learningsorting/algos"
)

func CreateArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func CheckSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func SortArray(arr []int, algo func([]int) []int, name string, inPlace bool) {
	start := time.Now()
	if inPlace {
		algo(arr)
	} else {
		arr = algo(arr)
	}
	end := time.Now()
	if !CheckSorted(arr) {
		fmt.Printf("%s failed to sort the array\n", name)
		fmt.Println("Array: ", arr)
	}
	fmt.Printf("%s took %s to sort %d elements\n", name, end.Sub(start), len(arr))
}

func main() {
	size := 10000
	SortArray(CreateArray(size), algos.SelectionSort, "Selection Sort", true)
	SortArray(CreateArray(size), algos.BubbleSort, "Bubble Sort", true)
	SortArray(CreateArray(size), algos.InsertionSort, "Insertion Sort", true)
	SortArray(CreateArray(size), algos.MergeSort, "Merge Sort", false)
	SortArray(CreateArray(size), algos.MergeSortParallel, "Merge Sort Parallel", false)
	SortArray(CreateArray(size), algos.MergeSortParallelV2, "Merge Sort Parallel V2", false)
	SortArray(CreateArray(size), algos.MergeSortParallelV3, "Merge Sort Parallel V3", false)
	SortArray(CreateArray(size), algos.MergeSortIterative, "Merge Sort Iterative", false)
	SortArray(CreateArray(size), algos.QuickSort, "Quick Sort", false)
	SortArray(CreateArray(size), algos.QuickSortParallel, "Quick Sort Parallel", false)
	SortArray(CreateArray(size), algos.QuickV2, "Quick Sort Hoare", false)	
}
