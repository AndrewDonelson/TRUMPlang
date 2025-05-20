// file: internal/cmd/build.go
// description: Build command for the TRUMP language

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AndrewDonelson/trumplang/internal/errors"
	"github.com/AndrewDonelson/trumplang/internal/lexer"
	"github.com/AndrewDonelson/trumplang/internal/parser"
)

/*
 * BuildTrump builds a TRUMP program from source to a .djt file.
 * It compiles the program into an abstract syntax tree representation.
 *
 * Input: A .trump source file
 * Output: A .djt executable file
 */
func BuildTrump(args []string, verbose bool, noFakeNews bool) {
	if len(args) < 1 {
		fmt.Println(errors.NewTrumpError(errors.MISSING_ARGUMENT, "Please specify a .trump file to build", 0, 0))
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

	// Parse the program
	l := lexer.New(string(input))
	p := parser.New(l)
	program := p.Parse()

	// Check for parsing errors
	if len(p.Errors()) > 0 {
		fmt.Println(errors.NewTrumpError(errors.SYNTAX_ERROR, "Parsing errors", 0, 0))
		for _, err := range p.Errors() {
			fmt.Println("   ", err)
		}
		os.Exit(1)
	}

	// Get output file name
	outputFile := strings.TrimSuffix(inputFile, ".trump") + ".djt"

	// Just write the AST as the "compiled" output
	err = os.WriteFile(outputFile, []byte(program.String()), 0644)
	if err != nil {
		fmt.Println(errors.NewTrumpError(errors.FILE_WRITE_ERROR, "Cannot write output file", 0, 0))
		os.Exit(1)
	}

	if verbose {
		fmt.Println("TREMENDOUS SUCCESS! Compiled", inputFile, "to", outputFile)
		fmt.Println("No problems found. BELIEVE ME, this code is PERFECT!")
		fmt.Println("Compilation finished in 0.5s - FASTER THAN ANYONE'S EVER SEEN BEFORE!")
	} else {
		fmt.Println("TREMENDOUS SUCCESS!", inputFile, "->", outputFile)
	}
}
