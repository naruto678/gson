package main

import (
	"flag"
	"os"
        "fmt"
	"strings"
        Internal "github.com/naruto678/gson/internal"
)



func main() {
        
        fileName := flag.String("file", "", "file to parse")
        flag.Parse()
        
        if strings.Trim(*fileName, " ")==""{
                panic("file name cannot be empty")
        }
        
        if content, err := os.ReadFile(*fileName); err!=nil {
                lexer := Internal.NewLexer(string(content))
                fmt.Println(lexer)
        } else {
                panic(err)
        }

}
