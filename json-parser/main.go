package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const LEFT_BRACE = "LEFT_BRACE"
const RIGHT_BRACE = "RIGHT_BRACE"

type Token struct {
	Type  string
	Value string
}

func main() {
	content, err := os.ReadFile("./step1/valid.json")
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	lexer(content)
}

func lexer(content []byte) []Token {

	re := []Token{}
	tempText := ""
	isString := false

	for _, val := range content {
		char := string(val)

		// 1. Deteksi tanda kutip (misal menggunakan " bukan \)
		if char == "\"" {
			if !isString {
				isString = true
				continue // Lewati tanda kutip pembuka
			} else {
				isString = false
				// Opsional: Simpan tempText ke 're' di sini sebagai satu Token String
				re = append(re, Token{Type: "STRING", Value: tempText})
				tempText = "" // Reset penampung
				continue
			}
		}

		// 2. Jika di dalam mode string, kumpulkan karakternya
		if isString {
			tempText += char
			continue // Kunci utama: lompat ke karakter berikutnya, jangan append di bawah
		}

		// 3. Logika untuk karakter normal (bukan string)
		re = append(re, Token{
			Type:  SetToken(char),
			Value: char,
		})
	}
	jsonData, _ := json.MarshalIndent(re, "", "  ")
	fmt.Println(string(jsonData))

	return re
}

func SetToken(char string) string {
	switch char {
	case "{":
		return "LEFT_BRACE"
	case "}":
		return "RIGHT_BRACE"
	case ":":
		return "COLON"
	default:
		return char
	}
}
