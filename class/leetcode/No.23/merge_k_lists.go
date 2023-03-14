package No_23

type ListNode struct {
	Val  int
	Next *ListNode
}

// 顺序依次合并
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = merge(ans, lists[i])
	}
	return ans
}

// 合并过程采用归并
func merge(n1, n2 *ListNode) *ListNode {
	ans := &ListNode{}
	cur := ans
	for n1 != nil && n2 != nil {
		if n1.Val <= n2.Val {
			cur.Next = n1
			n1 = n1.Next
			cur = cur.Next
		} else {
			cur.Next = n2
			n2 = n2.Next
			cur = cur.Next
		}
	}
	for n1 != nil {
		cur.Next = n1
		n1 = n1.Next
		cur = cur.Next
	}

	for n2 != nil {
		cur.Next = n2
		n2 = n2.Next
		cur = cur.Next
	}
	return ans.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return process(lists, 0, len(lists)-1)
}

func process(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := l + (r-l)>>1
	return merge(process(lists, l, mid), process(lists, mid+1, r))
}
