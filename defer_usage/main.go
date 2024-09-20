package main

import "fmt"

func main() {
	// defer burada fonksiyon diğer yazılan işlemlei bittikten sonra gelipbu işlemi yapmasını ifade eder
	// defer fmt.Println("hello")
	// fmt.Println("Word")

	//deferin kullanım alanları temizlik diye kast ettiğimiz bir işlem yaptık veritabanı bağlantısı dosya vb kapatma gibi

	//defer yapıları stack veri türünü kullanır ilk giren defer son çıkar
	// defer fmt.Println("1.defer")
	// defer fmt.Println("2.defer")
	// defer fmt.Println("3.defer")

	// fmt.Println("main.go çalıştı")

	//burda ise return main fonksiyonunu kestiği için 2.defer hiçbirzaman stack e atılmayacak buda onun return den sonra tanımlanmasınıanlamsız kılacak çalışmayacak
	// var condition = true
	// defer fmt.Println("1.defer")
	// if condition {
	// 	return
	// }
	// defer fmt.Println("2.defer")

	//burda defer fonsiyon bittikten sonra çalıştığı için x y değerleri değişmişse deferden sonra anlamsız kalır defer kendinden önce tanımlanan değerleri baz alır
	// x := 10
	// y := 20

	// x = 30
	// defer fmt.Println("x:", x)
	// defer fmt.Println("y:", y)

	// x = 100
	// y = 200

	// fmt.Println("x:", x)
	// fmt.Println("y:", y)

	//burada panic() methodunu kullandık go programlama dilinde bulunan kullanıldığında programı durdurup hata mesajı veren method
	// bu methoddan sonraki hiçbir kod çalışmaz
	// var condition = true
	// if condition {
	// 	panic("problem var!!")
	// }
	// Cleanup()

	//yukarıdaki örneğe benzer ama panicden önce defer kllanıldığı için panic() methodu görüldükten sonra defer yine program bitince çalışır
	var condition = true
	if condition {
		defer Cleanup()
		panic("problem var!!")
	}
	//2:55
}

func Cleanup() {
	fmt.Println("Clean Work çalıştı")
}
