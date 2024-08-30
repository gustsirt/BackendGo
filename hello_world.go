package main

import "fmt"

func main() {
	var myString string = "Esto es texto"
	fmt.Println(myString)

	myInt := 12

	if myInt < 0 {
		fmt.Println("Negaivo")
	} else if myInt == 0 {
		fmt.Println("Cero")
	} else {
		fmt.Println("Positivo")
	}
}
