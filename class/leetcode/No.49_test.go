package leetcode

import "testing"

func Test49(t *testing.T) {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(data)
	t.Log(ans)
}
