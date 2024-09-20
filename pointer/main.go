package main

import "fmt"

func main() {
	// var a int
	// var p *int

	// a = 10
	// p = &a

	// fmt.Println(a)
	// fmt.Println(p)
	// fmt.Println(*p)

	// //a değişkenin tutulduğu pointerı güncelledik a değişkenide güncellendi
	// *p = 20
	// //bakalım
	// fmt.Println(a)

	// var a int
	// var b int
	// var p *int

	// a = 10

	// b = a
	// p = &a

	// *p = 30
	// fmt.Println(b, a)

	// var n1 int

	// n1 = 50

	// nAdd(n1)

	// fmt.Println("main değeri: ", n1)
	// //değişkenlerimizi fonksiyon ile değerini değiştirdiğimizde pointer kullanarak adresini değiştirmez ise değişkenimizin değeri sabit kalıyor

	// //////////////////////////////////////
	// nAddPointer(&n1)
	// fmt.Println("main değeri: ", n1)
	//bu şekilde değişken güncellenmiştir
	//////////////////////////////////////

	var number = []int{1, 2, 3, 4, 5, 6}

	fmt.Println(number)
	changeValue(number)
	fmt.Println(number)
	//Burada pointer kullanmadan slice deeğerimiz güncellendi nedeni ise go-lang int string float boolean temel değişken türlerinin değerlerini array ve slicesların ise direk değişken referencesını gönderdiğinden ötürü pointer kullanmadan değer güncelleniyor

}

func nAdd(x int) { //pass by value
	x = x + 12
	fmt.Println("nAdd değeri: ", x)
}

func nAddPointer(x *int) { //pass by reference
	*x = *x + 12
	fmt.Println("nAddPointer değeri: ", x)
}

func changeValue(numbers []int) {
	numbers[0] = 5000
}
