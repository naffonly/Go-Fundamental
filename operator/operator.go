package main

import (
	"fmt"
)

func main() {
	// Operator Aritmatika
	var value = (((2+6)%3)*4 - 2) / 3

	// Operator Relasional
	var isEqual = value == 2

	// Operator Logika (&& || !)
	var hasValue = false && isEqual

	// Conditional if
	if hasValue == true {
		fmt.Println("result Conditional if : Weeee")
	} else {
		fmt.Println("result Conditional if : ewww")
	}

	// Conditional switch
	switch hasValue {
	case true:
		fmt.Println("result Conditional switch : true")
	case false:
		fmt.Println("result Conditional switch : false")
	default:
		fmt.Println("Nothing")
	}
	// Looping
	for i := 0; i <= 3; i++ {
		fmt.Println("resut Looping : ", i)
	}

	fmt.Println("result Aritmatika : ", value)
	fmt.Println("result Relasional : ", isEqual)
	fmt.Println("result Logika : ", hasValue)

}
