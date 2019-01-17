package main

import "fmt"

func main() {
	// declare and initialize an array
	var personArr [4]string
	personArr[0] = "Robert"
	personArr[1] = "Griesemer"
	personArr[2] = "39"
	personArr[3] = "Developer"

	fmt.Println(personArr)
	//access index values
	fmt.Println(personArr[3])
	//access index values using range
	fmt.Println(personArr[1:3]) //prints out index 1 to 3 excluding 3

	//shorthand declarations of an array
	assetArr := [4]string{"House", "Car"}
	fmt.Println(assetArr)
	fmt.Println(assetArr[3])

	//len of an array
	fmt.Println("Length of personArr is :", len(personArr))
	fmt.Println("Length of assetArr is :", len(assetArr))
	//cap of an array
	fmt.Println("Capacity of personArr is :", cap(personArr))
	fmt.Println("Capacity of assetArr is :", cap(assetArr))

}
