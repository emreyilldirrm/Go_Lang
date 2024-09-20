package main

import "fmt"

func main() {
	//2:08
	// var customer1 Customer
	var customer1 = Customer{646544, "Emre", 25, Address{city: "istanbul", district: "ataşehir"}}
	var customer2 = Customer{88888, "Iraz", 24, Address{city: "istanbul", district: "ataşehir"}}

	customer2.id = 4444444
	fmt.Println(customer1)
	fmt.Println(customer2)

	customer1.ToString()

	changeName(customer2)
	customer2.ToString()
	//Burada fonksiyon ile name değerimizi değiştirmeye çalıştık ama olmadı nedeni pointer kullanmadık

	changeNamePointer(&customer2)
	customer2.ToString()

	//dışardan fonksiyon tanımlamayız doğru olan direk structa fonksiyon tanımlanmasıdır ve veri değişikliği yapılıcaksa pointer kullanmamız gerekir
	customer2.changeNamePointer2("Tuana")
	customer2.ToString()
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}
type Address struct {
	city     string
	district string
}

func (customer Customer) ToString() {
	fmt.Printf("Name : %s , Age : %d\n", customer.name, customer.age)
}

func changeName(customer Customer) {
	customer.name = "Selina"
}

func changeNamePointer(customer *Customer) {
	//array slice structlarda değişken başlarına yıldız yazmaız gerekmez
	customer.name = "Selina"
}

//dışardan fonksiyon tanımlamayız doğru olan direk structa fonksiyon tanımlanmasıdır ve veri değişikliği yapılıcaksa pointer kullanmamız gerekir
func (customer *Customer) changeNamePointer2(newname string) {
	//array slice structlarda değişken başlarına yıldız yazmaız gerekmez
	customer.name = newname
}
