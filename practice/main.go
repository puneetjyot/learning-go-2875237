package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 34, 45
	intSum := i1 + i2 + i3
	fmt.Println(intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	f4 := 23
	floatSum := f1 + f2 + f3 + float64(f4)
	fmt.Println(floatSum)

	fmt.Println("Math")

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println(floatSum)

	radius := 15.5
	circumfrence := radius * 2 * math.Pi
	fmt.Printf("the circumfrence is %.2f\n", circumfrence)

}
