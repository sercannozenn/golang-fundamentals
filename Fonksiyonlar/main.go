package main

import (
	"errors"
	"fmt"
)

func main(){
	greet("Alin")

	result, err := divide(5.5,2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bölüm:", result)
	}

	alan := rectangleArea(4.0, 5.0)
	fmt.Println("Dikdörtgen Alanı: ", alan)

	anonim := multiply(5,6)
	fmt.Println("Anonim Sonuç:", anonim)

	fact := factorial(5)
	fmt.Println("Faktöriyel Sonucu:", fact)



	counter := makeCounter()
	fmt.Println("Sayaç 1:", counter())
	fmt.Println("Sayaç 2:", counter())
	fmt.Println("Sayaç 3:", counter())
}

func greet(name string){
	fmt.Println("Merhaba", name)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		//return 0, fmt.Errorf("Bölme hatası: sıfıra bölünemez")
		return 0, errors.New("Bölme hatası: sıfıra bölünemez")
	}
	return a / b, nil
}
// İsimlendirilmiş fonksiyon
func rectangleArea (width, height float64) (area float64){
	area = width * height
	return 
}

// anonim fonksiyon
var multiply = func (a, b int) int {
	return a * b
}

// Recursive Fonksiyon - Kendi kendini çağıran fonksiyonlar
func factorial(n int) int{
	if n == 0{
		return 1
	}

	return n * factorial(n-1)
}
// 5! => 5*4*3*2*1 =
// 5 * factorial(4) -> 5 * (4 * factorial(3)) 
// 5*4*3*2*1* factorial(0) => 5*4*3*2*1*1 = 120

// Closure Fonksiyon 
// Durum bilgisini (state) saklayabilen fonksiyonlardır.
func makeCounter() func() int {
	count := 0
	return func() int{
		count++
		return count
	}
}

