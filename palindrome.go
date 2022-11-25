package bluebird

import (
	"regexp"
	"strings"
)

func CleanString(str string) string {
	result := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(str, "")
	return strings.ToLower(result)
}

func ValidatePalindrome(str string) bool {
	str = CleanString(str)
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}

	return true
}
