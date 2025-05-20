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
			tok = token.Token{Type: token.EQ, Literal: literal, Line: l.line, Column: l.column - 1}
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
			tok = token.Token{Type: token.NOT_EQ, Literal: literal, Line: l.line, Column: l.column - 1}
		} else {
			tok = newToken(token.BANG, l.ch, l.line, l.column)
		}
	case '*':
		// Check for end of multi-line comment
		if l.peekChar() == '/' {
			l.readChar()        // Skip '*'
			l.readChar()        // Skip '/'
			tok = l.NextToken() // Skip the comment end and return the next token
			return tok
		}
		tok = newToken(token.ASTERISK, l.ch, l.line, l.column)
	case '/':
		// Check for single-line comment
		if l.peekChar() == '/' {
			l.readChar() // Skip second '/'
			comment := l.readComment()
			tok = token.Token{Type: token.COMMENT, Literal: comment, Line: l.line, Column: l.column - 1}
			return tok
		}
		// Check for multi-line comment start
		if l.peekChar() == '*' {
			l.readChar() // Skip '*'
			comment := l.readMultiLineComment()
			tok = token.Token{Type: token.COMMENT, Literal: comment, Line: l.line, Column: l.column - 1}
			return tok
		}
		tok = newToken(token.SLASH, l.ch, l.line, l.column)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LT_EQ, Literal: literal, Line: l.line, Column: l.column - 1}
		} else {
			tok = newToken(token.LT, l.ch, l.line, l.column)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GT_EQ, Literal: literal, Line: l.line, Column: l.column - 1}
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
		tok.Line = l.line
		tok.Column = l.column - len(tok.Literal) - 1 // Adjust for the opening quote and string length
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			identifier := l.readIdentifier()
			tok.Literal = identifier
			// Check if the identifier is 'FAKE' followed by 'NEWS:'
			if identifier == "FAKE" && l.peekWord() == "NEWS:" {
				l.skipWhitespace() // Skip whitespace after FAKE
				l.readWord()       // consume "NEWS:"
				comment := l.readFakeNewsComment()
				tok = token.Token{Type: token.COMMENT, Literal: "FAKE NEWS: " + comment, Line: l.line, Column: l.column - len(identifier)}
				return tok
			}

			tok.Type = token.LookupIdent(identifier)
			tok.Line = l.line
			tok.Column = l.column - len(identifier)
			return tok
		} else if isDigit(l.ch) {
			numToken := l.readNumber()
			numToken.Line = l.line
			numToken.Column = l.column - len(numToken.Literal)
			return numToken
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.line, l.column)
			errorMsg := errors.NewTrumpError(errors.ILLEGAL_CHARACTER, "Illegal character found", l.line, l.column)
			l.addError(errorMsg)
		}
	}

	l.readChar()
	return tok
}
