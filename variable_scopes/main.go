package main

import "fmt"

var global_degisken = "globalim"
var test_global_degisken = "test global değişken"

func main() {
	// var condition = true
	// if condition {
	// 	var a = 23
	// 	fmt.Println(a)
	// }

	// // if scobunda tanımlanan değşken scobe dışında tanınmaz
	// fmt.Println(a)

	// var condition = true
	// var a = 23
	// if condition {

	// 	fmt.Println(a)
	// }

	// //  değşken main func içinde tanınır dışında tanınmaz
	// fmt.Println(a)

	// test()
	// //test func içindeki değişkenlere sadece test func içinden erişilir
	// fmt.Println(x, y)

	//global değişken bu dosya içerisinde her yerde tanınır
	// fmt.Println(global_degisken)

	// //global değişken örnek
	// fmt.Println(test_global_degisken)
	// test2()
	// fmt.Println(test_global_degisken)

	// //global değişken örnek burada test func içinden global değişkene erişim değişiklik sağladık
	fmt.Println(test_global_degisken)
	test3()
	fmt.Println(test_global_degisken)

	//3:03:00

}

func test() {
	var x = 10
	var y = 20
	fmt.Println(x, y)
}

func test2() {
	var test_global_degisken = 4
	fmt.Println(test_global_degisken)
}

func test3() {
	test_global_degisken = "değiştim"
	fmt.Println(test_global_degisken)
}
