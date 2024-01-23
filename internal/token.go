package internal

import "fmt"

type TokenType rune


const (
        LEFT_BRACE TokenType = '{'
        RIGHT_BRACE TokenType = '}'
        COMMA = ','
        SEMI_COLON = ':'
        LEFT_SQUARE_BRACE = '['
        RIGHT_SQUARE_BRACE = ']'
        STRING = ' '
        NUMBER = ' '
        EOF = '\r'
        TRUE = ' '
        FALSE = ' '
)

type Token struct {
        line int 
        literal any
        tokenType TokenType
}


func (t Token) String() string{
        return fmt.Sprintf("Token[(line = %d, literal = %v , tokenType = %v)]",  t.line, t.literal, t.tokenType)
}

func NewToken(line int, literal any, tokenType TokenType ) *Token{
        return &Token{
                line : line, 
                literal: literal,
                tokenType: tokenType,
        }
}
                
