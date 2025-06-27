package main

import "sort"

func longestSubsequenceRepeatedK(s string, k int) string {
	// Conta a frequência de cada caractere
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	// Filtra caracteres que aparecem pelo menos k vezes
	var validChars []byte
	for char, count := range freq {
		if count >= k {
			validChars = append(validChars, char)
		}
	}

	// Ordena os caracteres válidos em ordem decrescente (para lexicograficamente maior)
	sort.Slice(validChars, func(i, j int) bool {
		return validChars[i] > validChars[j]
	})

	// Função auxiliar para verificar se uma subsequência pode ser repetida k vezes
	canRepeat := func(subseq string) bool {
		needed := make(map[byte]int)
		for i := 0; i < len(subseq); i++ {
			needed[subseq[i]] += k
		}

		for char, count := range needed {
			if freq[char] < count {
				return false
			}
		}

		// Verifica se subseq*k é realmente uma subsequência de s
		target := ""
		for i := 0; i < k; i++ {
			target += subseq
		}

		j := 0
		for i := 0; i < len(s) && j < len(target); i++ {
			if s[i] == target[j] {
				j++
			}
		}
		return j == len(target)
	}

	// BFS para construir subsequências
	queue := []string{""}
	result := ""

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(current) > len(result) {
			result = current
		}

		// Tenta adicionar cada caractere válido
		for _, char := range validChars {
			next := current + string(char)
			if canRepeat(next) {
				queue = append(queue, next)
			}
		}
	}

	return result
}
