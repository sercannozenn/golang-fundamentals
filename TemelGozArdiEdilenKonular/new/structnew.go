package main

import "fmt"

type User struct{
	Name string
	Age int
}

func main(){
	// User struct ı için bellek ayırır
	// Alanlar zero value olur ("", 0)
	// *User (Pointer) döner
	u:= new(User)

	fmt.Println("Adres:", u)
	fmt.Println("Değerleri:", *u)
	
	u.Name = "Alin"
	u.Age = 1
	fmt.Println("Değerleri:", *u)

	

}