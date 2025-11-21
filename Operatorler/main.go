package main

import "fmt"

func main(){

	// Aritmetik Operatörler

	a := 10
	b := 4

	toplama := a + b
	cikarma := a - b
	carpma := a * b
	bolme := a / b
	modalma := a % b 

	fmt.Println("Aritmetik Operatörler")
	fmt.Println("Toplama: ", toplama)
	fmt.Println("Çıkarma: ", cikarma)
	fmt.Println("Çarpma: ", carpma)
	fmt.Println("Bölme: ", bolme)
	fmt.Println("Mod: ", modalma)

	// Karşılaştırma 
	x := 10
	y := 10

	esit := x == y 
	esitDegil := x != y
	kucuk := x < y
	buyuk := x > y
	kucukEsit := x <= y
	buyukEsit := x >= y

	fmt.Println(" Karşılaştırma Operatörleri")
	fmt.Println("Eşit mi? ", esit)
	fmt.Println("Eşit değil mi?", esitDegil)
	fmt.Println("Küçük mü?", kucuk)
	fmt.Println("Büyük mü?", buyuk)
	fmt.Println("Küçük eşit mi?", kucukEsit)
	fmt.Println("Büyük eşit mi?", buyukEsit)


	// Mantıksal Operatörler
	p := true
	q := false

	ve := p && q 
	println("VE:" ,ve)
	veya := p || q
	println("VEYA: ", veya)
	degil := !q
	println("Değil:", degil)

	// Atama Operatörleri
	num := 20
	num += 5
	fmt.Println(num)
	num -=2
	fmt.Println(num)
	num *= 3
	fmt.Println(num)
	num /= 4
	fmt.Println(num)
	num %= 3
	fmt.Println(num)

	// Arttırma ve Azaltma Operatörleri

	count := 0
	count++
	fmt.Println(count)
	count--
	fmt.Println(count)

	// Go da pre-increment pre-decrement operatörleri yoktur
	// Yani ++count veya --count yok

}