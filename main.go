package main

import (
	"flag"
	"os"
	"strings"

	Internal "github.com/naruto678/gson/internal"
)

func main() {

	fileName := flag.String("file", "", "file to parse")
	flag.Parse()

	if strings.Trim(*fileName, " ") == "" {
		panic("file name cannot be empty you mother fucking cunt")
	}

	if content, err := os.ReadFile(*fileName); err == nil {
		_ = Internal.NewLexer(string(content))
	
	} else {
		panic(err)
	}

}
