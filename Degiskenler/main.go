package main

import "fmt"

var ad string = "sercan"
func main(){

	// Go değişken isimlendirme kuralları
	// 1. Değişken isimleri harf (A-Z, a-z) veya _ ile başlamalıdır.
	// 2. Değişken isimler harf, rakam (0-9) veya alt çizgi içerebilir. 
	// 3. Değişken isimleri boşluk içeremez. 
	// 4. Değişken isimleri büyük/küçük harfe duyarlıdır. Case Sensitive
	// 5. Büyük harflerle başlayan değişken isimleri paket dışından erişilebilir. Public
	// 6. Küçük harfle başlayan değişken isimleri paket dışından erişilemez. Private
	// 7. Go da bazı anahtar kelimeleri değişken ismi olarak kullanılamaz. (var, func, if, else, return, package, import)
	// 8. camelCase veya snake_case kullanılır. Go topluluğu camelCase daha çok tercih edilir. 
	// 9. Sabitlerde genellikle PascalCase kullanılır. MaxValue 
	//    ALL_CAPS MAX_VALUE
	// 10. Kısa ve anlamlı isimler tercih edilmelidir. 

	// Klasik Değişken Tanımlama
	var name string = "Alin" // Küçük harfle başladığı için sadece paket içinden erişilebilir. 
	var Name string = "Sercan"

	// var : anahtar kelime
	// name: değişken isimlendirme
	// string: veri tipi
	// "Alin": değeri

	fmt.Println(name)
	fmt.Println(Name)

	// Tip çıkarımlı değişken tanımlama
	var city = "İstanbul"
	// Go burada otomatik olarak string tipini algılar
	fmt.Println(city)

	// Kısa Değişken Tanımlama
	country := "Türkiye"
	// := operatörü ile değişken tanımlanır ve değeri atanır
	// := operatörü sadece fonksiyon içinde kullanılabilir
	fmt.Println(country)

	fmt.Println(ad)

	// Çoklu Değişken Tanımlama
	var (
		age int = 30
		height float64 = 1.78
		isStudent bool = false
	)
	fmt.Println(age, height, isStudent)

	// Çoklu Değişken Tanımlama 2
	var (
		x int = 10
		y float32 
		z bool
		w string
	)
	fmt.Println(x,y,z,w)

	// Çoklu Dğeişken Tanımlama 3
	a, b, c := 1, 2.5, "Alin"

	fmt.Println(a,b,c)

	firstName, lastName := "Alin", "Özen"
	fmt.Println(firstName, lastName)

	name = "Tuğçe"
	fmt.Println(name)

	name = "Sercan"

	// Sabit Tanımlama
	const MaxRetryCount int = 5
	const DEFAULT_TIMEOUT = 30
	const ApiBaseUrl = "https://api.example.com" 

	fmt.Println(MaxRetryCount, DEFAULT_TIMEOUT, ApiBaseUrl)

}
