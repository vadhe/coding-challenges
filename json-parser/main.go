package main

import (
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

	for _, val := range content {
		r := lexer(string([]byte{val}))
		fmt.Println(r)
	}
}

func lexer(char string) Token {
	re := Token{}
	if char == "{" {
		re.Type = LEFT_BRACE
		re.Value = "{"
	} else if char == "}" {
		re.Type = RIGHT_BRACE
		re.Value = "}"
	}
	return re
}
