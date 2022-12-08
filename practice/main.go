package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 68

	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum is ", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Floating Sum is ", floatSum)

	floatSum = math.Round(floatSum*100)/100
	fmt.Println("The floating sum is now ", floatSum)
}
