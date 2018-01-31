package acronym

import "strings"

// Abbreviate converts a phrase to its acronym.
// Example: Three Letter Acronym -> TLA
func Abbreviate(s string) string {
	replacer := strings.NewReplacer("-", " ")
	s = replacer.Replace(s)

	var firstLetters []string
	for _, word := range strings.Fields(s) {
		firstLetters = append(firstLetters, string(word[0]))
	}
	acronym := strings.ToUpper(strings.Join(firstLetters, ""))
	return acronym
}
