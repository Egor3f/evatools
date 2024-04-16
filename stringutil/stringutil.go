package stringutil

import (
	"fmt"
	"strings"
)

// Join unlike strings.join, it supports any types (uses fmt.fprint underhood).
func Join[T any](input []T, sep string) string {
	result := strings.Builder{}
	for i, s := range input {
		if i > 0 {
			result.WriteString(sep)
		}
		_, err := fmt.Fprint(&result, s)
		if err != nil {
			panic("join error")
		}
	}
	return result.String()
}

// EscapeMarkdownV2 escapes string for using in telegram bots. Don't use it besides telegram bot api.
func EscapeMarkdownV2(s string) string {
	chars := []string{"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
	for _, c := range chars {
		s = strings.Replace(s, c, `\`+c, -1)
	}
	return s
}
