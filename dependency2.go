package main

import (
	"strings"
)

func WhenYouLeave(s string) string {
	return "When you leave you say " + strings.ToLower(s)
}
