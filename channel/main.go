package main

import (
	"fmt"
)

// Fonction qui calcule la somme et envoie le résultat via un channel
func sum(numbers []int, resultChan chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	resultChan <- sum // Envoi du résultat via le channel
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int) // Création du channel

	go sum(numbers, resultChan) // Lancement de la goroutine

	result := <-resultChan // Récupération du résultat
	fmt.Println("Sum :", result)
}
