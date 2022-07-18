package execises8

import "unicode/utf8"

// 👎👍

func Application(a string) string {
	for _, myLog := range a {

		if myLog == '❓' {
			return "question mark"
		}
		if myLog == '👎' {
			return "ok hand"
		}
		if myLog == '🏃' {
			return "RUNNER"
		}
		if myLog == '❓' || myLog == '👎' {
			return "question mark"
		}

	}
	return "defult"
}

func Replace(log string, ok1 rune, ok2 rune) rune {

	for _, myLogs := range log {
		if myLogs == ok1 {
			ok1 = ok2
		}
	}
	return ok1

}

func WithinLimit(logLine string, characterLimit int) bool {

	if utf8.RuneCountInString(logLine) == characterLimit {
		return true
	}

	return false

}
