package main

import (
	"errors"
	"fmt"
)

func main() {
	// var x int
	// var y bool
	// var c float32
	// var e string

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(c)
	// fmt.Println(e)
	// //biz eğer temel tiplere değer verme isek go bizim yerime bir değer tanımlar fakt pointer yani reference tipler de öyle değildir

	// var r *int
	// fmt.Println(r)
	// //reference tipe yani pointere <nil> atadı bizde fonksiyonlarımızdan error yapılarımızı <nil> dönmüşmü dönmemişmi kontrol edicez

	// //örnek
	// var er *int
	// if er == nil {
	// 	fmt.Println("er herhangi bir reference içermiyor")
	// }

	//////////////////////
	//error handlik uygulamalarımızda kullanıyoruz ve yapıs bu şekilde
	// fileData,er:=os.ReadFile("sample.txt")
	//eğer error kontrolu yapmayacaksınız bu şkeilde _ çizgi koyun
	// fileData, _ := os.ReadFile("sample.txt")
	//////////////////////////

	//error handling örnek
	// fileData, err := os.ReadFile("sample.txt")
	// if err != nil {
	// 	fmt.Println("error", err)
	// } else {
	// 	fmt.Println(fileData)
	// }

	// //fonksiyon error handlin ile örnek
	// result, err_ := divide(10, 0)
	// if err_ != nil {
	// 	fmt.Println("error", err_)
	// } else {
	// 	fmt.Println("result", result)
	// }

	// //3:44:00
	P_S := ProductService{}

	err := P_S.Add(Product{id: 1, name: "", price: 1000})
	if err != nil {
		fmt.Println("error", err)
	}
}

type Product struct {
	id    int
	name  string
	price int
}
type ProductService struct {
}

func (productservice *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return errors.New("Ürün ismi boş olamaz")
	}
	if product.price < 0 {
		return errors.New("Ürün fiyatı 0'dan büyük olmalıdır")
	}
	fmt.Println("Ürün eklendi")
	return nil
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("payda sıfır olamaz")
	} else {
		return x / y, nil
	}
}
