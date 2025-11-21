package main

import (
	"errors"
	"fmt"
)

// Error Handling
// Go'da exception kullanılmaz
// Fonksiyonlar hata döndürür.
func main(){

	result, err := Divide(10,2)
	if err != nil {
		fmt.Println("Hata oluştu:", err)
	} else{
		fmt.Println("Sonuç:", result)
	}

	_, err2 := Divide(20, 0)
	if err2 != nil {
		fmt.Println("Hata oluştu:", err2)
	}

	user1, err := GetUser(1)
	if err != nil {
		fmt.Println("Kullanıcı hatası:", err)
	}else {
		fmt.Println("Kullanıcı bulundu:", user1)
	}
	_, err = GetUser(99)
	if err != nil {
		fmt.Println("Kullanıcı hatası:", err)
	}

}

func Divide(a, b float64)(float64, error){
	if b == 0 {
		return 0, errors.New("0'a bölme hatası")
	}

	return a / b, nil
}

func GetUser(id int) (string, error){
	users := map[int]string{
		1:"Sercan",
		2:"Tuğçe",
		3:"Alin",
	}
	name, exists := users[id]
	
	if !exists {
		return "",fmt.Errorf("Kullanıcı bulunamadı: id:%d", id)
	}

	return name, nil

}