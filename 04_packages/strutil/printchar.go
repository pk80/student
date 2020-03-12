package strutil

import "fmt"

// Printchar is a func to print each char from string
func Printchar(s string) {
	chars := string([]byte(s))
	for i := 0; i < len(chars); i++ {
		fmt.Println(string(chars[i]))
	}
}
