package main

import "fmt"

func main(){

	age := 17

	if age >= 18 {
		fmt.Println("Ehliyet alabilir.")
	} else {
		fmt.Println("Henüz küçük")
	}

	if score:=45; score >=50 {
		fmt.Println("Geçti")
	} else {
		fmt.Println("Kaldı")
	}
	// Burada tanımlanan score değişkeni yalnızca if bloğunun içinde geçerlidir. Scoped Variable

	temp := 5

	if temp < 10 {
		fmt.Println("Hava soğuk")
	} else if temp < 25 {
		fmt.Println("Hava ılık")
	} else {
		fmt.Println("Hava sıcak")
	}

	// Switch case
	
	day := 5

	switch day {
	case 1:
		fmt.Println("Pazartesi")
	case 2: 
		fmt.Println("Salı")
	case 3: 
		fmt.Println("Çarşamba")
	default:
		fmt.Println("Bilinmeyen gün")
	}

	grade := "D"

	switch grade {
	case "A", "B":
		fmt.Println("Başarılı")
	case "C":
		fmt.Println("Yeterli")
	default:
		fmt.Println("Tekrar gerekli")
	}

	for i:=1; i <= 5; i++{
		fmt.Println(i)
	}
	fmt.Println("-------------------")
	//while
	count := 0
	for count <=5 {
		fmt.Println(count)
		count++
	}

	for {
		fmt.Println("Bu sürekli çalışır")
		break // çıkmak için kullanılır
	}
	fmt.Println("-------------------")
	fmt.Println("-------------------")

	for i:=1; i <=9; i++ {
		if i == 3 {
			continue // Atla demek
			// 3 ü atla
		}

		if i == 5 {
			break
		}

		fmt.Println(i)
	}
}
