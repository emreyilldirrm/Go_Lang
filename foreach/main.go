package main

import "fmt"

func main() {
	// var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// for index, value := range numbers {
	// 	fmt.Println(index, value)
	// }

	// for _, value := range numbers {
	// 	fmt.Println(value)
	// }

	// var text = "Emre Yıldırım"
	// for _, value := range text {
	// 	fmt.Println(value)
	// }

	var text = "Emre Yıldırım"
	for _, value := range text {
		fmt.Printf("value %c\n", value)
	}

	names := map[string]int{
		"selina": 10,
		"irem":   40,
		"şule":   20,
	}

	for index, value := range names {
		fmt.Printf("key : %s , value : %d \n  ", index, value)
	}
}
