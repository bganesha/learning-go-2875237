
package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Memory address of anInt is ", p)
	fmt.Println("Value of pointer is ", *p)
}
