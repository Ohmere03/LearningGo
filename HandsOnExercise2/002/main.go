package main

import "fmt"

/*
Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables.
solution: https://play.golang.org/p/76R-poSzaY

*/

func main() {
	a := 10
	b := 20
	c := (a == b)
	d := (a <= b)
	e := (a >= b)
	f := (a!=b)
	g := (a < b)
	h := (a >b)
	// aussi possible a := (12==23)
	fmt.Println("This is the result of the comparison:", c,d,e,f,g,h)


}
