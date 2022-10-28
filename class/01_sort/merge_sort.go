package sort

// T(n) = 2*T(n/2)+O(1)
// 时间复杂度：O(n*log(n))
// 归并排序
// 每次取数组的中间值mid，然后递归处理左右2边的数据，使他们有序，然后合并左右2边的数据到一个临时数组中
// 合并的过程使用双指针，左边的数据小于等于右边的数据的时候，左边的数据copy到临时数组中，然后左指针向右移动
func mergeSort(arr []int) {
	process(arr, 0, len(arr)-1)
	return
}

func process(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	process(arr, l, mid)
	process(arr, mid+1, r)
	merge(arr, l, mid, r)
	return
}

func merge(arr []int, l, mid, r int) {
	p1 := l
	p2 := mid + 1

	help := make([]int, 0, 0)

	for p1 <= mid && p2 <= r {
		if arr[p1] <= arr[p2] {
			help = append(help, arr[p1])
			p1++
		} else {
			help = append(help, arr[p2])
			p2++
		}
	}

	for p1 <= mid {
		help = append(help, arr[p1])
		p1++
	}

	for p2 <= r {
		help = append(help, arr[p2])
		p2++
	}
	for i := 0; i < len(help); i++ {
		arr[l+i] = help[i]
	}
	return
}
