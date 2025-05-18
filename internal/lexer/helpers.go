// file: internal/lexer/lexer_helpers.go
// description: Helper functions for the TRUMP language lexer

package lexer

import (
	"unicode"
	"unicode/utf8"

	"github.com/AndrewDonelson/trumplang/internal/errors"
	"github.com/AndrewDonelson/trumplang/internal/lexer/token"
)

// Read a potential "FAKE NEWS:" comment
func (l *Lexer) readFakeNewsComment() string {
	startPosition := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

// Helper function to check if the next word matches expected
func (l *Lexer) peekWord() string {
	// Save current position
	currentPos := l.position
	currentReadPos := l.readPosition
	currentCh := l.ch
	currentLine := l.line
	currentColumn := l.column

	// Read the next word
	l.readChar() // Move past the current character
	word := l.readWord()

	// Restore position
	l.position = currentPos
	l.readPosition = currentReadPos
	l.ch = currentCh
	l.line = currentLine
	l.column = currentColumn

	return word
}

// Read a complete word (identifier)
func (l *Lexer) readWord() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == ':' {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Read a multi-line comment
func (l *Lexer) readComment() string {
	startPosition := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

// Read a string literal
func (l *Lexer) readString() string {
	position := l.position + 1 // Skip the opening double quote

	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
		// Handle escape sequences
		if l.ch == '\\' {
			l.readChar() // Skip the backslash
			// Could handle specific escape sequences here
		}
	}

	if l.ch == 0 {
		errorMsg := errors.NewTrumpError(errors.UNTERMINATED_STRING, "Unterminated string", l.line, l.column)
		l.addError(errorMsg)
	}

	return l.input[position:l.position]
}

// Read a numeric literal (integer or float)
func (l *Lexer) readNumber() token.Token {
	startLine, startColumn := l.line, l.column
	position := l.position
	isFloat := false

	for isDigit(l.ch) {
		l.readChar()
	}

	// Check for decimal point
	if l.ch == '.' && isDigit(l.peekChar()) {
		isFloat = true
		l.readChar() // Consume the decimal point

		for isDigit(l.ch) {
			l.readChar()
		}
	}

	tokenType := token.TokenType(token.INT)
	if isFloat {
		tokenType = token.TokenType(token.FLOAT)
	}

	return token.Token{
		Type:    tokenType,
		Literal: l.input[position:l.position],
		Line:    startLine,
		Column:  startColumn,
	}
}

// Read an identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Skip whitespace characters
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}
		l.readChar()
	}
}

// Read the next character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII NUL or EOF
	} else {
		// Get the next rune (Unicode character)
		r, size := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.ch = r
		l.position = l.readPosition
		l.readPosition += size
		l.column++
	}
}

// Peek at the next character without advancing the lexer
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return r
}

// Add an error to the lexer's error list
func (l *Lexer) addError(err string) {
	l.errors = append(l.errors, err)
}

// Helper function to create a new token
func newToken(tokenType token.TokenType, ch rune, line, column int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: line, Column: column}
}

// Helper function to check if a character is a letter
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// Helper function to check if a character is a digit
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}