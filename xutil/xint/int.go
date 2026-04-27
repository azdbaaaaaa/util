package xint

func ContainsInt64(l []int64, t int64) bool {
	for _, e := range l {
		if e == t {
			return true
		}
	}
	return false
}

func ContainsInt32(l []int32, t int32) bool {
	for _, e := range l {
		if e == t {
			return true
		}
	}
	return false
}

func ContainsInt(l []int, t int) bool {
	for _, e := range l {
		if e == t {
			return true
		}
	}
	return false
}
