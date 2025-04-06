package main

import (
	"fmt"
	"sync"
	"time"
)

type Stock struct {
	Name  string
	Value float64
}

// Nasdaq is a stock exchange in the US
func Nasdaq(chanNasdaq chan Stock, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	chanNasdaq <- Stock{
		Name:  "Nasdaq",
		Value: 6.2487,
	}
}

// CAC40 is a stock exchange in France
func CAC40(chanCAC40 chan Stock, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(300 * time.Millisecond)
	chanCAC40 <- Stock{
		Name:  "CAC40",
		Value: 1.8898,
	}
}

// Nikkei is a stock exchange in Japan
func Nikkei(chanNikkei chan Stock, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)
	chanNikkei <- Stock{
		Name:  "Nikkei",
		Value: 5.8874,
	}
}
func main() {
	stocks := make(chan Stock, 3)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go Nasdaq(stocks, &wg)
	go CAC40(stocks, &wg)
	go Nikkei(stocks, &wg)

	for i := 0; i < 3; i++ {
		stock := <-stocks
		fmt.Println(stock)
	}

	wg.Wait()
}
