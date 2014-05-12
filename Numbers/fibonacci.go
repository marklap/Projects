/*
:summary: Calculate Fibonacci sequence to N numbers or up to N

:license: Creative Commons Attribution-ShareAlike 3.0 Unported
:author: Mark LaPerriere
:contact: marklap@mindmind.com
:copyright: Mark LaPerriere 2013
*/

package main

import (
	"fmt"
	"math"
	"strings"
)

var MAX_N = math.MaxInt32

func fib_count(N) {

}

func fib_max(N) {

}

func main() {
	var selection string
	fmt.Printf("Would you like to specify a count of numbers or a max number? ('C=count, M=max') : ")
	_, err := fmt.Scanf("%s", &selection)
	lc_selection := strings.ToLower(selection)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	switch {
	case lc_selection == "c":
		fmt.Printf("You typed a '%s'\n", selection)
	case lc_selection == "m":
		fmt.Printf("You typed an '%s'\n", selection)
	default:
		fmt.Println("Error: Please enter either C or M")
		return
	}

	fmt.Println("On to the next thing...")
}
