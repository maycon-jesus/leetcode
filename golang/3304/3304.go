package main

var alphabet = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func kthCharacter(k int) byte {
	min, max := 97, 122
	nstr := "a"
	for len(nstr) < k {
		generated := make([]byte, len(nstr))
		for i := 0; i < len(nstr); i++ {
			generated[i] = alphabet[(int(nstr[i])-min+1)%(max-min+1)]
		}
		nstr += string(generated)
	}
	return nstr[k-1]
}
