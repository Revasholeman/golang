package main

import (
	"bytes"
	"fmt"
)

func spamMasker(str string) string {
	http := [7]string{"h", "t", "t", "p", ":", "/", "/"}
	count := 0
	var buffer bytes.Buffer
	for _, w := range str {
		switch {
		case string(w) == " ":
			buffer.WriteByte(byte(w))
			count = 0
		case count == 7 && string(w) != " ":
			buffer.WriteByte('*')
		case string(w) == http[count]:
			buffer.WriteByte(byte(w))
			count++
		default:
			buffer.WriteByte(byte(w))
		}
	}
	return buffer.String()
}

func main() {
	text := `Here's my spammy page: http://hehefouls.netHAHAHA see you http://hehefouls.netHAHAHA`

	fmt.Println(spamMasker(text))
}
