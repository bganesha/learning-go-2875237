package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors) - 1])
	fmt.Println(colors)

	numbers := make([]int,5)

	numbers[0] = 43
	numbers[1] = 9812
	numbers[2] = 7121
	numbers[3] = -812
	numbers[4] = 31
	fmt.Println(numbers)

	numbers = append(numbers, 98)
	numbers = append(numbers, -51)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
