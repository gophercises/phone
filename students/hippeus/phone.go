package phone

import (
	"bytes"
	"log"
)

// Normalize filters all non-number characters from an in string.
func Normalize(in string) string {
	var buf bytes.Buffer
	for _, r := range in {
		if r >= '0' && r <= '9' {
			_, err := buf.WriteRune(r)
			if err != nil {
				log.Fatal("Write operation failed\n")
			}
		}
	}
	return buf.String()
}
