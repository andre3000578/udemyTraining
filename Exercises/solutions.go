package Exercises

import "fmt"

func Helloworld() {
  b := "D'Andre"
  fmt.Println("Hello, my name is ", b)
}

func Name() {
  var name string;
  fmt.Println("Enter your name")

  fmt.Scan(&name)

  fmt.Println("Hello ", name)
}

func Number() {
  var snumber int;
  var bnumber int;

  fmt.Println("Enter a small number")

  fmt.Scan(&snumber)

  fmt.Println("Enter a big number")
  fmt.Scan(&bnumber)

  fmt.Println(bnumber % snumber)

}


func Even() {

  for i := 0; i <= 100; i++ {
    if i % 2  == 0 {
      fmt.Println(i)
    }
  }
}

func Fizzbuzz() {
  for i := 1; i < 100; i++ {
    if i % 3 == 0 && i % 5 == 0 {
      fmt.Println("FizzBuzz")
    } else if i % 3 == 0  {
      fmt.Println("Fizz")
    } else if i % 5 == 0 {
      fmt.Println("Buzz")
    }
  }
}

func Naturals() {
  var sum int;
  for i := 0; i < 10; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }

  }
  fmt.Println(sum)
}
