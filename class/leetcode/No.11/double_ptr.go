package No_11

func maxArea(height []int) int {
	N := len(height)
	l, r := 0, N-1
	var area int
	for l < r {
		if height[l] > height[r] {
			area = max(area, (r-l)*height[r])
			r--
		} else {
			area = max(area, (r-l)*height[l])
			l++
		}
	}
	return area
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
