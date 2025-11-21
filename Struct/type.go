package main

import "fmt"

type Age int
type Email string


func main(){

	var yas Age = 31
	fmt.Println(yas)

	var eposta Email = "sercan@example.com"
	fmt.Println(eposta)

	var sayi int = 28
	
	yas = Age(sayi)
	fmt.Println(yas)

}
