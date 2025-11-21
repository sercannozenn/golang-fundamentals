package main

import "fmt"

func main(){

	fmt.Println("Gün Sonu Kafe Raporu")
	fmt.Println("----------------------")

	coffeePrice := 140.90
	cakePrice := 225.90
	coffeeSold := 37
	cakeSold := 20
	
	totalCoffeeIncome := float64(coffeeSold) * coffeePrice
	totalCakeIncome := float64(cakeSold) * cakePrice
	totalIncome := totalCoffeeIncome + totalCakeIncome

	fmt.Printf("Kahve Geliri: %.2f₺\n", totalCoffeeIncome)
	fmt.Printf("Pasta Geliri: %.2f₺\n", totalCakeIncome)
	fmt.Printf("Toplam Gelir: %.2f₺\n", totalIncome)

	fmt.Println("-------------------------")

	rent := 1250.0
	rent += 500 // PAsta ve kahve gideri
	rent -= 100 // indiirm

	fmt.Printf("Günlük Gider (kira + masraf): %.2f₺\n", rent)
	fmt.Println("-------------------------")

	isWeekendCampaign := true
	hasMembershipCampaign := false

	fmt.Println("Hafta sonu ve üyelere özel kapmanya var mı?", isWeekendCampaign && hasMembershipCampaign)
	fmt.Println("Hafta sonu veya üyelere özel kampanya var mı?", isWeekendCampaign || hasMembershipCampaign)
	fmt.Println("Hafta sonu değili", !isWeekendCampaign)

	// Kasadaki bozuk para sayma
	coins := 23
	remainder := coins % 5
	fmt.Printf("Kasada 5'lik bozuklardan kalan %d adet\n", remainder)

}