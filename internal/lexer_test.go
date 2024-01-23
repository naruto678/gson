package internal

import (
	"os"
	"testing"
        "fmt"
)

func TestParse(t *testing.T){
        content, err := os.ReadFile("../test.json")
        if err!=nil {
                panic(err)
        }
        fmt.Println(string(content))
        lexer := NewLexer(string(content))
        fmt.Println(lexer.tokens)
}
