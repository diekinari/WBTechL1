package taskOne

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %v years old and  you can call me on %s\n", h.Age, h.Name)
}

type Action struct {
	Human
}

func (a Action) Do() {
	fmt.Println("I'm doing!")
}
