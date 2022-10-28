package No_9

import "strconv"

func isPalindrome(x int) bool {
	if x==0 {
		return true
	}
	if x < 0 {
		return false
	}

	xStr := strconv.Itoa(x)
	n := len(xStr)
	qiou := n%2
	mid := (n-1)>>1
	// 偶数
	if qiou == 0 {
		left := mid
		right := mid +1
		for left >= 0 && right < n {
			if xStr[left] != xStr[right] {
				return false
			}
			left--
			right++
		}
	} else { //奇数
		left := mid
		right := mid
		for left >= 0 && right < n {
			if xStr[left] != xStr[right] {
				return false
			}
			left--
			right++
		}
	}
	return true
}


