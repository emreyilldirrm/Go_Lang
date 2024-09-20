package main

import (
	"fmt"
	"reflect"
)

func main() {
	// names := make(map[int]string)

	// names[1] = "emre"
	// names[2] = "emre"
	// names[3] = "emre"

	// lastnames := make(map[string]int)

	// lastnames["emre"] = 2
	// lastnames["nes"] = 3
	// lastnames["met"] = 4

	// fmt.Println(lastnames)

	names := map[string]int{
		"Emre": 7,
		"ka":   8,
	}

	fmt.Println(reflect.TypeOf(names))
	fmt.Println(names)

}
