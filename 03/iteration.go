package iteration

import "fmt"

func Repeat(character string, repeatCount int) string {
	var repeated string

	fmt.Scanln(repeatCount)

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
