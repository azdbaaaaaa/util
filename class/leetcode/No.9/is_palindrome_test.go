package No_9

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	data := 10
	//data := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	ans := isPalindrome(data)
	t.Log(ans)
}
