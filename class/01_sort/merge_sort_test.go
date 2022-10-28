package sort

import "testing"

func TestMergeSort(t *testing.T) {
	data := []int{1, 5, 7, 4, 2, 1, 9, 10, 8}
	mergeSort(data)
	t.Log(data)
}
