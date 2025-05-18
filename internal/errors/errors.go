// file: internal/errors/errors.go
// description: Error handling for the TRUMP programming language

package errors

import (
	"fmt"
	"math/rand"
	"time"
)

// Initialize random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Error types
const (
	// Lexer errors
	ILLEGAL_CHARACTER   = "ILLEGAL_CHARACTER"
	UNTERMINATED_STRING = "UNTERMINATED_STRING"

	// Parser errors
	UNEXPECTED_TOKEN    = "UNEXPECTED_TOKEN"
	EXPECTED_IDENTIFIER = "EXPECTED_IDENTIFIER"
	EXPECTED_EXPRESSION = "EXPECTED_EXPRESSION"
	SYNTAX_ERROR        = "SYNTAX_ERROR"

	// File system errors
	FILE_NOT_FOUND    = "FILE_NOT_FOUND"
	FILE_WRITE_ERROR  = "FILE_WRITE_ERROR"
	PERMISSION_DENIED = "PERMISSION_DENIED"
	PAYLOAD_TOO_LARGE = "PAYLOAD_TOO_LARGE"
	INVALID_FILE_TYPE = "INVALID_FILE_TYPE"
	MISSING_ARGUMENT  = "MISSING_ARGUMENT"

	// Runtime errors
	NULL_POINTER   = "NULL_POINTER"
	STACK_OVERFLOW = "STACK_OVERFLOW"
	OUT_OF_MEMORY  = "OUT_OF_MEMORY"
	RUNTIME_ERROR  = "RUNTIME_ERROR"

	// Mathematical errors
	DIVISION_BY_ZERO     = "DIVISION_BY_ZERO"
	FLOATING_POINT_ERROR = "FLOATING_POINT_ERROR"
	INTEGER_OVERFLOW     = "INTEGER_OVERFLOW"
)

// TrumpError represents a custom error in TRUMP language
type TrumpError struct {
	Code    string
	Message string
	Line    int
	Column  int
}

// String returns a formatted error message
func (e *TrumpError) String() string {
	return fmt.Sprintf("%s at %d:%d: %s", e.Code, e.Line, e.Column, e.Message)
}

// Error implements the error interface
func (e *TrumpError) Error() string {
	return e.String()
}

// Map of error types to Trump-themed error messages
var errorMessages = map[string][]string{
	// Lexer errors
	ILLEGAL_CHARACTER: {
		"BAD HOMBRE! Illegal character found at position %d:%d. We need EXTREME VETTING for all code!",
		"YOUR CODE HAS AN INVADER at %d:%d! This is why we need a WALL around our syntax!",
		"FAKE CHARACTER! This symbol at %d:%d is written like it came from CHINA!",
	},
	UNTERMINATED_STRING: {
		"SAD! Unterminated string at %d:%d. You open quotes but don't close them. TOTAL CHAOS!",
		"YOUR QUOTES ARE A DISASTER at %d:%d! People say I know the best quotes. These are terrible quotes!",
		"WHERE'S THE CLOSING QUOTE? at %d:%d This is worse than the NATIONAL DEBT!",
	},

	// Parser errors
	UNEXPECTED_TOKEN: {
		"WRONG TOKEN! It's not what anyone expected at %d:%d. It's like showing up to a DEAL with NO LEVERAGE!",
		"UNEXPECTED TOKEN! This token at %d:%d is a TOTAL LOSER! We only accept the BEST tokens.",
		"YOUR TOKEN IS FIRED at %d:%d! It's not doing the job. We need TREMENDOUS tokens!",
	},
	EXPECTED_IDENTIFIER: {
		"WHERE'S THE NAME? Expected an identifier at %d:%d. You can't make AMERICA GREAT without proper NAMING!",
		"NO IDENTIFIER FOUND at %d:%d! How can we know what we're talking about? It's like the FAKE MEDIA!",
		"NAMELESS DISASTER at %d:%d! We need to identify variables. I identify the BEST variables!",
	},
	SYNTAX_ERROR: {
		"FAKE SYNTAX! Your code is a MESS at %d:%d. I've seen better syntax from a CHILD!",
		"YOUR SYNTAX IS TERRIBLE! at %d:%d Whoever taught you to code should be FIRED immediately!",
		"THIS SYNTAX IS A DISASTER at %d:%d. You need to get the BEST people to help you with your code!",
	},

	// File system errors
	FILE_NOT_FOUND: {
		"404 JUSTICE NOT FOUND! The file at %d:%d has been TOTALLY WIPED, like Hillary's emails!",
		"WHERE'S THE FILE? at %d:%d It's gone missing, just like those 30,000 emails!",
		"YOUR FILE IS A HOAX at %d:%d! It claims to exist, but nobody can find it. SAD!",
	},
	FILE_WRITE_ERROR: {
		"CAN'T WRITE TO FILE! at %d:%d The system is RIGGED against your code!",
		"FILE WRITE FAILURE! at %d:%d Your file system is CORRUPT like the deep state!",
		"WRITE DISASTER! at %d:%d Can't save your beautiful code. It's a WITCH HUNT!",
	},
	INVALID_FILE_TYPE: {
		"BAD FILE! at %d:%d We only accept the BEST file types, like .trump!",
		"WRONG FILE TYPE! at %d:%d That's not a Trump file. It's probably from CHINA!",
		"FILE TYPE REJECTED! at %d:%d This is why we need EXTREME VETTING for all files!",
	},
	MISSING_ARGUMENT: {
		"DISASTER DETECTED! at %d:%d You can't run with no arguments. That's like making a deal with NO LEVERAGE!",
		"WHERE ARE THE ARGUMENTS? at %d:%d You can't make AMERICA GREAT without proper INPUTS!",
		"NO ARGUMENTS FOUND! at %d:%d This command needs arguments. I know arguments, I have the BEST arguments!",
	},

	// Runtime errors
	NULL_POINTER: {
		"EMPTY PROMISES DETECTED at %d:%d! You referenced something that doesn't exist, like the OPPOSITION'S HEALTHCARE PLAN!",
		"NULL DETECTED at %d:%d! You're pointing at NOTHING. It's like the DEMOCRATS' BORDER POLICY!",
		"YOUR POINTER IS BANKRUPT at %d:%d! It's got NOTHING. ZERO. Like my opponents' ideas!",
	},
	RUNTIME_ERROR: {
		"RUNTIME DISASTER! at %d:%d Your program is CRASHING like the ECONOMY under the previous administration!",
		"EXECUTION FAILURE! at %d:%d This program is falling apart faster than FAKE NEWS ratings!",
		"PROGRAM COLLAPSE! at %d:%d Your code is having a MELTDOWN like the liberal media when I tweet!",
	},

	// Mathematical errors
	DIVISION_BY_ZERO: {
		"IMPOSSIBLE DEAL at %d:%d! You can't divide by zero, that's like trying to get MEXICO to pay without leverage!",
		"ZERO DIVISION DISASTER at %d:%d! Even I, with the best brain, know you can't divide by zero!",
		"MATH CATASTROPHE at %d:%d! Dividing by zero? That's worse economics than OBAMA!",
	},
}

// Custom countries to blame in errors
var countriesBlamed = []string{
	"CHINA",
	"MEXICO",
	"RUSSIA",
	"IRAN",
	"NORTH KOREA",
	"VENEZUELA",
}

// NewTrumpError creates a new TrumpError with a random Trump-themed message
func NewTrumpError(code, defaultMessage string, line, column int) string {
	messages, ok := errorMessages[code]
	if !ok {
		// If no specific message for this error code, use default with random enhancement
		return fmt.Sprintf("%s at %d:%d! %s! %s is probably behind this!",
			defaultMessage, line, column,
			getRandomSuperlative(),
			getRandomCountryToBlame())
	}

	// Select a random message from the available ones
	message := messages[rand.Intn(len(messages))]

	// Format the message with line and column
	return fmt.Sprintf(message, line, column)
}

// NewTrumpErrorObj creates a new TrumpError object with a random Trump-themed message
func NewTrumpErrorObj(code, defaultMessage string, line, column int) *TrumpError {
	messages, ok := errorMessages[code]
	if !ok {
		// If no specific message for this error code, use default with random enhancement
		return &TrumpError{
			Code: code,
			Message: fmt.Sprintf("%s! %s is probably behind this!",
				getRandomSuperlative(),
				getRandomCountryToBlame()),
			Line:   line,
			Column: column,
		}
	}

	// Select a random message from the available ones
	message := messages[rand.Intn(len(messages))]

	// Format the message with line and column
	return &TrumpError{
		Code:    code,
		Message: fmt.Sprintf(message, line, column),
		Line:    line,
		Column:  column,
	}
}

// Add random Trump-like hyperbole to errors
func getRandomSuperlative() string {
	superlatives := []string{
		"DISASTER",
		"TERRIBLE",
		"WORST EVER",
		"COMPLETE CATASTROPHE",
		"TOTAL MESS",
		"NOBODY'S EVER SEEN ANYTHING LIKE IT",
		"EVERYONE AGREES IT'S BAD",
		"BELIEVE ME",
		"SO SAD",
	}

	return superlatives[rand.Intn(len(superlatives))]
}

// Get a random country to blame
func getRandomCountryToBlame() string {
	return countriesBlamed[rand.Intn(len(countriesBlamed))]
}
