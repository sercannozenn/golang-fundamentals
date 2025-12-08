package main

import (
	"fmt"
	"time"
)

// Eşzamanlılık
func main() {
	slowJob("1. İşlem")
	slowJob("2. İşlem")

	// Goroutine nedir?
	// Normal bir fonksiyon tek başına sırayla çalışır.(senkron)
	// Goroutine fonksiyonu arka planda asenkron çalıştırır.

	go slowJob("3. İşlem")
	go slowJob("4. İşlem")

	time.Sleep(3 * time.Second)
}

func slowJob(name string){
	for i:=1; i<=3; i++{
		fmt.Println(name, "-> adım:", i)
		time.Sleep(500 * time.Millisecond)
	}
}
