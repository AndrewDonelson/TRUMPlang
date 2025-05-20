// file: internal/cmd/usage.go
// description: Command-line usage information for TRUMP

package cmd

import (
	"fmt"
)

// PrintUsage displays the command-line usage information
func PrintUsage() {
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
