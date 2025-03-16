package algos

import (
	"golang.org/x/exp/constraints"
)

func SelectionSort[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		mini := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[mini] {
				mini = j
			}
		}
		arr[i], arr[mini] = arr[mini], arr[i]
	}
	return arr
}

func BubbleSort[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}

func InsertionSort[T constraints.Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			continue
		}

		// Method 1:
		// pos, element := i, arr[i]
		// for pos > 0 && arr[pos-1] > element {
		// 	arr[pos] = arr[pos-1]
		// 	pos--
		// }
		// arr[pos] = element

		// Method 2:
		pos, end := 0, i
		// O(log n)
		for pos < end {
			mid := int(uint(pos+end) >> 1)
			if arr[mid] < arr[i] {
				pos = mid + 1
			} else {
				end = mid
			}	
		}

		temp := arr[i]
		copy(arr[pos+1:], arr[pos:i]) // O(n)
		arr[pos] = temp

	}
	return arr
}