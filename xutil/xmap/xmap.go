package xmap

import "strconv"

func KeyString(m map[string]string, k string) string {
	if v, ok := m[k]; ok {
		return v
	}
	return ""
}

func KeyInt64(m map[string]string, k string) int64 {
	if v, ok := m[k]; ok {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}
