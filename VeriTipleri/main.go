package main

import (
	"fmt"
	"strconv"
)

func main(){
	// Go statik tipli bir dildir. 
	// Yani her değişken tipi derleme (compile time) zamanında belirlenir.

	// Hatalar kod çalıştırılmadan önce tespit edillir.
	// Kod okunabilirliğini artar. 
	// Performans yükü azalır.



	// int8 tipi -128 ile 127 arasında değer alabilir 		  				1 byte yer kaplar.
	// int16 tipi -32,768 ile 32,767 arasında değer alabilir  				2 byte yer kaplar.
	// int32 tipi -2,147,483,648 ile 2,147,483,647 arasında değer alabilir  4 byte yer kaplar.
	// int64 tipi -9,223,372,036,854,775,808 ile 9,223,372,036,854,775,807 arasında değer alabilir  
	// 																		8 byte yer kaplar.
	// int tipi ise sistem mimarisine bağlı olarak 32 bit veya 64 bit olabilir. 
	// 															 Bu da göre 4 veya 8 byte yer kaplar.

	// Optimizasyon gerekmedikçe int tipi kullanılması önerilir. 
	var age int = 30
	fmt.Println(age)

	var age2 int16 = 40
	fmt.Println(age2)

	// uint veri tipi sadece pozitif değerler alabilir. 
	var a uint16 = 42
	fmt.Println(a)

	// float32 ve float64 tipleri ondalıklı sayılar için kullanılır. 

	var pi float32 = 3.14
	var e float64 = 2.7128781
	fmt.Println(pi)
	fmt.Println(e)

	// string tipi metin verileri için kullanılır ve "" arasında her alır.
	var name string = "Sercan"
	fmt.Println(name)

	// bool tipi true veya false değerlerini alabilir
	var isActive bool = true
	fmt.Println(isActive)

	// bool tipi için varsayılan değer false
	var isAvailable bool
	fmt.Println(isAvailable)

	// Tür Dönüşümleri

	var x int = 10
	var y float64 = float64(x) // int'ten float64 e dönüşüm gerçekleşir.
	fmt.Println(y)

	// float to int
	var z float64 = 9.99
	var w int = int(z)
	fmt.Println(w)

	// int to string
	var num int = 100
	var sayi string = fmt.Sprintf("%d", num)
	fmt.Println(sayi)

	// string to int
	var sayi2 string = "200"
	var num2 int
	fmt.Sscanf(sayi2, "%d", &num2)
	fmt.Println(num2)

	// strconv paketi ile dönüşümler yapılabilir. 

	// string to int
	// Atoi'nin anlammı ASCII to integer. ASCII string ifadeleri tanımlıyor.
	numStr := "12345"
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Hata Alındı: ", err)
		return
	}
	fmt.Println(numInt)

	// int to string
	// Itoa integer to ASCII

	numInt2 := 645
	numStr2 := strconv.Itoa(numInt2)
	fmt.Println(numStr2)

	// float64'ten stringe dönüşüm
	floatVal2 := 2.7984
	floatStr2 := strconv.FormatFloat(floatVal2, 'f', 3, 64)
	fmt.Println(floatStr2)
	// 'f' : format türü (f: ondalıklı)
	//  3  : virgülden sonra kaç basamak gösterileceği. 
	// 64  : float64 için

	// Değişken Veri Tipini Öğrenme
	
	fmt.Printf("FloatVal2 Değişkeninin Türü: %T\n", floatVal2)
	fmt.Printf("floatStr2 Değişkeninin Türü: %T\n", floatStr2)


}
