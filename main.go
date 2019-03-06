package main

import "fmt"

type Number int

func (number *Number) Invert() (resultNumber int) {
	resultNumber = -1 * int(*number)
	*number = Number(resultNumber)
	return
}

func main() {
	number := Number(10)
	fmt.Println(number.Invert())
	fmt.Println(number.Invert())
	fmt.Println(number.Invert())
	fmt.Println(number.Invert())
}
