package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		if r == '❗' {
			return "recommendation"
		} else if r == '🔍' {
			return "search"
		} else if r == '☀' {
			return "weather"
		}
	}
	return "default"
	/* rune, _ := utf8.DecodeRuneInString(log)
	switch rune {
	case '❗':
		return "recommendation"
	case '🔍':
		return "search"
	case '☀':
		return "weather"
	default:
		return "default"
	} */
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var modifiedLog string = ""
	for _, r := range log {
		if r == oldRune {
			modifiedLog += string(newRune)
			continue
		}
		modifiedLog += string(r)
	}
	return modifiedLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)
	if numberOfRunes <= limit {
		return true
	}
	return false
}
