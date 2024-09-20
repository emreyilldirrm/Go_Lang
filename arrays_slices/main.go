package main

import "fmt"

func main() {
	// var arr [3]string

	// arr[0] = "emre"
	// arr[1] = "kayra"
	// arr[2] = "caner"

	// fmt.Println(arr)

	// var arr = [4]string{"emre", "yıldırım", "ıraz", "sinem"}
	// fmt.Println(arr)
	// fmt.Println(arr[0.:2])

	var name = []string{"sinem", "mihriban", "tuana"}
	fmt.Println(name)
	name = append(name, "selina")
	fmt.Println(name)

}
