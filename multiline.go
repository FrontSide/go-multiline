package multiline

import (
	"fmt"
	"regexp"
)

// String removes all leading whitespace characters from each line
// in the given string s and returns the new string.
func String(s string) string {
	return regexp.MustCompile(`(?m)^(\s*)(\S+)`).ReplaceAllString(s, "$2")
}

// StringWithDelimiter removes all characters up until and inluding
// the first occurrence of the delimiter in each line.
// It returns the new string
func StringWithDelimiter(s string, delimiter rune) string {
	return regexp.MustCompile(
		fmt.Sprintf(`(?m)^(.*)(%x)(\S+)`, delimiter),
	).ReplaceAllString(s, "$2")
}
