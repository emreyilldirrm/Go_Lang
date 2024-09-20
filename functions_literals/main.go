package main

import "fmt"

func main() {

	// //tanımlandığı yerde kullanan fonksiyon kullanımı
	// func() {
	// 	fmt.Println("F1")
	// }()

	// //parametreli kullanım
	// func(a int, b int) {
	// 	fmt.Println(a + b)
	// }(2, 4)

	// //fonksiyonu değişkene atmak ve değişkenle birlikte fonksiyonnu kullanmak
	// sum := func(a int, b int) int {
	// 	return (a + b)
	// }

	// fmt.Println(reflect.TypeOf(sum))
	// fmt.Println(sum(3, 4))

	//fonksiyona parametre olarak fonksiyon gönderme (kullanımı)
	add := func(x int, y int) int {
		return x + y
	}
	multiply := func(x int, y int) int {
		return x * y
	}

	n1, n2 := calculator(3, 3, add, multiply)
	fmt.Println("toplamı :", n1, "çarpımı:", n2)

	//3:30:00
}

// fonksiyona parametre olarak fonksiyon gönderme (fonksiyon tanımı)
func calculator(n1 int, n2 int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(n1, n2), f2(n1, n2)
}
