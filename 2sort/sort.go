package sort

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// quickSort divide and conquer algorithm
// The array is repeatedly partitioned around a pivot element.
// After partitioning, the pivot is in its correct position in the sorted array.
// The process is repeated recursively for the left and right subarrays until the entire array is sorted.
func quickSort(arr []int) []int {
	qs(arr, 0, len(arr)-1)
	return arr
}

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	// partitioning incoming array
	pivIdx := partition(arr, lo, hi)
	// sorting everything left of the pivot
	qs(arr, lo, pivIdx-1)
	//sorting everything right of the pivot
	qs(arr, pivIdx+1, hi)

}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi] // chose random pivot
	pivIdx := lo - 1 // this is the position where the next smaller element will be placed
	for i := lo; i < hi; i++ {
		if arr[i] < pivot {
			pivIdx++
			arr[pivIdx], arr[i] = arr[i], arr[pivIdx] // group together small elements
		}
	}
	pivIdx++
	// pivot is swapped with the element at pivotIdx + 1 placing it in its final sorted position
	arr[hi], arr[pivIdx] = arr[pivIdx], pivot

	return pivIdx
}
