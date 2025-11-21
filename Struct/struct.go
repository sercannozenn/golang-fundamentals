package main

import "fmt"

// Struct: Farklı türdeki verileri bir arada tutmak için kullanılır.

type User struct {
	Name string
	Age int
	Email string
}

type Address struct{
	City string
	Country string
}

type Employee struct{
	UserInfo User
	Address Address
	Salary float64
}

func main(){
	user1 := User{
		Name: "Sercan",
		Age: 31,
		Email: "sercan@example.com",
	}

	fmt.Println("Kullanıcı Bilgileri:")
	fmt.Println("Ad:", user1.Name)
	fmt.Println("Yaş:", user1.Age)
	fmt.Println("Eposta:", user1.Email)

	var user2 User
	user2.Name = "Alin"
	user2.Age = 1
	
	fmt.Println("Kullanıcı2 Bilgileri:")
	fmt.Println("Ad:", user2.Name)
	fmt.Println("Yaş:", user2.Age)
	fmt.Println("Eposta:", user2.Email)


	employee := Employee{
		UserInfo: User{
			Name: "Ahmet",
			Age: 25,
			Email: "ahmet@example.com",
		},
		Address: Address{
			City: "İstanbul",
			Country: "Türkiye",
		},
		Salary: 45000.0,
	}

	fmt.Println("\nÇalışan Bilgileri:")
	fmt.Println("Ad:", employee.UserInfo.Name)
	fmt.Println("Yaş:", employee.UserInfo.Age)
	fmt.Println("Eposta:", employee.UserInfo.Email)
	fmt.Println("Şehir:", employee.Address.City)
	fmt.Println("Ülke:", employee.Address.Country)
	fmt.Println("Maaş:", employee.Salary)

	employees := []Employee {
		{
			UserInfo: User{"Ali",26,"ali@example.com"},
			Address: Address{"Ankara", "Türkiye"},
			Salary: 40000.0,
		},
		{
			UserInfo: User{"Zeynep",23,"zeynep@example.com"},
			Address: Address{"Erzurum", "Türkiye"},
			Salary: 50000,
		},
	}

	fmt.Println("\nTüm Çalışan Listesi:")
	for _,employee := range employees {
		/*fmt.Printf("- %s (%d) | %s, %s | Maaş: %.2f₺\n",
		e.UserInfo.Name,
		e.UserInfo.Age,
		e.Address.City,
		e.Address.Country,
		e.Salary)*/
		PrintEmployee(employee)
	}

}

func PrintEmployee(e Employee){
	fmt.Println("-------------------------")
	fmt.Printf("- %s (%d) | %s, %s | Maaş: %.2f₺\n",
		e.UserInfo.Name,
		e.UserInfo.Age,
		e.Address.City,
		e.Address.Country,
		e.Salary)
}
