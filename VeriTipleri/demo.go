package main

import "fmt"

func main(){

	fmt.Println("Kafe Sipariş Uygulamasına Hoş Geldiniz!")

	var customerName string = "Alin"
	var drink string = "Latte"

	var quantity int = 2

	var price float64 = 144.90

	var isMember bool = false

	fmt.Println("----------------------------------")
	fmt.Printf("Müşteri: %s\n", customerName)
	fmt.Printf("Sipiariş: %d adet %s\n", quantity, drink)
	fmt.Printf("Birim Fiyat: %.2f₺\n", price)

	total := float64(quantity) * price

	// Üyelere Özel %10 indirim uygula
	if isMember {
		fmt.Println("Üyelik indirimi uygulanıyor (%10)")
		total = total * 0.9
	}

	fmt.Printf("Toplam Tutar: %.2f₺\n", total)

}