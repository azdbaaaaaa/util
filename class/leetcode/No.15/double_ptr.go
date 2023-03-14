package No_15

import "sort"

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0, 0)
	n := len(nums)
	if n <= 2 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l = l + 1
				}
				for l < r && nums[r] == nums[r-1] {
					r = r - 1
				}
				l = l + 1
				r = r - 1
			} else if sum > 0 {
				r = r - 1
			} else {
				l = l + 1
			}
		}
	}
	return ans
}
