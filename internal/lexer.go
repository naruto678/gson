package internal

import (
        "unicode"
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
                        l.addToken(NewToken(l.line, currVal, LEFT_SQUARE_BRACE ))
                        l.advance()
                case ']': 
                        l.addToken(NewToken(l.line, currVal, RIGHT_SQUARE_BRACE))
                        l.advance()
                case '{': 
                        l.addToken(NewToken(l.line, currVal, LEFT_BRACE))
                        l.advance()
                case '}': 
                        l.addToken(NewToken(l.line, currVal, RIGHT_BRACE))
                        l.advance()
                case ':': 
                        l.addToken(NewToken(l.line, currVal, SEMI_COLON))
                        l.advance()
                
                case ',': 
                        l.addToken(NewToken(l.line, currVal, COMMA))
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
                        l.addToken(NewToken(l.line, l.content[startPos:l.currPos-1], STRING))
                        l.advance()
                default: 
                        if l.isBoolean(){
                                l.currPos+=4 
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
                l.addToken(NewToken(l.line, "true", TRUE))
                return true 
        } 
        if l.content[l.currPos:l.currPos+4] == "false" {
                l.addToken(NewToken(l.line, "false", FALSE))
                return true 
        }
        return false 
}


func (l *Lexer) addToken(token *Token) {
        l.tokens = append(l.tokens, token)
}



