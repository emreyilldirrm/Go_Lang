package main

import "fmt"

func main() {
	// index := 0
	// for index < 10 {
	// 	fmt.Println(index)
	// 	index++
	// }

	// for index := 0; index < 10; index++ {
	// 	fmt.Println(index)

	// }
	// for index := 0; index < 10; index++ {
	// 	if index == 3 {
	// 		break
	// 	}
	// 	fmt.Println(index)
	for index := 0; index < 10; index++ {
		if index == 3 || index == 4 {
			continue
		}
		fmt.Println(index)

	}

}
