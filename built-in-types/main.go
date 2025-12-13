package main

import "fmt"

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Tesla",
		Model:         "Model S",
		Year:          2022,
	}

	fmt.Printf("My car is a %d %s %s with %d tires. Luxury: %t, Bucket Seats: %t\n",
		myCar.Year, myCar.Make, myCar.Model, myCar.NumberOfTires, myCar.Luxury, myCar.BucketSeats)
}
