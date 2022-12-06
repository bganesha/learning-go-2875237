package main

import (
	"fmt"
)

const av float64 = 6.023

func main() {
	var aString string = "This is from Go!"
	fmt.Println(aString)
	fmt.Printf("The variable type for aString  is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)
	fmt.Printf("The variable type for anIntenger is %T\n", anInteger)

	myInteger := 84
	fmt.Println(myInteger)
	fmt.Printf("The variable type for myIntenger is %T\n", myInteger)

	fmt.Println(av)
	fmt.Printf("The variable type for av is %T\n", av)
}
