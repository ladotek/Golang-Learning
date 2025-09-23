package main

import "fmt"

type dog struct { // Define a struct type dog with one field, first (the dog’s name).
	first string
}

func (d dog) walk() { // Method walk has a value receiver (d dog).
	fmt.Println("My name is", d.first, "and I'm walking.") // It reads d.first and prints a line.
}

//Because it has a value receiver, it works on a copy of the dog; it cannot mutate the caller’s dog.

func (d *dog) run() { // Method run has a pointer receiver (d *dog).
	d.first = "Rover" // It sets d.first = "Rover"—this mutates the original dog the pointer points to—then prints.
	fmt.Println("My name is", d.first, "and I'm running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"} // Create a value of type dog with name "Henry".
	d1.walk()          // Calls the value-receiver method. Prints: My name is Henry and I'm walking. -- d1 remains "Henry".
	d1.run()           // Although run requires *dog, this works because d1 is addressable; Go implicitly takes &d1 for the call.
	//run sets the original d1.first to "Rover" and prints: My name is Rover and I'm running

	// youngRun(d1) --> does not work
	// When you have a type value (not a pointer) you can only use value receivers to implement an interface

	d2 := &dog{"Padget"} // Create a pointer to a dog with name "Padget".
	d2.walk()            // walk has a value receiver, but calling it on *dog is allowed—Go implicitly dereferences the pointer.Prints: My name is Padget and I'm walking. No mutation.
	d2.run()             // Pointer receiver on a pointer—no conversion needed. Sets d2.first = "Rover" and prints: My name is Rover and I'm running

	youngRun(d2) // This works! When you have a value that is of type pointer, you can use pointer receivers or value receivers to implement an interface.
}

/*
Value receiver (func (d dog)): works on a copy; no mutation of the caller.

Pointer receiver (func (d *dog)): can mutate the original.

Method-call conveniences:

If you have a value and call a pointer-receiver method, Go can implicitly take the address (when the value is addressable).

If you have a pointer and call a value-receiver method, Go can implicitly dereference it.

After d1.run(), d1.first becomes "Rover". After d2.run(), d2.first also becomes "Rover".
*/
