package Exercises

import (
	"strings"
)

func Acronym(phrase string) string {
	var q string
	word := strings.SplitAfter(phrase, " ")
	for i := 0; i < len(word); i++ {
		q += strings.ToUpper(string(word[i][0]))
	}
	return q
}
