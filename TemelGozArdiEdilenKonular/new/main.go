package main

import "fmt"

func main(){
	// new (int)
	// Heap'te int için bellek ayrılır
	// int in zero value'su atanır. 0
	// Geriye *int (pointer) döner

	p := new(int)
	fmt.Println("Pointer adresi", p)
	fmt.Println("P değeri:", *p)
	
	*p = 44
	fmt.Println("P değeri:", *p)
}
