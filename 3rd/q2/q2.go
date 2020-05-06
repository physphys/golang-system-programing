package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	randText, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer randText.Close()
	io.CopyN(randText, rand.Reader, 1024)
}
