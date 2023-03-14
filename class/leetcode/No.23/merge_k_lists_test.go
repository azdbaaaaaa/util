package No_23

import (
	"testing"
)

func Test23(t *testing.T) {
	var data []*ListNode
	{
		data = []*ListNode{{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		}, {
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}, {
			Val: 2,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
		}
	}

	ans := mergeKLists2(data)

	for ans != nil {
		t.Log(ans.Val)
		ans = ans.Next
	}
}
