// file: internal/lexer/token/token.go
// description: Token types and related functions for the TRUMP language lexer

package token

// TokenType represents the type of token
type TokenType string

// Token represents a lexical token in the TRUMP programming language
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Token types
const (
	// Special tokens
	ILLEGAL = "ILLEGAL" // Unknown token
	EOF     = "EOF"     // End of file
	COMMENT = "COMMENT" // Comment (includes single-line and multi-line)

	// Identifiers and literals
	IDENT  = "IDENT"  // Variable or function name
	INT    = "INT"    // Integer literal
	FLOAT  = "FLOAT"  // Float literal
	STRING = "STRING" // String literal

	// Operators
	ASSIGN   = "=" // =
	PLUS     = "+" // +
	MINUS    = "-" // -
	BANG     = "!" // !
	ASTERISK = "*" // *
	SLASH    = "/" // /

	EQ     = "==" // ==
	NOT_EQ = "!=" // !=
	LT     = "<"  //
	GT     = ">"  // >
	LT_EQ  = "<=" // <=
	GT_EQ  = ">=" // >=

	// Delimiters
	COMMA     = "," // ,
	SEMICOLON = ";" // ;
	LPAREN    = "(" // (
	RPAREN    = ")" // )
	LBRACE    = "{" // {
	RBRACE    = "}" // }
	LBRACKET  = "[" // [
	RBRACKET  = "]" // ]

	// Keywords
	FUNCTION        = "FUNCTION"
	YUGE            = "YUGE"
	TREMENDOUS      = "TREMENDOUS"
	BUILD           = "BUILD"
	WALL            = "WALL"
	IF              = "IF"
	ELSE            = "ELSE"
	RETURN          = "RETURN"
	MAKE            = "MAKE"
	DEALS           = "DEALS"
	WHILE           = "WHILE"
	TWEET           = "TWEET"
	RALLY           = "RALLY"
	EXECUTIVE_ORDER = "EXECUTIVE_ORDER"
	WINNING         = "WINNING"
	LOSER           = "LOSER"
	BIGLY           = "BIGLY"
	BILLIONS        = "BILLIONS"
	FAKE_NEWS       = "FAKE_NEWS"
	RATED           = "RATED"
	BORDER          = "BORDER"
	FOR             = "FOR"
	AMERICA         = "AMERICA"
	GREAT           = "GREAT"
	AGAIN           = "AGAIN"
)

// Map of keywords to their token types
var keywords = map[string]TokenType{
	"FUNCTION":        FUNCTION,
	"YUGE":            YUGE,
	"TREMENDOUS":      TREMENDOUS,
	"BUILD":           BUILD,
	"WALL":            WALL,
	"IF":              IF,
	"ELSE":            ELSE,
	"RETURN":          RETURN,
	"MAKE":            MAKE,
	"DEALS":           DEALS,
	"WHILE":           WHILE,
	"TWEET":           TWEET,
	"RALLY":           RALLY,
	"EXECUTIVE_ORDER": EXECUTIVE_ORDER,
	"WINNING":         WINNING,
	"LOSER":           LOSER,
	"BIGLY":           BIGLY,
	"BILLIONS":        BILLIONS,
	"FAKE_NEWS":       FAKE_NEWS,
	"RATED":           RATED,
	"BORDER":          BORDER,
	"FOR":             FOR,
	"AMERICA":         AMERICA,
	"GREAT":           GREAT,
	"AGAIN":           AGAIN,
}

// LookupIdent checks if the given identifier is a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
