package main

import (
	"fmt"

	"github.com/abhisuta/table-driven-test/calculator"
)

func main() {
	a := "1"
	b := "2"
	c, err := calculator.Add(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s + %s = %d\n", a, b, c)
	}
}
