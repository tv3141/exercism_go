package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, " ", "", -1)

	seen := make(map[rune]bool)
	for _, rune := range word {
		if seen[rune] {
			return false
		}
		seen[rune] = true
	}
	return true
}
