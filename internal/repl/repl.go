package repl

import "strings"

func CleanInput(text string) []string {
	trimmed := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(text, "\t", " "), "\n", " "))
	if len(trimmed) == 0 {
		return []string{}
	}
	splitted := strings.Split(strings.ToLower(trimmed), " ")
	output := []string{}
	for _, word := range splitted {
		if len(word) != 0 {
			output = append(output, word)
		}
	}
	return output
}
