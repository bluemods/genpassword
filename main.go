package main

import (
	"flag"
	"fmt"
	"math/big"
	"slices"

	"crypto/rand" // Uses secure random generator managed by the OS
)

var (
	alphaNumChars            = []byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`)
	alphaNumCharsWithSymbols = slices.Concat(alphaNumChars, []byte(`@!$#^-_=+/\`))
)

func main() {
	passwordSize := flag.Int("s", 32, "password size")
	alphaNum := flag.Bool("a", false, "alphanumeric only")
	flag.Parse()

	if *passwordSize < 8 {
		fmt.Println("passwordSize cannot be less than 8")
		return
	}
	fmt.Println(string(randChars(*passwordSize, *alphaNum)))
}

func randChars(passwordSize int, alphaNum bool) []byte {
	var chars []byte
	if alphaNum {
		chars = alphaNumChars
	} else {
		chars = alphaNumCharsWithSymbols
	}
	for range 1000 {
		shuffle(chars)
	}

	output := make([]byte, passwordSize)
	for i := range passwordSize {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err)
		}
		output[i] = chars[n.Int64()]
	}
	return output
}

func shuffle[T any](items []T) {
	for i := len(items) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			panic(err)
		}
		jj := j.Int64()
		items[i], items[jj] = items[jj], items[i]
	}
}
