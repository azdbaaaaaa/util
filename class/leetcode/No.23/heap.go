package No_23

func heap(lists []*ListNode) *ListNode {
	type heapListNode
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

// 当子节点位置的值比父节点位置的值大的时候进行交换
func heapInsert(arr []int, idx int) {
	for arr[idx] > arr[(idx-1)/2] {
		swap(arr, idx, (idx-1)/2)
		idx = (idx - 1) / 2
	}
}

func heapify(arr []int, idx int, heapSize int) {
	left := 2*idx + 1

	for left < heapSize {
		// 找到左右孩子中比较大的一个
		largest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}

		if arr[idx] >= arr[largest] {
			break
		}
		swap(arr, idx, largest)
		idx = largest
		left = 2*idx + 1
	}
}

func swap(arr []int, a, b int) {
	t := arr[a]
	arr[a] = arr[b]
	arr[b] = t
}
