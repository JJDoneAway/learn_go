package iterations

import (
	"bytes"
	"strings"
)

func Repeat(char string, length int) string {
	var b bytes.Buffer
	for i := 0; i < length; i++ {
		b.WriteString(char)
	}
	return b.String()
}

func Repeat2(char string, length int) string {

	return strings.Repeat(char, length)
}
