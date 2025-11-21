package main

import "fmt"

func main(){
	var username string = "user123"
	password := "pass123"
	const MaxLoginAttempts = 3

	fmt.Printf("Kullanıcı Adı: %s Parola: %s\nMaksimum Giriş Denemesi: %d\n", username, password, MaxLoginAttempts)
}