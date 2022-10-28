package sort

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
	}
	return
}

func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
