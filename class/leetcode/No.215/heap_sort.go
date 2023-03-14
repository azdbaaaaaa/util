package No_215

// https://leetcode.cn/problems/kth-largest-element-in-an-array/?favorite=2cktkvj
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	heapSize := len(nums)

	// 堆化数组
	for i := 0; i < heapSize; i++ {
		heapInsert(nums, i)
	}

	swap(nums, 0, heapSize-1)
	heapSize--

	for i := 0; i < k && heapSize > 0; i++ {
		heapify(nums, 0, heapSize)

		swap(nums, 0, heapSize-1)
		heapSize--
	}

	return nums[len(nums)-k]
}

// i位置左孩子的位置 2*i+1
// i位置右孩子的位置 2*i+2
// i位置父(i-1)/2

// idx位置上的值不停的与父位置的数比较，当大于父位置的数的时候进行交换，否则停止
func heapInsert(arr []int, idx int) {
	for arr[idx] > arr[(idx-1)/2] {
		swap(arr, idx, (idx-1)/2)
		idx = (idx - 1) / 2
	}
}

// idx指从哪个位置开始往下做heapify
// heapSize 管理堆大小
// idx节点不停的与左右孩子的最大值进行比较，当小于它的时候进行交换，否则停止
func heapify(arr []int, idx int, heapSize int) {
	if heapSize == 0 {
		return
	}
	// 左右孩子的下标
	left := 2*idx + 1

	// idx有孩子的时候进入for循环
	for left < heapSize {
		// 获取左右孩子的最大值
		largest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}
		// 如果idx节点比左右孩子节点的最大值还要大，就跳出循环
		if arr[idx] >= arr[largest] {
			break
		}
		// 否则就把父节点和左右孩子节点的最大值的节点交换
		// 并且重新设置idx为之前左右孩子节点的最大值的节点
		swap(arr, idx, largest)
		idx = largest
		left = 2*idx + 1
	}
	return
}

func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
