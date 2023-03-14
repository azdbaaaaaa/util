package No_56

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	merged := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if len(merged) == 0 {
			merged = append(merged, intervals[i])
			continue
		}
		if merged[len(merged)-1][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
		} else {
			merged[len(merged)-1][1] = max(intervals[i][1], merged[len(merged)-1][1])
		}
	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
