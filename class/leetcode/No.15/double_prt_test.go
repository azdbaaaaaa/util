package No_15

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	data := []int{3,2,3}
	ans := threeSum(data)
	t.Log(ans)


}

 type ListNode struct {
	     Val int
	     Next *ListNode
	 }
func hasCycle(head *ListNode) bool {
	m := make(map[*listNode]struct{})
	cur := head
	for cur != nil {
		if _, ok := m[cur]; ok {
			return true
		}
		m[cur]=struct{}{}
		cur = cur.Next
	}
	return false
}
