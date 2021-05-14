package xstring

import "bytes"

func Join(in []string, sep string) string {
	// TODO copy
	//bs := make([]byte, 100000)
	//bl := 0
	//for n := 0; n < 100000; n++ {
	//	bl += copy(bs[bl:], "test")
	//}

	var buffer bytes.Buffer
	for i, s := range in {
		buffer.WriteString(s)
		if i != len(in)-1 {
			buffer.WriteString(sep)
		}
	}
	return buffer.String()
}

func Contains(l []string, t string) bool {
	for _, e := range l {
		if e == t {
			return true
		}
	}
	return false
}
