package main

import "fmt"

type character struct {
	name  string
	grade int
}

func main() {

	var s1 character
	s1.name = "alan"
	s1.grade = 2

	fmt.Println("nama : ", s1.name)
	fmt.Println("grade : ", s1.grade)
}
