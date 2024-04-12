package main

import "fmt"

const (
	BEIJING = iota*10
	SHANGHIA
	SHENZHEN
)

const (
	a, b = iota+1, iota+2
	c, d
	e, f


	g, h =iota * 2, iota * 3
	i, k
)


func main() {
	const length int = 10
	fmt.Print("length =", length)
	fmt.Print("a =", a)
	fmt.Print("g =", g)
	fmt.Print("BEIJING =", BEIJING)
}