package Exercises

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	j := 0

	var final string
	arr := make([]string, len(s))
	output := make([]string, len(s))
	arr = strings.Split(s, "")

	for i := len(arr) - 1; i >= 0; i-- {
		output[j] = arr[i]
		final = strings.Join(output, "")
		j++

	}
	fmt.Println(final)

	return "hi"
}
