package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("Valeur reçue de ch1:", num)
			case msg := <-ch2:
				fmt.Println("Valeur reçue de ch2:", msg)
			}
		}
	}()

	go func() {
		ch1 <- 10
	}()

	go func() {
		ch2 <- "Bonjour !"
	}()

	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	res, _ := http.Get("https://google.com")
	fmt.Println(res.Body)

	time.Sleep(time.Second)
}
