package Exercises

import "fmt"

func Looping() {

  for i := 0; i <= 100; i++ {
      fmt.Println(i)
  }
}

func Nestedloop() {
  for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
      fmt.Println(i, " _ ", j)
    }
  }

}

func Conditions() {
    i := 0
    for i < 10 {
      fmt.Println(i)
      i++
    }
}

func Switching() {
  a := 90

  switch a {
  case 3 :
    fmt.Println("The number is 3")
  case 5:
    fmt.Println("The number is 5")
  case 90:
    fmt.Println("The number is 90")
  default:
    fmt.Println("Idk")
  }
}
