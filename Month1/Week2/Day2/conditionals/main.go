package main

import "fmt"

func main() {
// 	var num int = 37
// 	if num >= 0 {
// 		fmt.Printf("%d is positive.", num)
// 	} else if num < 0 {
// 		fmt.Printf("%d is negative.", num)
// 	} else {
// 		fmt.Printf("%d is neutral.", num)
// 	}

const age int = 20
 if age >= 0 && age <= 12 {
	fmt.Print("User is a child.\n")
} else if age > 12 && age <= 19 {
	fmt.Print("User is a teenager.\n")
} else if age >= 20 {
	fmt.Print("User is an adult.\n")
} else {
	fmt.Print("Invalid age.\n")
}
}
