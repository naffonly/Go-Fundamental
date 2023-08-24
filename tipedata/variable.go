package main

import "fmt"

func main() {
	// variable tipedata constant
	const firstName = "alan "
	var lastName string = "ragnarson "

	var bilBulat uint8 = 22
	var bilDes = 1.22
	var humanRac bool = true

	// Example Type data Nill/Zero Value
	// var exampleNil string = ""
	// var exampleBool bool = false
	// var exampleNodesimal = 0
	// var example Desimal = 0.0

	// Variabel Pointer
	var num1 int = 4
	var num2 *int = &num1

	*num2 = 2
	// num2 => num1 = 2
	lastName = "Bjorson"

	fmt.Println("Nama : ", firstName, lastName)
	fmt.Println("Umur : ", bilBulat)
	fmt.Println("PowerUp : ", bilDes)
	fmt.Println("Manusia : ", humanRac)
	// num1  = 2
	fmt.Print("pointer value : ", num1)

}
