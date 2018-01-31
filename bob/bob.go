package bob

import "strings"

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if isShoutedQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}
	if isQuestion(remark) {
		return "Sure."
	}
	if isShout(remark) {
		return "Whoa, chill out!"
	}
	if isSilence(remark) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func isSilence(remark string) bool { return remark == "" }

func isShout(remark string) bool {
	return remark == strings.ToUpper(remark) &&
		remark != strings.ToLower(remark)
}

func isQuestion(remark string) bool { return strings.HasSuffix(remark, "?") }

func isShoutedQuestion(remark string) bool {
	return isQuestion(remark) &&
		isShout(remark)
}
