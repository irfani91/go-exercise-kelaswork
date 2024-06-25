package main

import (
	"fmt"
)

func main() {
	age := 19

	if age < 18 {
		fmt.Println("Not enough! \n")
	} else if age > 60 {
		fmt.Println("Too enough! \n")
	} else {
		fmt.Println("Welcome! \n")
	}

	grade := "D"

	switch grade {
	case "A":
		fmt.Println("Good Job! \n")
	case "B", "C":
		fmt.Println("Nice Job! \n")
		fmt.Println("You can do better! \n")
	case "D":
		fmt.Println("Please see teacher! \n")
		fallthrough
	default:
		fmt.Println("You have try again! \n")
	}
}
