// file: internal/lexer/helpers.go
// description: Helper functions for the TRUMP language lexer

package lexer

import (
	"strings"
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
	word := ""
	l.skipWhitespace()
	if isLetter(l.ch) {
		word = l.readIdentifier()
	}

	// Restore position
	l.position = currentPos
	l.readPosition = currentReadPos
	l.ch = currentCh
	l.line = currentLine
	l.column = currentColumn

	return word
}

// Read a word that might include colons (for FAKE NEWS:)
func (l *Lexer) readWord() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == ':' {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Read a single-line comment
func (l *Lexer) readComment() string {
	startPosition := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

// Read a multi-line comment
func (l *Lexer) readMultiLineComment() string {
	startPosition := l.position
	startLine := l.line
	startColumn := l.column

	for {
		l.readChar()
		// Check for end of comment
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar() // consume '*'
			l.readChar() // consume '/'
			break
		}

		// Check for unterminated comment
		if l.ch == 0 {
			errorMsg := errors.NewTrumpError(errors.SYNTAX_ERROR, "Unterminated multi-line comment", startLine, startColumn)
			l.addError(errorMsg)
			break
		}

		// Handle newlines in multi-line comments
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}
	}

	// Extract the comment excluding the /* and */ markers
	comment := l.input[startPosition : l.position-2]

	// Format multi-line comment
	return formatMultiLineComment(comment)
}

// Format a multi-line comment to clean up any common patterns like
// /* This
//   - is
//   - a comment
//     */
func formatMultiLineComment(comment string) string {
	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		// Trim spaces at beginning and end
		line = strings.TrimSpace(line)
		// Remove common * prefix from lines
		if strings.HasPrefix(line, "*") {
			line = strings.TrimPrefix(line, "*")
			line = strings.TrimSpace(line)
		}
		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

// Read a string literal
func (l *Lexer) readString() string {
	position := l.position + 1 // Skip the opening double quote
	startLine := l.line
	startColumn := l.column

	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
		// Handle escape sequences
		if l.ch == '\\' {
			l.readChar() // Skip the backslash
		}

		// Handle newlines in strings
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}
	}

	if l.ch == 0 {
		errorMsg := errors.NewTrumpError(errors.UNTERMINATED_STRING, "Unterminated string", startLine, startColumn)
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

	tokenType := token.INT
	if isFloat {
		tokenType = token.FLOAT
	}

	return token.Token{
		Type:    token.TokenType(tokenType),
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
