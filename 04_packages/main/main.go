package main

import (
	"fmt"
	"math" //importing package from standard library

	"github.com/pk80/student/04_packages/strutil" //importing third party package from github.com
)

func main() {
	fmt.Println(math.Floor(1.323))
	fmt.Println(math.Ceil(1.323))
	fmt.Println(math.Sqrt(2))

	fmt.Println(strutil.Reverse("Correct"))
}
