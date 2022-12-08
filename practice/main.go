package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("I ran this program at ", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at ", t)

	fmt.Println(t.Format(time.ANSIC))
}
