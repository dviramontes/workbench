package workbench

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0
	var result strings.Builder

	for i < len(word1) && j < len(word2) {
		result.WriteByte(word1[i])
		result.WriteByte(word2[j])
		i++
		j++
	}

	// drain the rest of the pointers into the word
	for i < len(word1) {
		result.WriteByte(word1[i])
		i++
	}
	for j < len(word2) {
		result.WriteByte(word2[j])
		j++
	}

	// return word
	return result.String()
}
