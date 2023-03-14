package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0, 0)
	pool := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		sorted := sortStr(strs[i])
		if v, ok := pool[sorted]; ok {
			v = append(v, strs[i])
			pool[sorted] = v
		} else {
			pool[sorted] = []string{strs[i]}
		}
	}
	for _, v := range pool {
		ans = append(ans, v)
	}
	return ans
}

func sortStr(s string) string {
	data := []byte(s)
	sort.Slice([]byte(data), func(i int, j int) bool {
		return data[i] > data[j]
	})
	return string(data)
}
