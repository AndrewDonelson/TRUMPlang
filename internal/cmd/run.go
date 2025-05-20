// file: internal/cmd/run.go
// description: Run command for the TRUMP language

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AndrewDonelson/trumplang/internal/errors"
	"github.com/AndrewDonelson/trumplang/internal/interpreter"
	"github.com/AndrewDonelson/trumplang/internal/lexer"
	"github.com/AndrewDonelson/trumplang/internal/parser"
)

// RunTrump runs a Trump program
func RunTrump(args []string, verbose bool) {
	if len(args) < 1 {
		fmt.Println(errors.NewTrumpError(errors.MISSING_ARGUMENT, "Please specify a .trump file to run", 0, 0))
		os.Exit(1)
	}

	inputFile := args[0]

	// If no file specified but it's a directory, look for main.trump
	if info, err := os.Stat(inputFile); err == nil && info.IsDir() {
		mainFile := inputFile
		if !strings.HasSuffix(mainFile, "/") {
			mainFile += "/"
		}
		mainFile += "main.trump"

		if _, err := os.Stat(mainFile); err == nil {
			inputFile = mainFile
		} else {
			fmt.Println(errors.NewTrumpError(errors.FILE_NOT_FOUND, "No main.trump found in directory", 0, 0))
			os.Exit(1)
		}
	}

	// Check file extension
	if !strings.HasSuffix(inputFile, ".trump") {
		fmt.Println(errors.NewTrumpError(errors.INVALID_FILE_TYPE, "Expected a .trump file", 0, 0))
		os.Exit(1)
	}

	// Read input file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(errors.NewTrumpError(errors.FILE_NOT_FOUND, "Cannot read file: "+inputFile, 0, 0))
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

	// Create an evaluator and run the program
	evaluator := interpreter.NewEvaluator()
	result := evaluator.Eval(program)

	// Check for evaluation errors
	if result != nil && result.Type() == interpreter.ERROR_OBJ {
		fmt.Println(errors.NewTrumpError(errors.RUNTIME_ERROR, "Execution failed", 0, 0))
		fmt.Println("   ", result.Inspect())
		os.Exit(1)
	}

	if verbose && result != nil && result.Type() != interpreter.NULL_OBJ {
		fmt.Println("\nProgram result:", result.Inspect())
	}
}
