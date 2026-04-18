package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\W[A-Z]+\W`)
	s := re.FindString(text)
	fmt.Println(s)

	switch s {
	case "[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]":
		return true
	default:
		return false
	}
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	sl := re.Split(text, -1)
	fmt.Println(sl)

	return sl
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)

	for _, v := range lines {
		if re.MatchString(v) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)

	if re.MatchString(text) {
		return re.ReplaceAllString(text, "")
	}

	return text
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\W+\w+`)
	reUser := regexp.MustCompile(`\w+$`)

	for i, v := range lines {
		found := re.FindString(v)
		if found != "" {
			user := "[USR] " + reUser.FindString(found) + " "
			lines[i] = user + lines[i]
		}
	}

	return lines
}
