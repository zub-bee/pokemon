package repl

import "strings"

func CleanInput(input string) []string {
	var list []string
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	list = strings.Fields(input)
	return list
}
