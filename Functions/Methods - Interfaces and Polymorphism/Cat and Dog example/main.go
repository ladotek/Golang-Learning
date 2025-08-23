package main

import "fmt"

type Sayer interface {
	Say()
}
type Dog struct{}

func (d Dog) Say() {
	fmt.Println("Woof!")
}

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("Meow!")
}

func LetThemSpeak(s Sayer) {
	s.Say()
}

func main() {
	dog := Dog{}
	cat := Cat{}
	LetThemSpeak(dog) // Output: Woof!
	LetThemSpeak(cat) // Output: Meow!
}
