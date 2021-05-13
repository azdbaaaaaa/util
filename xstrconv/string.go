package xstrconv

import "strconv"

func Int642Str(i int64) string {
	return strconv.FormatInt(i, 10)
}
