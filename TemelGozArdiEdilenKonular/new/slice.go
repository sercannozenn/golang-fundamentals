package main

import "fmt"

func main(){

	// Slice HEADER için bellek ayırır.
	// ptr, len, cap
	// ptr = nil
	// Elemanlar için HİÇ bellek ayrılmaz
	ps := new([]int)

	fmt.Println("ps adresi:", ps)
	fmt.Println("ps slice nil mi?:", *ps == nil)

	// append çalışır

	*ps = append(*ps, 1,2,3)

	fmt.Println("slice:", *ps)
	fmt.Println("len:", len(*ps), "cap:", cap(*ps))
}