package strutil

import "fmt"

//PrintChar is a func to print each char from string
func Printchar(s string) {
	chars := string([]byte(s))
	for i := 0; i < len(chars); i++ {
		fmt.Println(string(chars[i]))
	}
}
