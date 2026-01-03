package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	users := []string{"Ali", "Ahmet", "Fatma"}
	var wg sync.WaitGroup
	wg.Add(len(users))

	for _, u := range users{

		go func (userName string)  {
			wg.Done()

			fmt.Println(userName)
		}(u)
	}
		time.Sleep(1 * time.Second)
}
