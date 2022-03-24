package text

import "strings"

func GetFirstLetterLowered(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}
