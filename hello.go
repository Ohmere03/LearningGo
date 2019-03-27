package main

import "fmt"


var y = 43
type apex string
var test apex

func main() {
	x := 42 // Declare and assign a value to variabl with short declaration variable -> only used in a function body
	fmt.Println("This is my first programm in go\n I decladred the variable x which has a value of", x)
	y = 90
	secondFunction()
	test = "new variable type apex"

	fmt.Print(test, "is of type ")
	fmt.Printf("%T", test)


}

func secondFunction() {

	fmt.Println("This is the second function, with y which has a value of", y)


}

//this is how to add comment in go
