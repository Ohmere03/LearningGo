package main

import "fmt"

/*
Create TYPED and UNTYPED constants. Print the values of the constants.
solution: https://play.golang.org/p/NutvJXWUx2
video: 033

*/

type newType int
const (
	a int = 42
	b newType = 41
	c = "test"

)
func main() {

	fmt.Printf("%T\t%T\t%T\n", a,b,c)
	fmt.Println(a,b,c)
}
