package twofer

import "fmt"

func ShareWith(s string) string {
	if s == "" {
		s = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", s)
}
