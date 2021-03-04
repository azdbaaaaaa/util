package util

func removeDuplicateString(input []string) (output []string) {
	output = make([]string, 0, len(input))
	temp := map[string]struct{}{}
	for _, item := range input {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			output = append(output, item)
		}
	}
	return output
}
