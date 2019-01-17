package main

import "fmt"

func main() {
	// MAIN TYPES

	// string
	var name1 string                 //declare using var and can be used anywhere in main
	name1 = "Robert_1"               //initialize value
	var name2 = "Robert_2"           //declare and initialize using var
	name3 := "Robert_3"              //declare and initialize using shorthand and can be used inside func only
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
	fmt.Printf("%T : %v\n", age1, age1)

	// int int8 int16 int32 int64
	var age2 int32 = 32
	fmt.Printf("%T : %v\n", age2, age2)

	// uint uint8 uint16 uint32 uint64 uintptr  - unsigned int
	// byte - alias for uint

	// rune - alias for int32
	var runes rune = 23
	fmt.Printf("Rune, %T : %v\n", runes, runes)

	// float32 float64 - usually used float64
	height1 := 1.3 //by default it's a float64
	var height2 float32 = 1.4
	fmt.Printf("%T : %v\n", height1, height1)
	fmt.Printf("%T : %v\n", height2, height2)

	// complex64 complex 128

}
