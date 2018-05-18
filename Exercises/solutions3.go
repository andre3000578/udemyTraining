package Exercises

func makeslice() {
	apple := make([]int, 70, 100)

	dog := make([]string, 39, 200)

	apple = append(apple, dog...)

	//Deletes element 4
	apple = append(apple[:4], apple[5:])
}
