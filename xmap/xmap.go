package xmap

func KeyString(m map[string]string, k string) string {
	if v, ok := m[k]; ok {
		return v
	}
	return ""
}
