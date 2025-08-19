package main

import "fmt"

type Engine struct {
	electric bool
}

type Vehicle struct {
	Engine
	make  string
	model string
	doors int
	color string
}

func main() {
	car1 := Vehicle{
		Engine: Engine{
			electric: true,
		},
		make:  "Toyota",
		model: "Yaris",
		doors: 3,
		color: "silver",
	}

	car2 := Vehicle{
		Engine: Engine{
			electric: false,
		},
		make:  "Ford",
		model: "Sierra",
		doors: 5,
		color: "blue",
	}

	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car2.model)

	fmt.Println(car1.electric, car1.make, car1.model)
	fmt.Println(car2.electric, car2.make, car2.model)

	fmt.Println(car1.Engine.electric, car1.make, car1.model)
	fmt.Println(car2.Engine.electric, car2.make, car2.model)
}

/*
● Create a type engine struct, and include this field
	○ electric bool
● Create a type vehicle struct, and include these fields
	■ engine
	■ make
	■ model
	■ doors
	■ color
● Create two VALUES of TYPE vehicle
	○ use a composite literal
● Print out each of these values.
● Print out a single field from each of these values.
*/
