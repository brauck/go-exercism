package main

import (
	s "exercism/parsing_log_files/parsinglogfiles"
	"fmt"
)

var p = fmt.Println

func main() {
	/* p(s.IsValidLine("[ERR] A good error here"))
	p(s.IsValidLine("Any old [ERR] text"))
	p(s.IsValidLine("[BOB] Any old text")) */

	s.SplitLogLine("section 1<*>section 2<~~~>section 3")
	s.SplitLogLine("section 1<=>section 2<-*~*->section 3")

	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,     // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`,          // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one
	}
	p(s.CountQuotedPasswords(lines))

	p(s.RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27"))

	result := s.TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	})
	p(result)
}
