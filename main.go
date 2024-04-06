package main

import (
	"regexp"
)

func String(s string) string {
	return regexp.MustCompile(`(?m)^(\s*)(\S*)`).ReplaceAllString(s, "$2")
}
