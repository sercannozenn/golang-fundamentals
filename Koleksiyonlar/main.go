package main

import "fmt"

func main() {
	
	// Array => Sabit boyuutlu listelerdir -> Belirli sayıda eleman tutar
	// Tek bir veri tipi tutarlar
	// Value typetır. Kopyalanır referans tutulamaz.
	// Boyutu sonradan değiştirilemez
	// Yeni eleman eklenemez
	// Mevcut eleman silinemez
	// Eleman güncellenebilir. 

	var numbers [3]int = [3]int{10,20,30}
	fmt.Println("Array elemanları", numbers[0], numbers[1], numbers[2])

	numbers[1] = 99
	fmt.Println("Güncelleme sonrası array:", numbers)

	fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")

	// Array de aralık seçimi - slicing 
	subArray := numbers[0:2]
	// 0. indexten başlar 2.indexe kadar elemanları alır.
	// 2.index dahil değil.
	fmt.Println("Alt array:", subArray)

	subArray2 := numbers[1:]
	fmt.Println("Alt array2:", subArray2)

	subArray3 := numbers[:2]
	fmt.Println("Alt array3: ", subArray3)

	fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")

	// İç içe array örneği - matrix
	var matrix [2][3]int = [2][3]int{
		{1,2,3},
		{4,5,6},
	}

	fmt.Println("0.boyut 1.index:", matrix[0][1])
	fmt.Println("1.boyut 2.index:", matrix[1][2])

	fmt.Println("----------------------------------")
	fmt.Println("----------------------------------")

	// Slice
	// Dinamik uzunluktadır. Boyutu sonradan değiştirilebilir. 
	// Tek bir veri tipi tutarlar
	// Hafızada bir array referans tutarlar. 
	// reference typetır. Kopyalanamaz, referans tutulabilir. 
	// Gerçek hayatta da genellikle array den çok slicelar kullanılır. 

	var numberSlice []int = []int{10,20,30}
	fmt.Println("Slice elemanları", numberSlice[0], numberSlice[1], numberSlice[2])

	// Eleman ekleme 
	numberSlice = append(numberSlice, 40)
	fmt.Println("Eklendikten sonra slice:", numberSlice)

	// birden fazla eleman ekleme
	moreNumers := []int{50,60}
	numberSlice = append(numberSlice, moreNumers...)
	fmt.Println("Daha fazla eleman eklendikten sonra", numberSlice)

	// güncelleme
	numberSlice[0] = 45
	fmt.Println("Güncelleme sonrası:", numberSlice)

	// Slice aralık seçme slicing
	subSlice := numberSlice[1:4]
	// 1.indexten başlar 4. index e kadar 
	// 4. index dahil değil!
	fmt.Println("Alt slice:", subSlice)

	subSlice2 := numberSlice[2:]
	// 2.indexten başla sonuna kadar git
	fmt.Println("Alt slice2:", subSlice2)

	subSlice3 := numberSlice[:4]
	// 0. indexten başla 4 e kadar git. 4. index gahil değil
	fmt.Println("Alt slice3:", subSlice3)

	// Slice silme 
	// 2. indexteki elemanı sildirme
	indexToRemove := 2
	fmt.Println("2.indexe kadar elemanlar:", numberSlice[:indexToRemove])
	fmt.Println("2.indexten soranki elemanlar:", numberSlice[indexToRemove+1:])
	numberSlice = append(numberSlice[:indexToRemove], numberSlice[indexToRemove+1:]...)
	fmt.Println("Eleman silindikten sonra:", numberSlice)

	// iç içe slice örneği

	matrixSlice := [][]int{
		{1,2,3},
		{4,5,6},
	}

	fmt.Println("Matrix slice elemanı:", matrixSlice[0][1])
	fmt.Println("Matrix slice elemanı2:", matrixSlice[1][1])

	fmt.Println("--------------------------")
	fmt.Println("--------------------------")

	// Map
	// Anahtar - Değer key -value
	// hızlı veri erişimi sağlarlar
	// Anahtarlar benzersizdir unique
	// Değerler herhangi bir veri tipi olabilir. 

	prices := map[string]float64{
		"Kahve": 54.90,
		"Çay": 25.90,
		"Pasta": 140.90,
	}
	fmt.Println("Kahve Fiyatı:", prices["Kahve"])
	
	prices["Kek"] = 30.0
	fmt.Println("Güncel Liste:", prices)
	
	prices["Pasta"] = 300.56
	fmt.Println("Güncel Liste:", prices)
	
	delete(prices, "Çay")
	fmt.Println("Güncel Liste:", prices)


	//İç içe map

	persons := map[string]map[string]string{
		"123": {
			"name": "Sercan",
			"lastname": "Özen",
		},
		"456": {
			"name": "Alin",
			"lastname": "Özen2",
		},
	}

	fmt.Println("Person 123 adı: ", persons["123"]["name"])
	fmt.Println("Person 456 soyadı: ", persons["456"]["lastname"])

	// MAp kontrol etme
	price, exists := prices["Pasta2"]
	if exists {
		fmt.Println("Pasta Fiyatı:", price)
	}else {
		fmt.Println("Pasta bulunamadı")
	}

	// Döngüler
	fmt.Println("--------------------")
	fmt.Println("--------------------")
	fmt.Println("--------------------")

	fmt.Println("Numbers Arrayinin Eleman Sayısı:", len(numbers))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index:", i, "Değer:",numbers[i])
	}

	fmt.Println("Array üzerinde range ile döngü kurma")
	for i,v := range numbers {
		fmt.Println("Index:", i, "Değer:", v)
	}

	fmt.Println("Slice üzerinde range")
	for i,v := range numberSlice {
		fmt.Println("Index:", i, "Değer:", v)
	}

	fmt.Println("Map üzerinde range")
	for k,v := range prices {
		fmt.Println(k, "Fiyat:",v)
	}

	fmt.Println("map sadece keyler")
	for k := range prices{
		fmt.Println("Anahtar:", k)
	}

	fmt.Println("map sadece değerler")
	for _, v:= range prices {
		fmt.Println("Değer:", v)
	}
}
