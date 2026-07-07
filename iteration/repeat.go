package iteration

import "strings"

func Repeat(char string, num int) string {
	var repeated strings.Builder
	for i := 0; i < num; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
