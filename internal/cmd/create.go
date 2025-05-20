// file: internal/cmd/create.go
// description: Create command for the TRUMP language

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AndrewDonelson/trumplang/internal/errors"
)

/*
 * CreateTrump creates a new TRUMP project with sample code.
 *
 * This function:
 * 1. Creates a new directory
 * 2. Creates a sample main.trump file
 * 3. Creates a README file
 */
func CreateTrump(args []string) {
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

	// Create a sample main.trump file
	sampleCode := `// file: main.trump
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
	filePath := projectName + "/main.trump"
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

    trumpc run ${projectName}

Or directly with:

    trumpc run main.trump

To build your program:

    trumpc build main.trump

This will create a main.djt file.

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
	readmeContent = strings.Replace(readmeContent, "${projectName}", projectName, -1)
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
	fmt.Println("  trumpc run", projectName)
}
