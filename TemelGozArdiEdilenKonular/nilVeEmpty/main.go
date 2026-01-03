package main

import "fmt"

func main() {
	// nil = hiçbir şey yok, bellek adresi yok, başlatılmamış.
	// Sadece reference typelarda nil olabilir.

	// empty = var ama içi boş
	/*
	

	Önemli not: Len aynı olabilr ancak nil ve empty farklıdır.
	*/

	//nil slice
	var s1 []int
	fmt.Println(s1 == nil)
	fmt.Println(len(s1))

	// empty slice
	/*
	Bellek var
	Kullanıma hazır
	Eleman sayısı 0
	*/
	s2 := []int{}
	fmt.Println(s2 == nil)
	fmt.Println(len(s2))

	// make ile empty slice
	s3 := make([]int, 0)
	fmt.Println(s3 == nil)
	fmt.Println(len(s3))



	// nil map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	// m1["a"] = 1 // panic
	
	// empty map 
	m2 := map[string]int{}
	fmt.Println(m2 == nil)
	m2["a"] = 1

	m1 = make(map[string]int)
	m1["a"] = 1

	fmt.Println(m1)
}