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

/*
  func main() {
  data := []float64{45,53,34,56,22,44}
  n := average(data...) //passes in a slice into a variatic function. you need ...
}
*/
