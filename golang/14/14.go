package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		j := 0

		for j < len(prefix) && j < len(str) && prefix[j] == str[j] {
			j++
		}

		prefix = prefix[:j]
		if prefix == "" {
			return ""
		}
	}

	return prefix
}
