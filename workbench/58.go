package workbench

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	return len(s) - strings.LastIndex(s, " ") - 1
}
