package internal

import "fmt"

type TokenType string

const (
        LEFT_BRACE TokenType = "{"
        RIGHT_BRACE TokenType = "}"
        COMMA = ","
        SEMI_COLON = ";"
        LEFT_SQUARE_BRACE = "["
        RIGHT_SQUARE_BRACE = "]"
        STRING = "STRING"
        NUMBER = "NUMBER"
        EOF = "EOF"
        TRUE = "TRUE"
        FALSE = "FALSE"
)

type Token struct {
        line int 
        literal any
        tokenType TokenType
}


func (t Token) String() string{
        return fmt.Sprintf("Token[(line = %d, tokenType = '%s', literal = %v)]",  t.line, t.tokenType, t.literal)
}

func NewToken(line int, literal any, tokenType TokenType ) *Token{
        return &Token{
                line : line, 
                literal: literal,
                tokenType: tokenType,
        }
}
                
