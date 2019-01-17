package main

import "fmt"

func main() {
	// MAIN TYPES

	// string
	var name1 string                 //declare using var
	name1 = "Robert_1"               //initialize value
	var name2 = "Robert_2"           //declare and initialize using var
	name3 := "Robert_3"              //declare and initialize using shorthand
	fmt.Println(name1, name2, name3) //fmt package doc: godoc.org/pkg/fmt
	fmt.Printf("%T : %v\n", name1, name1)

	// bool
	isMale := true
	isMale = false
	fmt.Printf("%T : %v\n", isMale, isMale)

	// const
	const isCool = true
	//isCool = false //value cannot be changed as it is a constant
	fmt.Printf("%T : %v\n", isCool, isCool)

	// int
	age1 := 32
	var age2 int32 = 32
	fmt.Printf("%T : %v\n", age1, age1)
	fmt.Printf("%T : %v\n", age2, age2)

	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr  - unsigned int
	// byte - alias for uint
	// rune - alias for int
	// float32 float64 - usually used float64
	// complex64 complex 128

}
