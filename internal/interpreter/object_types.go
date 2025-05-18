// file: internal/interpreter/object_types.go
// description: Object type implementations for the TRUMP language interpreter

package interpreter

import (
	"fmt"
	"strings"

	"github.com/AndrewDonelson/trumplang/internal/parser"
)

// Integer represents an integer value
type Integer struct {
	Value int64
}

func (i *Integer) Type() string { return INTEGER_OBJ }
func (i *Integer) Inspect() string {
	// Easter egg: If value is 45, add a special message
	if i.Value == 45 {
		return fmt.Sprintf("%d - THE BEST NUMBER, BELIEVE ME!", i.Value)
	}
	return fmt.Sprintf("%d", i.Value)
}

// Float represents a floating-point value
type Float struct {
	Value float64
}

func (f *Float) Type() string    { return FLOAT_OBJ }
func (f *Float) Inspect() string { return fmt.Sprintf("%g", f.Value) }

// Boolean represents a boolean value (WINNING or LOSER)
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string {
	if b.Value {
		return "WINNING"
	}
	return "LOSER"
}

// String represents a string value
type String struct {
	Value string
}

func (s *String) Type() string    { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value }

// Null represents a null value
type Null struct{}

func (n *Null) Type() string    { return NULL_OBJ }
func (n *Null) Inspect() string { return "COVFEFE" }

// ReturnValue wraps a returned value
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() string    { return RETURN_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error represents an error value
type Error struct {
	Message string
}

func (e *Error) Type() string    { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Function represents a function definition
type Function struct {
	Parameters []*parser.Identifier
	Body       *parser.BlockStatement
	Env        *Environment
	Rating     string // Optional rating (e.g., "10/10")
}

func (f *Function) Type() string { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out strings.Builder

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("YUGE FUNCTION")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")

	if f.Rating != "" {
		out.WriteString("RATED ")
		out.WriteString(f.Rating)
		out.WriteString(" ")
	}

	out.WriteString(f.Body.String())

	return out.String()
}

// Builtin represents a built-in function
type Builtin struct {
	Fn func(args ...Object) Object
}

func (b *Builtin) Type() string    { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string { return "BUILT-IN FUNCTION" }

// Array represents an array value
type Array struct {
	Elements []Object
}

func (a *Array) Type() string { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out strings.Builder

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
