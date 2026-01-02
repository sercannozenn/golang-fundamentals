package main

import (
	"fmt"
	"sync"
	"time"
)

// 1 ve 2 sıralı - senkron
func sendOrdered(ch chan string){
	ch <- "Mesaj 1"
	ch <- "Mesaj 2"
	time.Sleep(500 * time.Millisecond)
}

// 3-4-5 asenkron
func sendAsync(ch chan string, msg string, wg *sync.WaitGroup){
	defer wg.Done()
	ch <- msg
}

func main(){
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(4)

	// 1 ve 2 sıralı geliyor
	go func ()  {
		defer wg.Done()
		sendOrdered(ch)
	}()

	// 3-4-5 sırasız
	go sendAsync(ch, "Mesaj 3", &wg)
	go sendAsync(ch, "Mesaj 4", &wg)
	go sendAsync(ch, "Mesaj 5", &wg)

	go func ()  {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("Aldım:", msg)
	}
}