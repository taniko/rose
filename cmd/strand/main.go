package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

const characters = "abcdefghijklmnopqrstuvwxyz0123456789"

func main() {
	length := flag.Int("l", 16, "length")
	flag.Parse()
	characters := strings.Split(characters, "")
	size := len(characters)

	var result []string
	for i := 0; i < *length; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(int64(size)))
		if err != nil {
			panic(err)
		}
		result = append(result, characters[b.Int64()])
	}
	fmt.Println(strings.Join(result, ""))
}
