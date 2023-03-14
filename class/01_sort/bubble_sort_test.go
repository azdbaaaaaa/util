package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	data := []int{1, 5, 7, 4, 2, 1, 9, 10, 8}
	bubbleSort(data)
	t.Log(data)
}
