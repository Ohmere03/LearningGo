package main

import "fmt"

/*
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
solution: https://play.golang.org/p/MDLF3v5EGT
video: 036

 */

const (
	year1 = 2019 + iota
	year2 = 2019 + iota
	year3 = 2019 + iota
	year4 = 2019 + iota

)
func main() {

	fmt.Println(year1, year2, year3,year4)

}