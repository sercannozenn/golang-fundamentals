package main

import (
	"fmt"
	"slices"
)

func main()  {
	fmt.Println("=== KAFE SİPARİŞ YÖNETİMİ ===")

	var tables [3]string = [3]string {"Masa 1", "Masa 2", "Masa 3"}
	fmt.Println("Masalar:", tables)
	
	tables[1] = "Bahçe Masası"

	fmt.Println("Güncellenmiş Masalar:", tables)

	menu := map[string]float64{
		"Kahve": 54.9,
		"Çay": 25.0,
		"Pasta": 45.0,
		"Su": 10.0,
		"Sandviç": 70.0,
	}
	fmt.Println("Menü Listesi:")
	for item,price := range menu{
		fmt.Printf("- %s: %.2f₺\n", item, price)
	}

	menu["Kek"] = 30.0
	fmt.Println("Yeni Ürün Eklendi: Kek - 30 TL")

	orders := []string{}

	fmt.Println("Sipariş Alınıyor...")
	orders = append(orders, "Kahve", "Pasta", "Su", "Şeker")

	fmt.Println("Aktif Siparişler:", orders)

	orders = append(orders, "Çay")
	fmt.Println("Yeni bir sipariş eklendi:", orders)

	orders = removeAt(orders, 2) // Su silinsin
	fmt.Println("Su iptal edildi", orders)



	var total float64

	for _,item := range orders{
		price, ok := menu[item]
		if ok {
			total += price
		}else {
			fmt.Println("Menüde olmaya ürün:", item)
		}
	}
	fmt.Printf("Toplam Tutar: %2.f ₺ \n", total)

	customers := map[string]map[string]string{
		"123": {
			"name": "Sercan",
			"lastname": "Özen",
			"masa": "Masa 12",
		},
		"456": {
			"name": "Alin",
			"lastname": "Tanrıverdi",
			"masa": "Bahçe Masası",
		},
	}

	fmt.Println("Müşteri Listesi:")
	for id, info := range customers{

		if slices.Contains(tables[:], info["masa"])  {
			fmt.Printf("- %s: %s %s (%s)\n", id, info["name"], info["lastname"], info["masa"])
		}else{
			fmt.Println("Masa bulunamadı.")
			break
		}

	}

	fmt.Println("Sipariş Yönetimi Tamamlandı. ")




}
// T bir type parametresidir, herhangi bir tip olabilir.
// any ise T'nin herhangi bir tip olabileceğini belirtir.
// Bu sayede fonksiyon her türlü slice ile çalışabilir. 
// Yani string, int, float64 gibi tiplerle kullanılabilir. 
func removeAt[T any](s []T, idx int) []T {
 // s []T: Silinecek elemanın bulunduğu slice
 // idx int: Silinecek elemanın indexi
 // []T: Dönüş değeri - Yeni slice (orjinal değişmez)

 // Güvenlik kontrolü
 if idx < 0 || idx >= len(s) {
	return s// Geçersiz index ise orjinal slice ı döndür. 
 }

 return append(s[:idx], s[idx+1:]...)

}