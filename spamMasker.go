package main

import (
	"fmt"
)

func spamMasker(message string) string {
	http := [7]rune{'h', 't', 't', 'p', ':', '/', '/'}
	count := 0
	result := make([]rune, 0)
	for _, w := range message {
		switch {
		case w == ' ':
			result = append(result, w)
			count = 0
		case count == 7 && w != ' ':
			result = append(result, '*')
		case w == http[count]:
			result = append(result, w)
			count++
		default:
			result = append(result, w)
		}
	}
	return string(result)
}

func main() {
	text := `Here's my spammy page: http://hehefouls.netHAHAHA see you http://hehefouls.netHAHAHA`

	fmt.Println(spamMasker(text))
}
