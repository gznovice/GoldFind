package main

import(
	"fmt"
	"sync"
	"time"
)

var once sync.Once

var findGold = func(){
	fmt.Print("GOLD found")
}

func doSearch(i int){
	fmt.Println("Team", i, "is searching GOLD")
	time.Sleep(time.Second)
	once.Do(findGold)
}

func main(){
	const N = 5
	var wg sync.WaitGroup
	wg.Add(N)
	for  i := 1; i <= 5; i++{		
		go func(i int){
			doSearch(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	
}