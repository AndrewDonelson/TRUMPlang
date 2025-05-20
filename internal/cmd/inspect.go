// file: internal/cmd/inspect.go
// description: Inspect command for the TRUMP language

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AndrewDonelson/trumplang/internal/errors"
	"github.com/AndrewDonelson/trumplang/internal/lexer"
	"github.com/AndrewDonelson/trumplang/internal/parser"
)

/* inspectTrump analyzes a TRUMP program without compiling it.
 * It provides detailed information about:
 * - Lexical analysis of tokens
 * - Syntax analysis of the program structure
 * - Token and statement statistics
 *
 * This is great for debugging and educational purposes,
 * showing the tremendous inner workings of the language!
 */
func InspectTrump(args []string) {
	if len(args) < 1 {
		fmt.Println(errors.NewTrumpError(errors.MISSING_ARGUMENT, "Please specify a .trump file to inspect", 0, 0))
		os.Exit(1)
	}

	inputFile := args[0]

	// Check file extension
	if !strings.HasSuffix(inputFile, ".trump") {
		fmt.Println(errors.NewTrumpError(errors.INVALID_FILE_TYPE, "Expected a .trump file", 0, 0))
		os.Exit(1)
	}

	// Read input file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(errors.NewTrumpError(errors.FILE_NOT_FOUND, "Cannot read file", 0, 0))
		os.Exit(1)
	}

	// Create a lexer for the input
	l := lexer.New(string(input))

	// Check for lexer errors
	fmt.Println("INSPECTING FILE:", inputFile)
	fmt.Println("==========================")

	// Phase 1: Lexical analysis
	fmt.Println("PHASE 1: LEXICAL ANALYSIS")

	// Count tokens for stats
	tokenCount := 0
	identifierCount := 0
	keywordCount := 0
	literalCount := 0
	commentCount := 0

	// Preview the first few tokens
	previewTokens := []string{}
	previewLimit := 10

	for {
		tok := l.NextToken()
		tokenCount++

		// Count token types
		switch {
		case tok.Type == "IDENT":
			identifierCount++
		case tok.Type == "INT" || tok.Type == "FLOAT" || tok.Type == "STRING":
			literalCount++
		case tok.Type == "COMMENT":
			commentCount++
		case tok.Type != "EOF" && tok.Type != "ILLEGAL" &&
			tok.Type != ";" && tok.Type != "," && tok.Type != "(" && tok.Type != ")" &&
			tok.Type != "{" && tok.Type != "}" && tok.Type != "[" && tok.Type != "]" &&
			tok.Type != "=" && tok.Type != "+" && tok.Type != "-" && tok.Type != "*" && tok.Type != "/":
			keywordCount++
		}

		// Add to preview
		if len(previewTokens) < previewLimit {
			tokenDesc := fmt.Sprintf("%s (%s)", tok.Literal, tok.Type)
			previewTokens = append(previewTokens, tokenDesc)
		}

		if tok.Type == "EOF" {
			break
		}
	}

	// Print token preview
	fmt.Println("Token preview:")
	for i, token := range previewTokens {
		fmt.Printf("  %d. %s\n", i+1, token)
	}
	fmt.Println("  ...")

	// Print token stats
	fmt.Println("\nToken statistics:")
	fmt.Printf("  Total tokens: %d\n", tokenCount)
	fmt.Printf("  Identifiers: %d\n", identifierCount)
	fmt.Printf("  Keywords: %d\n", keywordCount)
	fmt.Printf("  Literals: %d\n", literalCount)
	fmt.Printf("  Comments: %d\n", commentCount)

	// Reset the lexer and parse the program
	l = lexer.New(string(input))
	p := parser.New(l)
	program := p.Parse()

	// Phase 2: Syntax analysis
	fmt.Println("\nPHASE 2: SYNTAX ANALYSIS")

	// Check for parsing errors
	if len(p.Errors()) > 0 {
		fmt.Println(errors.NewTrumpError(errors.SYNTAX_ERROR, "Parsing errors found", 0, 0))
		for i, err := range p.Errors() {
			fmt.Printf("  %d. %s\n", i+1, err)
		}
		os.Exit(1)
	}

	// Count statement types
	statementCount := len(program.Statements)

	// Count statement types
	letCount := 0
	returnCount := 0
	ifCount := 0
	loopCount := 0
	tweetCount := 0
	rallyCount := 0
	execOrderCount := 0
	expressionCount := 0

	// This is a basic approximation since we can't traverse the AST easily without adding methods
	for _, stmt := range program.Statements {
		stmtStr := stmt.String()

		switch {
		case strings.HasPrefix(stmtStr, "YUGE") || strings.HasPrefix(stmtStr, "TREMENDOUS"):
			letCount++
		case strings.HasPrefix(stmtStr, "RETURN"):
			returnCount++
		case strings.HasPrefix(stmtStr, "BUILD WALL IF"):
			ifCount++
		case strings.HasPrefix(stmtStr, "MAKE DEALS WHILE") || strings.HasPrefix(stmtStr, "MAKE AMERICA GREAT AGAIN FOR"):
			loopCount++
		case strings.HasPrefix(stmtStr, "TWEET"):
			tweetCount++
		case strings.HasPrefix(stmtStr, "RALLY"):
			rallyCount++
		case strings.HasPrefix(stmtStr, "EXECUTIVE_ORDER"):
			execOrderCount++
		default:
			expressionCount++
		}
	}

	// Print syntax stats
	fmt.Println("No syntax errors found!")
	fmt.Println("\nStatement statistics:")
	fmt.Printf("  Total statements: %d\n", statementCount)
	fmt.Printf("  Variable declarations: %d\n", letCount)
	fmt.Printf("  Return statements: %d\n", returnCount)
	fmt.Printf("  If statements: %d\n", ifCount)
	fmt.Printf("  Loop statements: %d\n", loopCount)
	fmt.Printf("  Tweet statements: %d\n", tweetCount)
	fmt.Printf("  Rally statements: %d\n", rallyCount)
	fmt.Printf("  Executive orders: %d\n", execOrderCount)
	fmt.Printf("  Expression statements: %d\n", expressionCount)

	// Overall assessment
	fmt.Println("\nOVERALL ASSESSMENT")
	fmt.Println("THIS CODE IS TREMENDOUS! No errors found. BELIEVE ME!")
	fmt.Println("It's ready to MAKE PROGRAMMING GREAT AGAIN!")

	if commentCount > 0 {
		fmt.Println("Great documentation with comments. SO TRANSPARENT! THE MOST TRANSPARENT EVER!")
	}

	if tweetCount > 0 {
		fmt.Println("Lots of tweets in this program. GREAT COMMUNICATION!")
	}

	if ifCount > 0 {
		fmt.Println("Building walls with conditional statements. VERY SECURE!")
	}

	if loopCount > 0 {
		fmt.Println("Making great deals in loops. THE BEST DEALS!")
	}

	fmt.Println("\nInspection complete. This program is PERFECT - like my phone call with Ukraine!")
}
