package Exercises

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

// half := func(n int) (int, bool) {
//   return n/2, n%2 == 0
// }

func Greatest(list ...int) int {
	var great int
	for _, i := range list {
		if i > great {
			great = i
		}
	}
	return great
}

func Foo(num ...int) {
	fmt.Println(num)
}
