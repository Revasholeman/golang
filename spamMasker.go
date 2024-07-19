package main

import (
	"fmt"
)

func spamMasker(str string) string {
	http := [7]string{"h", "t", "t", "p", ":", "/", "/"}
	count := 0
	var buffer []byte
	for _, w := range str {
		switch {
		case string(w) == " ":
			buffer = append(buffer, byte(w))
			count = 0
		case count == 7 && string(w) != " ":
			buffer = append(buffer, '*')
		case string(w) == http[count]:
			buffer = append(buffer, byte(w))
			count++
		default:
			buffer = append(buffer, byte(w))
		}
	}
	return string(buffer)
}

func main() {
	text := `Here's my spammy page: http://hehefouls.netHAHAHA see you http://hehefouls.netHAHAHA`

	fmt.Println(spamMasker(text))
}
