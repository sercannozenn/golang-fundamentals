package main

import (
	"encoding/json"
	"fmt"
)

// Json => Veriyi düzenli bir şekilde saklamak ve başka sistemlere göndermek için kullanılır.

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func main(){

	jsonData := `{"name":"Sercan","age":30,"email":"sercan@example.com"}`

	// Unmarshal
	// Json -> struct 
	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("JSON çözümleme hatası:", err)
		return
	}
	fmt.Println("Json -> Struct Sonucu:")
	fmt.Println("İsim:", user.Name)
	fmt.Println("Yaş:", user.Age)
	fmt.Println("E-posta:", user.Email)

	//Marshal
	// Struct -> JSON
	newUser := User{
		Name: "Alin",
		Age: 1,
		Email: "alin@example.com",
	}

	jsonBytes, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println("Json oluşturma hatası:", err)
		return
	}
	fmt.Println("\nStruct -> Json Sonucu:")
	fmt.Println(string(jsonBytes))

	// Daha okunabilir bir json
	prettyJson, _ := json.MarshalIndent(newUser, "", " ")
	//Her satırın başına ne eklenecek
	// her seviye için kaç boşşluk bırakılacağı
	fmt.Println("Daha okunabilir json:")
	fmt.Println(string(prettyJson))


}
