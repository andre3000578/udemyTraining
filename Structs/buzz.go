package Structs

import "fmt"

type baseball struct {
	equipment string
}

//anonymous interface
func hit(a interface{}) {
	fmt.Println(a)
}
