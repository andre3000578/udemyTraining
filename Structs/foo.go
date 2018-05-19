package Structs

type foo int

func typeexample() {
	var myage foo
}

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person("James", "Bond", 20)
	p2 := person("Miss", "Moneypenny", 24)
}


//exercise
type customerinfo struct {
  name string
  age int
  nationality string

  p1 := customerinfo("john", 23, "american")
  p2 := customerinfo("mary", 45, "yemen")

  fmt.println(p1.name, p1.age, p1.nationality)

  p1.name = "jake"

}
