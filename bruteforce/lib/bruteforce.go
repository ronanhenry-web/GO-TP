package lib

import (
	"context"
	"sync"
)

func getChars() []byte {
	return []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'_', '@',
	}
}

// generate explore récursivement les combinaisons de caractères
func generate(ctx context.Context, chars []byte, current string, target string, maxLen int, foundChan chan string) {
	// Si le contexte est annulé, on arrête
	select {
	case <-ctx.Done():
		return
	default:
	}

	// Vérification si on a trouvé le mot de passe
	if current == target {
		foundChan <- current
		return
	}

	// Si la longueur maximale est atteinte, on arrête la récursion
	if len(current) == maxLen {
		return
	}

	// Ajout récursif d'un caractère supplémentaire
	for _, c := range chars {
		// Vérification précoce en cas d'annulation
		select {
		case <-ctx.Done():
			return
		default:
		}
		generate(ctx, chars, current+string(c), target, maxLen, foundChan)
	}
}

func FindPasswordAsync(password string) string {
	chars := getChars()
	foundChan := make(chan string, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	maxLen := 5 // Longueur maximale à tester

	// On lance une goroutine pour chaque caractère de départ pour "paralléliser" la recherche
	for _, c := range chars {
		wg.Add(1)
		go func(prefix byte) {
			defer wg.Done()
			generate(ctx, chars, string(prefix), password, maxLen, foundChan)
		}(c)
	}

	// Une goroutine attend que toutes les recherches soient terminées et ferme le canal
	go func() {
		wg.Wait()
		close(foundChan)
	}()

	// Dès qu'un résultat est envoyé, on annule le contexte pour stopper les autres goroutines
	result, ok := <-foundChan
	if ok {
		cancel()
		return result
	}

	cancel()
	return ""
}

func FindPasswordNaive(password string) string {
	word := ""
	for k := 0; k < len(getChars()); k++ {
		word = string(getChars()[k])
		if word == password {
			return word
		}
		for i := range getChars() {
			word = string(getChars()[k]) + string(getChars()[i])
			if word == password {
				return word
			}
			for j := range getChars() {
				word = string(getChars()[k]) + string(getChars()[i]) + string(getChars()[j])
				if word == password {
					return word
				}
				for l := range getChars() {
					word = string(getChars()[k]) + string(getChars()[i]) + string(getChars()[j]) + string(getChars()[l])
					if word == password {
						return word
					}
					for m := range getChars() {
						word = string(getChars()[k]) + string(getChars()[i]) + string(getChars()[j]) + string(getChars()[l]) + string(getChars()[m])
						if word == password {
							return word
						}
					}
				}
			}
		}
	}

	return ""
}

func FindPasswordByCombination(password string) string {
	chars := getChars()
	n := len(chars)
	maxLen := 5

	for length := 1; length <= maxLen; length++ {
		indices := make([]int, length)

		for {
			word := make([]byte, length)
			for i, idx := range indices {
				word[i] = chars[idx]
			}

			if string(word) == password {
				return string(word)
			}

			i := length - 1
			for i >= 0 && indices[i] == n-1 {
				indices[i] = 0
				i--
			}

			if i < 0 {
				break
			}

			indices[i]++
		}
	}

	return ""
}

func FindPasswordByCombinationAsync(password string) string {
	chars := getChars()
	n := len(chars)

	wg := sync.WaitGroup{}
	maxLen := 5
	wg.Add(maxLen)
	chPassword := make(chan string)

	for length := 1; length <= maxLen; length++ {
		go func(length int) {
			defer wg.Done()
			indices := make([]int, length)

			for {
				word := make([]byte, length)
				for i, idx := range indices {
					word[i] = chars[idx]
				}

				if string(word) == password {
					chPassword <- string(word)
					return
				}

				i := length - 1
				for i >= 0 && indices[i] == n-1 {
					indices[i] = 0
					i--
				}

				if i < 0 {
					break
				}
				indices[i]++
			}
		}(length)
	}

	return <-chPassword
}
