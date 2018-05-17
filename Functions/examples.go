package Functions

import "fmt"

func greet(fname, lname string) (string, string) {
  return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

func average(sf ...float64) float64 {
  total := 0.0
  for _, v := range sf {
    total += v
  }
  return total / float64(len(sf))
}

func makeGreater() func() string {
  return func() string {
    return "Hello world!"
  }
}

//Example of  a callback
//Takes function as a parameter. A func that takes an int.
//Passes a func as an argument
func visit(numbers []int, callback func(int)) {
  for _, n := range numbers {
    callback(n)
  }
}

func filter(numbers []int, callback func(int) bool) []int {
  var xs []int{}

  for _, n := range numbers {
    if callback(n) {
      xs = append(xs, n)
    }
  }
  return xs
}
/*
  func main() {
  data := []float64{45,53,34,56,22,44}
  n := average(data...) //passes in a slice into a variatic function. you need ...
  greet := makeGreater()

  xs := filter([]int{1,2,3,4}, func(n int) bool {
  return n > 1
})

  fmt.Println(xs)
}
*/
