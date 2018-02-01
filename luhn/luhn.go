package luhn

import "strconv"
import "strings"

func Valid(n string) bool {
	replacer := strings.NewReplacer(" ", "")
	n = replacer.Replace(n)

	if len(n) <= 1 {
		return false
	}

	sum := 0
	for i := len(n) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(n[i]))
		if err != nil {
			return false
		}
		if (len(n)-i)%2 == 0 {
			digit = 2 * digit
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
