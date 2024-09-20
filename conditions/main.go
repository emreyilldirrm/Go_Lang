package main

import "fmt"

func main() {
	old := 18

	if old >= 18 {
		fmt.Println("oy kullanabilirsin")
	} else {
		fmt.Println("oy kullanamazsın")
	}

	s1 := 7777
	s2 := 800
	s3 := 78

	if s1 > s2 && s1 > s3 {
		fmt.Println("en büyük sayı s1")
	} else if s2 > s1 && s2 > s3 {
		fmt.Println("en büyük sayı s2")
	} else {
		fmt.Println("en büyük sayı s3")
	}
}
