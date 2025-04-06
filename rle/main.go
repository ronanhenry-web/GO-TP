package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	original := generateBytesSlice(1024)

	workers := runtime.NumCPU()

	sem := make(chan struct{}, workers)
	compressed := make([][]byte, workers)
	for i := 0; i < workers; i++ {
		sem <- struct{}{}
		slice := original[i*len(original)/workers : (i+1)*len(original)/workers]
		i := i
		go func(id int) {
			defer func() { <-sem }()
			compressed[i] = CompressByString(slice)
		}(i)
	}

	for i := 0; i < workers; i++ {
		sem <- struct{}{}
	}

	compressedBytes := make([]byte, 0, len(original))
	for i := 0; i < workers; i++ {
		compressedBytes = append(compressedBytes, compressed[i]...)
	}

	decompressed := make([]byte, 0, len(original))
	for i := 0; i < workers; i++ {
		decompressed = append(decompressed, Decompress(compressed[i])...)
	}

	fmt.Println("Original : ", string(original))
	fmt.Println("Compressed : ", string(compressedBytes))
	fmt.Println("Equal : ", bytes.Equal(original, decompressed))
	// Calculate the compression ration as a percentage for the original and compressed data
	fmt.Printf("Compression ratio : %.2f %% \n", float64(len(original))/float64(len(compressedBytes))*100-100)
}

func Decompress(compressed []byte) []byte {
	var result bytes.Buffer
	for len(compressed) > 0 {
		letterIndex := strings.IndexFunc(string(compressed), func(r rune) bool { return !unicode.IsDigit(r) })
		multiply := 1
		if letterIndex != 0 {
			multiply, _ = strconv.Atoi(string(compressed[:letterIndex]))
		}
		result.WriteString(strings.Repeat(string(compressed[letterIndex]), multiply))
		compressed = compressed[letterIndex+1:]
	}
	return result.Bytes()
}

func CompressByString(original []byte) []byte {
	result := bytes.Buffer{}
	for len(original) > 0 {
		firstLetter := original[0]
		originalLength := len(original)
		original = []byte(strings.TrimLeft(string(original), string(firstLetter)))
		if counter := originalLength - len(original); counter > 1 {
			result.WriteString(strconv.Itoa(counter))
		}
		result.WriteString(string(firstLetter))
	}
	return result.Bytes()
}

func CompressOptimized(original []byte) []byte {
	result := bytes.Buffer{}
	count := 1
	for i := range original {
		if i == len(original)-1 {
			if count > 1 {
				result.Write([]byte(strconv.Itoa(count)))
			}
			result.Write([]byte{original[i]})
			break
		}

		if original[i] != original[i+1] {
			if count > 1 {
				result.Write([]byte(strconv.Itoa(count)))
			}
			result.Write([]byte{original[i]})
			count = 1
		} else {
			count++
		}
	}
	return result.Bytes()
}

func generateBytesSlice(size int) []byte {
	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)
	const minAmount = 0
	const maxAmount = 1

	defaultCharsets := []byte{'a', 'b'}
	bytesSlice := make([]byte, 0, size)
	for len(bytesSlice) < size {
		bytesSlice = append(bytesSlice, defaultCharsets[random.Intn(maxAmount-minAmount+1)+minAmount])
	}
	return bytesSlice
}
