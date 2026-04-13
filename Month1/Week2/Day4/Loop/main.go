package main

import "fmt"

func main() {
	// Traditional for loop
	// x := 10
	// for x < 17 {
	// 	fmt.Println("Opemipo")
	// 	x++
	// }

	// Loop using for range
	// for range 5 {
	// 	fmt.Println("Opemipo")
	// }

	// Loop using val and range
	// for val := range 9 {
	// 	fmt.Println(val)
	// }

	// Loop converting bytes to string
	// for val1, val2 := range "Andrew" {
	// 	fmt.Println(val1, "----", string(val2))
	// }

	// for val1, val2 := range "Jesus" {
	// 	fmt.Println(val1, "----", string(val2))
	// }

	// infinite loop (to stop an infinite loop, it is either you return, break, or skip)
	// x := 1
	// for {
	// 	if x == 10 {
	// 		break
	// 	}
	// 	fmt.Println("Opemipo")
	// 	x++
	// }

	// y := 0
	// for {
	// 	if y == 15 {
	// 		break
	// 	}
	// 	fmt.Println("Web3")
	// }

	// count := 15
	// for {
	// 	if count == 0 {
	// 		break
	// 	}
	// 	fmt.Println(count)
	// 	count--
	// }

	// var choice string

    // for {
    //     fmt.Println("\n1. Start\n2. Settings\n3. Exit")
    //     fmt.Print("Enter choice: ")
    //     fmt.Scan(&choice)

    //     if choice == "1" || choice == "2" || choice == "3" {
    //         fmt.Println("Valid choice!")
    //         break
    //     }
    //     fmt.Println("Invalid choice. Try again.")
    // }

	// var num, total int

    // for {
    //     fmt.Print("Enter a number (0 to stop): ")
    //     fmt.Scan(&num)

    //     if num == 0 {
    //         break
    //     }
    //     total += num
    // }

    // fmt.Println("Total:", total)

var password string

    for {
        fmt.Print("Enter password: ")
        fmt.Scan(&password)

        if password == "secret" {
            fmt.Println("Access granted!")
            break
        }
        fmt.Println("Wrong password. Try again.")
    }
}
