package sort

// 快排2.0
// 数组最后一位的值作为基准值p，小于p的放在数组左边，大于p的放在数组右边，等于p的放在数组中间
// 最后p和大于p区间的第一个值交换，这样就保证基准值p的位置确定
// 递归以上行为，使整个数组有序
func quickSort(arr []int) {
	process2(arr, 0, len(arr)-1)
	return
}

func process2(arr []int, l, r int) {
	if l == r {
		return
	}
	// i 是从左向右的指针
	// p1 是左指针，记录小于基准值的区域
	// p2 是右指针，记录大于基准值的区域
	// r 是基准值
	i := l
	p1 := l
	p2 := r - 1
	for i <= p2 {
		// 当前值<基准值，左指针++,i++
		// 交换左指针和当前值
		if arr[i] < r {
			p1++
			i++
			swap(arr, p1+1, i)
		}
		// 当前值>基准值，右指针--
		// 交换当前值和右指针
		if arr[i] > r {
			swap(arr, p2, i)
			p2--
		}
		// 当前值==基准值，i++
		if arr[i] == r {
			i++
		}
	}
	// 基准值和大于区域的第一个值交换
	swap(arr, r, p2-1)
	// 递归处理大于区域和小于区域的数据
	process2(arr, l, p1)
	process2(arr, p2, r)
}

// [1, 4, 3, 5, 1, 4]
// [1, 4, 3, 5, 1] [4]
