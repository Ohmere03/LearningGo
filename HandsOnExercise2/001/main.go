package main

import (
	"fmt"
)

/*
Hands-on exercise #1
Write a program that prints a number in decimal, binary, and hex
solution: https://play.golang.org/p/bAQxcEuK8O
video: 031
*/

const (
	a = 42
)

func main() {

	fmt.Printf("%b\n%d\n%#X", a, a, a)
}