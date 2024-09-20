package main

//1:30:55 dakika
import "fmt"

func main() {
	//topla(8, 7)
	// fmt.Println(topla2(7, 7))
	// fmt.Println(calculate(77, 4))

	// total, difference, multiply := calculate_2(7, 7)
	// fmt.Println(total, difference, multiply)
	//var numbers = []int{1, 2, 3, 4, 5}
	// fmt.Println(Sum(numbers))
	fmt.Println(Sum2(1, 2, 3, 4))
}

func topla(x int, y int) {
	fmt.Println(x + y)
}

func topla2(x int, y int) int {
	return x + y
}

func calculate(x int, y int) (int, int) {
	return x + y, x - y
}

func calculate_2(x int, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum

}

func Sum2(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
