package main

import (
	"fmt"
	"sync"
	"time"
)

// Waitgroup: Goroutinlerin tamamlanmasını beklemek için kullanılan sayaçtır.
// wg.Add() -> Kaç tane goroutine çalışacak
// wg.Done() -> Bu goroutine işini bitirdi.
// wg.Wait() -> Tüm goroutineler Done diyene kadar bekle
func main(){
	fmt.Println("--- Senkron Çalışma (go olmadan) ---")
	slowJob("1. İşlem")
	slowJob("2. İşlem")

	fmt.Println("\n ********* GOROUTINE *********")

	var wg sync.WaitGroup

	wg.Add(1)
	fmt.Println("--- ASenkron Çalışma (go ile) ---")

	go slowJobAsync("3. İşlem", &wg, 500) // 1500

	go slowJobAsync("4. İşlem", &wg, 600) // 2400
	go slowJobAsync("5. İşlem", &wg, 800) // 2700
	go slowJobAsync("6. İşlem", &wg, 1000) // 3000

	wg.Wait()
}
func slowJob(name string){
	for i:=1; i<=3; i++{
		fmt.Println(name, "-> adım:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func slowJobAsync(name string, wg *sync.WaitGroup, ms int){
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		time.Sleep(time.Duration(ms)* time.Millisecond)
		fmt.Println(name, "-> adım:", i)
	}
}