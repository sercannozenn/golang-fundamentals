package main

import "fmt"

func main(){

	// Mapin mantığı: map bir reference type
	// runtime da yönetiği bir hash-table vardır.
	// mapler aslında bir hashtable ve bundan türer.

	//Key-value çiftlerini depolayan, hızlı etişim sağlayan dinamik veri yapısı
	
	//Sadece *map için bellekte yer ayırır - pointer döner
	// pointer var ama işaret ettiği map değeri zero value - nil
	// Nil map'in iç hashtabel ı olmadığı için yazma yapılamaz ama okuma yapılır
	pm := new(map[string]int)

	fmt.Println("okuma (*pm) [\"a\"] =", (*pm)["a"])
	fmt.Println("okuma (*pm) [\"b\"] =", (*pm)["b"])
	fmt.Println("*pm map nil mi?", *pm == nil)

	// panic assigment fırlatır
	// (*pm)["a"] = 1

	// Peki nasıl doldururuz? 

	/**
	*pm = make(map[string]int)
	fmt.Println("*pm make map nil mi?", *pm == nil)
	
	(*pm)["a"] = 1
	(*pm)["b"] = 2
	fmt.Println("*pm değerleri:", *pm)
*/

if *pm == nil {
	*pm = make(map[string]int)
}
(*pm)["c"] = 3



}