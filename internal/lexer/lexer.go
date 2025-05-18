// file: internal/lexer/lexer.go
// description: Lexical analyzer for the TRUMP programming language

package lexer

import (
	"github.com/AndrewDonelson/trumplang/internal/errors"
	"github.com/AndrewDonelson/trumplang/internal/lexer/token"
)

// Lexer represents a lexical analyzer for the TRUMP programming language
type Lexer struct {
	input        string   // Source code input
	position     int      // Current position in input (points to current char)
	readPosition int      // Current reading position in input (after current char)
	ch           rune     // Current character being examined
	line         int      // Current line number
	column       int      // Current column number
	errors       []string // Encountered errors
}

// New creates a new Lexer instance
func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
		errors: []string{},
	}
	l.readChar() // Initialize with first character
	return l
}

// Errors returns the list of errors encountered during lexing
func (l *Lexer) Errors() []string {
	return l.errors
}

// NextToken scans the next token from the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// Set token position
	tok.Line = l.line
	tok.Column = l.column

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal, Line: l.line, Column: l.column}
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.line, l.column)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch, l.line, l.column)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.line, l.column)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal, Line: l.line, Column: l.column}
		} else {
			tok = newToken(token.BANG, l.ch, l.line, l.column)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.line, l.column)
	case '/':
		if l.peekChar() == '/' {
			l.readChar() // Skip second '/'
			comment := l.readComment()
			tok = token.Token{Type: token.COMMENT, Literal: comment, Line: l.line, Column: l.column}
			return tok
		}
		tok = newToken(token.SLASH, l.ch, l.line, l.column)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LT_EQ, Literal: literal, Line: l.line, Column: l.column}
		} else {
			tok = newToken(token.LT, l.ch, l.line, l.column)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GT_EQ, Literal: literal, Line: l.line, Column: l.column}
		} else {
			tok = newToken(token.GT, l.ch, l.line, l.column)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.line, l.column)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.line, l.column)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.line, l.column)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.line, l.column)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.line, l.column)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.line, l.column)
	case '[':
		tok = newToken(token.LBRACKET, l.ch, l.line, l.column)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, l.line, l.column)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// Check if the identifier is 'FAKE' followed by 'NEWS:'
			if tok.Literal == "FAKE" && l.peekWord() == "NEWS:" {
				l.readWord() // consume "NEWS:"
				comment := l.readFakeNewsComment()
				tok = token.Token{Type: token.COMMENT, Literal: "FAKE NEWS: " + comment, Line: l.line, Column: l.column}
				return tok
			}

			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			return l.readNumber()
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.line, l.column)
			errorMsg := errors.NewTrumpError(errors.ILLEGAL_CHARACTER, "Illegal character found", l.line, l.column)
			l.addError(errorMsg)
		}
	}

	l.readChar()
	return tok
}