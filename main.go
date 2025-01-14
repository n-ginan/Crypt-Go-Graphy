package main

import (
	"fmt"
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

func AffineEncryption(text string, a int, b int) string {
	var plainText string
	for _, char := range text {
		plainText += string(((rune(a) * (char - 94) + rune(b) % 66) + 94)) // Continue Latur The Quick Maffs
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
	/*fi, err := os.Open("playground.txt")
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
	}*/
	//testString := ShiftDecryption(ShiftEncryption(output, 2), 2)
	//fmt.Print(testString)
	/*for position, _ := range testString {
		fmt.Println(testString[position:position+12])
	}*/
}

func main() {
	//Main()
	fmt.Println(AffineEncryption("affine", 9, 2))
}
