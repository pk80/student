package main

import "fmt"

func main() {
	// declare and initialize a slice
	var personSlice []string
	personSlice = []string{"Robert", "Griesemer", "39", "Developer"}
	fmt.Println(personSlice)

	fmt.Println(personSlice[3])   //access index values
	fmt.Println(personSlice[1:3]) //access index values using range, prints out index 1 to 3 excluding 3

	//decalre and initialize a slice in a slingle line
	var personsSlice = []string{"Robert", "Mike", "Fill"}
	fmt.Println(personsSlice)

	//shorthand declaration of a slice
	assetSlice := []string{"House", "Car", "Land"}
	fmt.Println(assetSlice)
	fmt.Println(assetSlice[2])

	//shorthand declaration of a slice using make keyword
	habitSlice := make([]string, 4, 6) //4 is the length and 6 is the capacity here
	habitSlice = []string{"Reading", "Research"}
	fmt.Println(habitSlice)

	//length and capcithy of a slice
	fmt.Println("Len and Cap of personSlice :", len(personSlice), cap(personSlice))
	fmt.Println("Len and Cap of personsSlice :", len(personsSlice), cap(personsSlice))
	fmt.Println("Len and Cap of assetSlice :", len(assetSlice), cap(assetSlice))
	fmt.Println("Len and Cap of habitSlice :", len(habitSlice), cap(habitSlice))

	//slicing an array
	myArray := [4]int{1, 2, 3, 4}
	mySlice := myArray[1:3] //myArray[lb:rb] left boundary inclusive and right boundary is exclusive
	mySlice1 := myArray[1:] //slicing myArray with its index 1 as 0 index or first index for slice till end
	mySlice2 := myArray[:2] //slicing myArray to slice with index ending index 2 in myArray
	fmt.Println(mySlice, len(mySlice), cap(mySlice))
	fmt.Println(mySlice1, len(mySlice1), cap(mySlice1))
	fmt.Println(mySlice2, len(mySlice2), cap(mySlice2))

	//append to a slice
	habitSlice = append(habitSlice, "Surfing")
	fmt.Println(habitSlice)
}
