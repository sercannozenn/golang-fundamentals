package main

import "fmt"

type User struct{
	Name string
	Age int
}

func Birthday(u *User){
	u.Age++
}


func main(){
	user1 := User{Name: "Sercan", Age: 30}
	fmt.Println("Önce",user1)
	Birthday(&user1)
	fmt.Println("Sonra:", user1)
}