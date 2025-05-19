// file: main.go
// description: Main entry point for the TRUMP language compiler and interpreter - updated to show multi-line comments

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/AndrewDonelson/trumplang/internal/errors"
	"github.com/AndrewDonelson/trumplang/internal/interpreter"
	"github.com/AndrewDonelson/trumplang/internal/lexer"
	"github.com/AndrewDonelson/trumplang/internal/parser"
)

func main() {
	// Define command-line flags
	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	inspectCmd := flag.NewFlagSet("inspect", flag.ExitOnError)

	// Add verbosity flags
	buildVerbose := buildCmd.Bool("verbose", false, "Enable verbose output")
	runVerbose := runCmd.Bool("verbose", false, "Enable verbose output")
	buildNoFakeNews := buildCmd.Bool("no-fake-news", false, "Suppress warnings")

	// Check for correct number of arguments
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Parse the command
	switch os.Args[1] {
	case "build":
		buildCmd.Parse(os.Args[2:])
		buildTrump(buildCmd.Args(), *buildVerbose, *buildNoFakeNews)
	case "run":
		runCmd.Parse(os.Args[2:])
		runTrump(runCmd.Args(), *runVerbose)
	case "create":
		createCmd.Parse(os.Args[2:])
		createTrump(createCmd.Args())
	case "inspect":
		inspectCmd.Parse(os.Args[2:])
		inspectTrump(inspectCmd.Args())
	default:
		printUsage()
		os.Exit(1)
	}
}

// Print command-line usage
func printUsage() {
	fmt.Println("Usage: trumpc <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  build <file.trump>    - Compile a .trump file to a .djt executable")
	fmt.Println("  run <file.trump>      - Run a .trump file directly")
	fmt.Println("  create <project>      - Create a new Trump project")
	fmt.Println("  inspect <file.trump>  - Check a program for errors without compiling")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --verbose             - Enable verbose output")
	fmt.Println("  --no-fake-news        - Suppress warnings")
}

/*
 * buildTrump builds a TRUMP program from source to a .djt file.
 * It compiles the program into an abstract syntax tree representation.
 *
 * Input: A .trump source file
 * Output: A .djt executable file
 */
func buildTrump(args []string, verbose bool, noFakeNews bool) {
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

// Run a Trump program
func runTrump(args []string, verbose bool) {
	if len(args) < 1 {
		fmt.Println(errors.NewTrumpError(errors.MISSING_ARGUMENT, "Please specify a .trump file to run", 0, 0))
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

/*
 * createTrump creates a new TRUMP project with sample code.
 *
 * This function:
 * 1. Creates a new directory
 * 2. Creates a sample hello_world.trump file
 * 3. Creates a README file
 */
func createTrump(args []string) {
	if len(args) < 1 {
		fmt.Println(errors.NewTrumpError(errors.MISSING_ARGUMENT, "Please specify a project name", 0, 0))
		os.Exit(1)
	}

	projectName := args[0]

	// Create project directory
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Println(errors.NewTrumpError(errors.FILE_WRITE_ERROR, "Cannot create project directory", 0, 0))
		os.Exit(1)
	}

	// Create a simple "Hello, World!" program
	sampleCode := `// file: hello_world.trump
// A tremendous TRUMP language program

/* 
 * This is a multiline comment in TRUMP language!
 * The best language, believe me.
 * Nobody codes better than this.
 */

// Say hello to the world, Trump style
TWEET "Hello, World! It's going to be TREMENDOUS!";

/* As everyone knows,
 * TRUMP language has the best functions
 * Much better than any other language!
 */
// Define a yuge function with a tremendous rating
YUGE FUNCTION greet(name) RATED 10/10 {
    TWEET "Hello, " + name + "! You're doing a FANTASTIC job!";
    RETURN "Greeted " + name;
}

// Call our function
greet("America");

// Use a BUILD WALL IF statement
YUGE value = 45;
BUILD WALL IF (value == 45) {
    RALLY "That's the BEST number, BELIEVE ME!";
} ELSE {
    TWEET "Not a great number. SAD!";
}

/* Make deals in a loop
 * The best loops, tremendous loops
 */
YUGE counter = 0;
MAKE DEALS WHILE (counter < 3) {
    TWEET "Making deal #" + counter;
    counter = counter + 1;
}

// For loop to make America great again
MAKE AMERICA GREAT AGAIN FOR (YUGE i = 0; i < 3; i = i + 1) {
    TWEET "Making America Great Again: " + i;
}

// Using some special built-in functions
YUGE numbers = [3, 1, 2, 45, 6];
TWEET "Original array: " + numbers;
TWEET "After TREMENDOUS_SORT: " + TREMENDOUS_SORT(numbers);
TWEET "After AMERICA_FIRST: " + AMERICA_FIRST(numbers);
`

	// Write the sample program to a file
	filePath := projectName + "/hello_world.trump"
	err = os.WriteFile(filePath, []byte(sampleCode), 0644)
	if err != nil {
		fmt.Println(errors.NewTrumpError(errors.FILE_WRITE_ERROR, "Cannot write sample program", 0, 0))
		os.Exit(1)
	}

	// Create a simple readme file
	readmeContent := `# ${projectName}

A tremendous project written in the TRUMP programming language!

## Running your program

To run your program:

    trumpc run hello_world.trump

To build your program:

    trumpc build hello_world.trump

This will create a hello_world.djt file.

## Language Features

- Variable declarations with YUGE and TREMENDOUS
- Function definitions with FUNCTION
- Conditionals with BUILD WALL IF/ELSE
- Loops with MAKE DEALS WHILE and MAKE AMERICA GREAT AGAIN FOR
- Output with TWEET, RALLY, and EXECUTIVE_ORDER
- Special built-in functions like TREMENDOUS_SORT and AMERICA_FIRST
- Single-line comments with //
- Multi-line comments with /* ... */

Enjoy making your code great again!
`
	readmeContent = strings.Replace(readmeContent, "${projectName}", projectName, 1)
	readmePath := projectName + "/README.md"
	err = os.WriteFile(readmePath, []byte(readmeContent), 0644)
	if err != nil {
		fmt.Println(errors.NewTrumpError(errors.FILE_WRITE_ERROR, "Cannot write README file", 0, 0))
		os.Exit(1)
	}

	fmt.Println("TREMENDOUS SUCCESS! Created new project:", projectName)
	fmt.Println("  - Sample program:", filePath)
	fmt.Println("  - README:", readmePath)
	fmt.Println("\nTo run your program:")
	fmt.Println("  trumpc run", filePath)
}

/* inspectTrump analyzes a TRUMP program without compiling it.
 * It provides detailed information about:
 * - Lexical analysis of tokens
 * - Syntax analysis of the program structure
 * - Token and statement statistics
 *
 * This is great for debugging and educational purposes,
 * showing the tremendous inner workings of the language!
 */
func inspectTrump(args []string) {
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
