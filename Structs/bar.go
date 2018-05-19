package Structs

//Embedded types
type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func main() {
	//Person is the first field. Then define person fields from person class
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: true,
	}
}
