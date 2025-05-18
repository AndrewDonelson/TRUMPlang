// file: internal/interpreter/object.go
// description: Object interface and environment for the TRUMP language interpreter

package interpreter

// Object types
const (
	INTEGER_OBJ  = "INTEGER"
	FLOAT_OBJ    = "FLOAT"
	BOOLEAN_OBJ  = "BOOLEAN"
	STRING_OBJ   = "STRING"
	NULL_OBJ     = "NULL"
	RETURN_OBJ   = "RETURN"
	ERROR_OBJ    = "ERROR"
	FUNCTION_OBJ = "FUNCTION"
	BUILTIN_OBJ  = "BUILTIN"
	ARRAY_OBJ    = "ARRAY"
)

// Object interface that all objects implement
type Object interface {
	Type() string
	Inspect() string
}

// Environment stores variable bindings
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment creates a new Environment
func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
		outer: nil,
	}
}

// NewEnclosedEnvironment creates a new Environment with an outer Environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get retrieves a value from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets a value in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// Check if an object is an error
func IsError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}

// Check if an object is truthy
func IsTruthy(obj Object) bool {
	switch obj := obj.(type) {
	case *Boolean:
		return obj.Value
	case *Null:
		return false
	default:
		return true
	}
}
