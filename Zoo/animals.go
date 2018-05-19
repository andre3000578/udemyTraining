package Zoo

import "fmt"

type Animals struct {
	name  string
	voice string
}

func (a Animals) talk() string {
	return a.name + " says " + a.voice
}

type Exhibit struct {
	a1   Animals
	a2   Animals
	name string
}

func (e Exhibit) startExhibit() {
	fmt.Println("The ", e.name, " exhibit is now starting")
	fmt.Println(e.a1.talk())
}

type Zoo struct {
	Exhibit
	name string
}

func (z Zoo) VisitAllExhibits() {
	z.startExhibit()
}
