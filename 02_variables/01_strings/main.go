package main

import "fmt"

func main() {
	//declaring a string
	var fname string
	//assigning value to a string variable
	fname = "Robert"
	//printing out the variable value
	fmt.Println(fname)

	//shorthand declare and assign value to a string
	lname := "Griesemer"
	fmt.Println(lname)

	//concatinating two variables with another string constraing
	fmt.Println("Golang is designed by", fname, lname, "in 2007.")

	//escape character
	fmt.Println("This language was also designed by \nRob Pik\nKen Thompson")  //next line
	fmt.Println("This language was also designed by \tRob Pik \tKen Thompson") //tab space
	fmt.Println("This language was also designed by \aRob Pik \aKen Thompson") //slips the character
}
