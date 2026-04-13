package main

import "fmt"

func main() {
	// // challenge 1
	// fmt.Println("Printing all Even Numbers between 20 and 55: ")

	// for i := 20; i <= 55; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Print(i, " ")
	// 	}
	// }
	// fmt.Println()

	// challenge 2
	// var myArray = [7]int{20, 22, 24, 26, 28, 30, 32}
	// fmt.Println("Printing new array: ")
	// newArray := [7]int{}

	// for idx, val := range myArray {
	// 	result := val * 2
	// 	newArray[idx] = result
	// }
	// fmt.Println(newArray)

	// challenge 3
	mySlice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Printing reversed slice")

	for i := len(mySlice) - 1; i >= 0; i-- {
		fmt.Print(mySlice[i], " ")
	}
	fmt.Println()

}
