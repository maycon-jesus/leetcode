package main

func divideString(s string, k int, fill byte) []string {
	size := len(s) / k
	if len(s)%k != 0 {
		size++
	}
	arr := make([]string, size)

	for i := 0; i < len(s); i += k {
		end := i + k
		if end > len(s) {
			end = len(s)
		}
		buf := make([]byte, k)
		copy(buf, s[i:end])
		if end-i < k {
			for j := end - i; j < k; j++ {
				buf[j] = fill
			}
		}
		arr[i/k] = string(buf)
	}
	return arr
}
