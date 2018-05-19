package Exercises

import "fmt"

type Course struct {
	name         string
	courseNumber float64
	instructor   string
}

type Course struct {
	name         string
	courseNumber float64
	instructor   string
}

func main() {
	var courses int
	var name string
	var number float64
	var instructor string

	fmt.Println("Enter how many courses:")

	fmt.Scan(&courses)

	class := make([]Course, courses)

	for i := 0; i < courses; i++ {
		fmt.Scan(&name)
		fmt.Scan(&number)
		fmt.Scan(&instructor)
		class[i] = Course{name, number, instructor}
	}

	fmt.Println(class)

}
