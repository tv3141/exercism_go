package reverse

func String(s string) string {
	// Reverse the order of code points
	result := ""
	chars := []rune(s) // string to codepoints
	for i := len(chars) - 1; i >= 0; i-- {
		result += string(chars[i])
	}
	return result
}
