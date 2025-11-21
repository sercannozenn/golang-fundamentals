package main

import "fmt"

func main(){

	fmt.Println("Fastfood Fün Sonu Uygulaması")
	fmt.Println("-----------------------------")

	burgerPRice := 220.40
	friesPrice := 120.0
	drinkPrice := 45.90

	var burgerCount, friesCount, drinkCount int
	
	fmt.Print("Satılan burger adedi:")
	fmt.Scanln(&burgerCount)

	fmt.Print("Satılan patates adedi:")
	fmt.Scanln(&friesCount)

	fmt.Print("Satılan içecek adedi:")
	fmt.Scanln(&drinkCount)

	total := (float64(burgerCount) * burgerPRice) + (float64(friesCount) * friesPrice) + (float64(drinkCount) * drinkPrice)
	fmt.Println("--------------------------------")

	fmt.Printf("Toplam Kazanç: %.2f₺\n", total)

	if total > 1000 {
		fmt.Println("Harika Satış Günü")
	} else if total >= 500 {
		fmt.Println("İyi satış günü")
	} else {
		fmt.Println("Bugün satışlar düşük")
	}
	fmt.Println("--------------------------------")

	fmt.Println("Kapanış işlemi başlıyor:")
	for i := 3; i > 0; i-- {
		fmt.Printf("%d...\n", i)
	}

	fmt.Println("KAsa kapandı. İyi akşamlar")

}