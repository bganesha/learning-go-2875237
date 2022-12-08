package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter some text: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("You entered the text: '", input, "'")
	
}
