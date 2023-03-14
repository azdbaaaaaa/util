package No_11

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	data := []int{1,8,6,2,5,4,8,3,7}
	ans := maxArea(data)
	t.Log(ans)
}
