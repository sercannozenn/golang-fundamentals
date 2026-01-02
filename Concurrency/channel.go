package main

import (
	"fmt"
	"time"
)



func sendMessage(ch chan string){
	time.Sleep(1 * time.Second)

	ch <- "Selam! Goroutineden mesaj geldi."
}

// Channel, goroutineler arasında veri alışverişini yapmak için kullanılır.
// Güvenli

func main(){
	ch := make(chan string)

	go sendMessage(ch)

	// Kanalı dinleme.
	msg := <-ch

	fmt.Println("Ana program mesaj aldı:", msg)

	// make([]int, 0, 1000) 

}
