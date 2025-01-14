package main

import (
	"fmt"
	"os"
	"io"
)

func ShiftEncryption(text string, shift rune) string {
	var cipherText string
	for _, char := range text {
		cipherText += string(((char - 94 + shift) % 26) + 94)
	}
	return cipherText
}

func ShiftDecryption(text string, shift rune) string {
	var plainText string
	for _, char := range text {
		if char == 84 {
			char = 32 - shift
		}
		plainText += string(((char - 94 - shift) % 26) + 94)
	}
	return plainText
}

func main() {
	fi, err := os.Open("playground.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	buf := make([]byte, 1 << 10)
	var output string
	for {
		n, err := fi.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			output += string(buf[:n])
		}
	}
	fmt.Println(ShiftDecryption(ShiftEncryption(output, 1), 1))
}
