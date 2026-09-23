// Package lexer implements a lexical scanner that reads source input
// and produces a sequence of tokens for parsing
package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.char {
	case '=':
		tok = newToken(token.ASSIGN, lexer.char)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.char)
	case '(':
		tok = newToken(token.LPAREN, lexer.char)
	case ')':
		tok = newToken(token.RPAREN, lexer.char)
	case ',':
		tok = newToken(token.COMMA, lexer.char)
	case '+':
		tok = newToken(token.PLUS, lexer.char)
	case '{':
		tok = newToken(token.LBRACE, lexer.char)
	case '}':
		tok = newToken(token.RBRACE, lexer.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// Early return is necessary because readIdentifier calls readChar() which advances our readPosition and position field.
		// So we don't need to call readChar() after the switch statement
		if isLetter(lexer.char) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.char) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readChar()
	return tok
}

// Whitespace only acts as a seperator of tokens in our language and doesn't have meaning.
// We need to skip over whitespace entirely including newlines
func (lexer *Lexer) skipWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.readChar()
	}
}

// Reads a number and advances our lexer's position until it encounters
// a non alphanumerical character
func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.char) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Reads an identifier and advances our lexer's positions until it encounters
// a non-letter-character
func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.char) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
