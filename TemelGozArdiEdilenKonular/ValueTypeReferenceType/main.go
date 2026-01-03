package main

import "fmt"

func main(){
	// Value Type (kopyalanır)
	// int, float, bool, string, struct, array

	// Reference Type (aynı belleği işaret eder)
	// slice, map, channel, func, pointer

	a := 10
	b := a // a'nın değeri kopyalanır.
	fmt.Println("b:", b)

	b = 20

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	u1 := User{Name: "Alin"}
	u2 := u1 // struct kopyalandı

	u2.Name = "Sercan"
	fmt.Println(u1.Name)
	fmt.Println(u2.Name)

	// Reference Type -> slice
	s1 := []int{1, 2, 3}
	s2 := s1 // aynı arrayi işaret eder

	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	c := 10
	fmt.Println(c)
	change(&c)
	fmt.Println(c)
}
func change(x *int){
	*x = 50 // pointer üzerinden asıl değeri değiştirmiş olur
}
type User struct{
	Name string
}
