# Golang LINQ
Golang version of .NET LINQ

```go
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
			owner: "fernandinho",
			model: "2",
		}, {
			year:  2022,
			owner: "pedrinho",
			model: "2",
		},
		{
			year:  2017,
			owner: "joaozinho",
			model: "2",
		},
	}

	output := From(cars).Where(func(car Car) bool {
		return car.year >= 2016
	}).Reverse().ToSlice()

	fmt.Println(output)
	// [{2017 joaozinho 2} {2022 pedrinho 2} {2018 fernandinho 2}]
}
```
