package main

import taskOne "WBTechL1/1"

func main() {
	h := taskOne.Human{
		Name: "Bob",
		Age:  25,
	}
	h.SayHi()

	a := taskOne.Action{Human: h}
	a.Do()
	a.SayHi()
}
