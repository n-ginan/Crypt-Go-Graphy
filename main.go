package main

import (
	"fmt"
	"io"
	"os"
)

func ShiftEncryption(text string, shift rune) string {
	var cipherText string
	for _, char := range text {
		if char == 10 {
			cipherText += "S)...?S)...?"
			continue
		}
		cipherText += string(((char - 94 + shift) % 66) + 94)
	}
	return cipherText
}

func ShiftDecryption(text string, shift rune) string {
	var plainText string
	var position int
	for {
		if position + 12 < len(text) {
			if text[position:position+12] == "S)...?S)...?" {
				plainText += "\n"
				position += 12
				continue
			}
		}
		if rune(text[position]) == 94 + shift {
			plainText += "-"
			position++
			continue
		}
		plainText += string(((rune(text[position]) - 94 - shift) % 66) + 94)
		position++
		if position == len(text) {
			break
		}
	}
	return plainText
}

func AffineEncryption(text string, a rune, b rune) string {
	// This only works with small letters with no extra special characters like spaces and symbols
	// The formula for this is just an affine function (ax + b)
	// I should go back into this because there are still a lot to understand
	// such as the euclidian algorithm to find the gcd
	var cipherText string
	for _, char := range text {
		cipherText += string(((a * (char - 97) + b) % 26) + 97) // Continue Latur The Quick Maffs
	}
	return cipherText
}

func AffineDecryption(text string) string {
	// Its possible formulation would be (3(y - 2), 3y - 6, 3y + 20)
	var plainText string
	for _, char := range text {
		plainText += string(((3 * (char - 97) + 20) % 26) + 97)
	}
	return plainText
}

func UtilMod() int {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("An error has occurred")
		return 0
	}
	return number % 26
}

func Main() {
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
	testString := AffineEncryption("itonlyworkswithsmallletters", 9, 2)
	fmt.Print(testString)
	/*for position, _ := range testString { a = 9, b = 2
		fmt.Println(testString[position:position+12])
	}*/
}

func main() {
	Main()
}
