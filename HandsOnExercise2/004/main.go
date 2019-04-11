package main

/*Write a program that
assigns an int to a variable
prints that int in decimal, binary, and hex
shifts the bits of that int over 1 position to the left, and assigns that to a variable
prints that variable in decimal, binary, and hex
solution: https://play.golang.org/p/Ms964T8SbH
video: 034
*/
import "fmt"

var x int = 30

func main() {
	fmt.Printf("%d\t%b\t%#X\n",x,x,x)
	y := x << 1

	fmt.Printf("%d\t%b\t%#X",y,y,y)


}
