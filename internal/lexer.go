package internal

import (
        "unicode"
        "fmt"
)

type Lexer struct{
        tokens []*Token
        currPos int 
        content string
        line int
}

func NewLexer(content string) *Lexer {
        lexer := &Lexer{
                tokens : []*Token{}, 
                content : content,
                line : 1, 
        }
        lexer.parseContent()
        for _, token := range lexer.tokens{
                fmt.Println(token)
        } 
        return lexer
} 

//
// parses contents and returns the tokens in the json file . 
// this is then feed to the AST which makes sense of all the things that are present in the file 
func (l *Lexer) parseContent() []*Token{
        for l.currPos < len(l.content){
                currVal := l.content[l.currPos]
                switch currVal {
                case '[': 
                        l.addToken(NewToken(l.line, nil, LEFT_SQUARE_BRACE ))
                        l.advance()
                case ']': 
                        l.addToken(NewToken(l.line, nil, RIGHT_SQUARE_BRACE))
                        l.advance()
                case '{': 
                        l.addToken(NewToken(l.line, nil, LEFT_BRACE))
                        l.advance()
                case '}': 
                        l.addToken(NewToken(l.line, nil, RIGHT_BRACE))
                        l.advance()
                case ':': 
                        l.addToken(NewToken(l.line, nil, SEMI_COLON))
                        l.advance()
                
                case ',': 
                        l.addToken(NewToken(l.line, nil, COMMA))
                        l.advance()
                case '\n': 
                        l.line++
                        l.advance()
                case ' ': 
                        l.advance()
                
                case '\t':
                        l.advance()
                case '"': 
                        l.advance()
                        startPos := l.currPos
                        for l.currPos< len(l.content) && l.content[l.currPos] != '"'{
                                l.advance()
                        }
                        l.addToken(NewToken(l.line, l.content[startPos:l.currPos], STRING))
                        l.advance()
                default:
                        if l.isBoolean(){
                                continue 
                        } else {
                                currNum := 0 
                                for l.currPos < len(l.content) && unicode.IsDigit(rune(currVal))  {
                                        currNum = currNum*10 + int(currVal-'0') 
                                        l.advance()
                                        currVal = l.content[l.currPos]
                                }
                                l.addToken(NewToken(l.line, currNum, NUMBER))
                                l.advance()
                        }
                
        }
                //fmt.Println("lexer", l.currPos, l.line)


                        
        }
        return l.tokens
        }
       
func (l *Lexer) advance(){
        l.currPos++
}

func (l *Lexer) isBoolean() bool {
        if l.content[l.currPos:l.currPos+4] == "true" {
                l.addToken(NewToken(l.line, nil, TRUE))
                l.currPos+=4
                return true 
        } 
        if l.content[l.currPos:l.currPos+5] == "false" {
                l.addToken(NewToken(l.line, nil, FALSE))
                l.currPos+=5
                return true
        }
        return false 
}


func (l *Lexer) addToken(token *Token) {
        l.tokens = append(l.tokens, token)
}



