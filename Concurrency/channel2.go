package main

import (
	"fmt"
	"sync"
	"time"
)




func sendMessage(ch chan string, msg string, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(1 * time.Second)
	ch <- msg
}

func main(){
	ch := make(chan string)
	var wg sync.WaitGroup

	// 3 tane goroutine çalıştıralım
	wg.Add(3)
	go sendMessage(ch, "Mesaj 1", &wg)
	go sendMessage(ch, "Mesaj 2", &wg)
	go sendMessage(ch, "Mesaj 3", &wg)

	// Gözetmen gorutine - işçiler bitince kanalı kapatacak
	go func ()  {
		wg.Wait() // 3 goroutine tamamlanana kadar bekle
		close(ch) // kanalı kapatır
	}()

	// Tüm mesajları dinle
	for mesaj := range ch{
		fmt.Println("Aldım:", mesaj)
	}

}
