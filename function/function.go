package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"My", "Skill"}
	printMsg("hallo", names)
}

func printMsg(msg string, arr []string) {
	// strings.Join => untuk menyatukan data pada sebuah array
	var nameString = strings.Join(arr, "")
	fmt.Println(msg, nameString)
}
