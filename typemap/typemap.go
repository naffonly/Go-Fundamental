package main

import "fmt"

func main() {
	// type map
	var date map[string]int
	date = map[string]int{}
	date["jan"] = 50
	date["feb"] = 20
	fmt.Println("isi type map : ", date["jan"])
}
