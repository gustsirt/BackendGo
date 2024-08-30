package main

import "fmt"

func main() {
	myInt := 12

	if myInt < 0 {
		fmt.Println("Negaivo")
	} else if myInt == 0 {
		fmt.Println("Cero")
	} else {
		fmt.Println("Positivo")
	}
}
