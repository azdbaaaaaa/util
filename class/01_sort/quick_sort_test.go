package sort

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int{1, 2, 3, 4}
	//quickSort(data)
	t.Log(findUnsortedSubarray(data))
	t.Log(data)
}

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	n := len(nums)
	numsCopy := make([]int, 0, n)
	numsCopy = append(numsCopy, nums...)

	sort.Ints(numsCopy)
	l := 0
	r := n - 1
	for l < r {
		if numsCopy[l] == nums[l] {
			l++
			continue
		}

		if numsCopy[r] == nums[r] {
			r--
			continue
		}
		break
	}
	return r - l + 1
}
//func findUnsortedSubarray(nums []int) int {
//	if sort.IntsAreSorted(nums) {
//		return 0
//	}
//	numsSorted := append([]int(nil), nums...)
//	sort.Ints(numsSorted)
//	left, right := 0, len(nums)-1
//	for nums[left] == numsSorted[left] {
//		left++
//	}
//	for nums[right] == numsSorted[right] {
//		right--
//	}
//	return right - left + 1
//}
