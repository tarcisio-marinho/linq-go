package main

import (
	"fmt"
	. "linq-go/Operations"
)

type Car struct {
	year         int
	owner, model string
}

func main() {
	cars := []Car{
		{
			year:  2016,
			owner: "tarcisio",
			model: "2",
		}, {
			year:  2018,
			owner: "tarcisio",
			model: "2",
		}, {
			year:  2022,
			owner: "tarcisio",
			model: "2",
		},
		{
			year:  2017,
			owner: "joaozinho",
			model: "2",
		},
		{
			year:  2017,
			owner: "joaozinho",
			model: "2",
		},
		{
			year:  2017,
			owner: "joaozinho",
			model: "2",
		},
	}

	for i := 0; i < 10; i++ {
		cars = append(cars, Car{
			year:  2012,
			owner: "joao",
			model: "2",
		})
	}

	othersCars := cars
	cars = append(cars, Car{
		year:  2022,
		owner: "zezinho",
		model: "Z",
	})
	cars = append(cars, Car{
		year:  2022,
		owner: "zezinho",
		model: "Z",
	})

	output := From(cars).Where(func(car Car) bool {
		return car.year >= 2015
	}).Except(From(othersCars)).Distinct().Reverse().ToSlice()

	fmt.Println(output)

}
