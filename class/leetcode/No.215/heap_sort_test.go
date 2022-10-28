package No_215

import (
	"testing"
)

func TestLargestKth(t *testing.T) {
	nums := []int{2,1}
	k := 2
	//data := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	ans := findKthLargest(nums, k)
	t.Log(ans)
}
