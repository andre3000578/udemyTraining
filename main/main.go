package main

import "fmt"
import "github.com/andre3000578/udemyTraining/Exercises"

func main() {
	num := []int{58, 98, 201, 4, 87, 689, 102}
	fmt.Println(Exercises.Greatest(num...))
}
