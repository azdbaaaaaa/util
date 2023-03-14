package sort

func insertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j <= i; j++ {
			if arr[i] <= arr[j] {
				swap(arr, j , i)
			}
		}
	}
	return
}

// [5, 4, 3, 2, 1, 4]
// [4, 5, 3, 2, 1, 4]
// [3, 4, 5, 2, 1, 4]
