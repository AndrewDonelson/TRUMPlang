// file: internal/interpreter/builtins.go
// description: Built-in functions for the TRUMP programming language

package interpreter

// Register built-in functions
func (e *Evaluator) registerBuiltins() {
	// Standard library functions
	e.builtins["len"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments for len. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	}

	e.builtins["first"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments for first. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return e.NULL
		},
	}

	e.builtins["last"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments for last. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return e.NULL
		},
	}

	e.builtins["rest"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments for rest. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]Object, length-1)
				copy(newElements, arr.Elements[1:])
				return &Array{Elements: newElements}
			}

			return e.NULL
		},
	}

	e.builtins["push"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments for push. got=%d, want=2", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*Array)
			length := len(arr.Elements)

			newElements := make([]Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &Array{Elements: newElements}
		},
	}

	// Trump-specific built-ins
	e.builtins["DEAL"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError("wrong number of arguments for DEAL. got=%d, want=2", len(args))
			}

			// Return an array with swapped values
			return &Array{Elements: []Object{args[1], args[0]}}
		},
	}

	e.builtins["BUILD"] = &Builtin{
		Fn: func(args ...Object) Object {
			// Constructor function - initializes new objects
			// For now, it just returns the first argument or NULL if none provided
			if len(args) < 1 {
				return e.NULL
			}
			return args[0]
		},
	}

	e.builtins["FIRE"] = &Builtin{
		Fn: func(args ...Object) Object {
			// Destructor function - in a real implementation, this might handle cleanup
			// For now, it just returns NULL
			return e.NULL
		},
	}

	e.builtins["TREMENDOUS_SORT"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments for TREMENDOUS_SORT. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to TREMENDOUS_SORT must be ARRAY, got %s", args[0].Type())
			}

			// This is just a simple bubble sort, but it's "the best sorting algorithm ever"
			array := args[0].(*Array)
			elements := array.Elements
			length := len(elements)

			// Make a copy to avoid modifying the original
			newElements := make([]Object, length)
			copy(newElements, elements)

			// Use a simple bubble sort (not efficient, but Trump wouldn't care about that)
			for i := 0; i < length; i++ {
				for j := 0; j < length-i-1; j++ {
					// Compare integers
					if elements[j].Type() == INTEGER_OBJ && elements[j+1].Type() == INTEGER_OBJ {
						if elements[j].(*Integer).Value > elements[j+1].(*Integer).Value {
							newElements[j], newElements[j+1] = newElements[j+1], newElements[j]
						}
					}
					// Compare floats
					if elements[j].Type() == FLOAT_OBJ && elements[j+1].Type() == FLOAT_OBJ {
						if elements[j].(*Float).Value > elements[j+1].(*Float).Value {
							newElements[j], newElements[j+1] = newElements[j+1], newElements[j]
						}
					}
					// Compare strings
					if elements[j].Type() == STRING_OBJ && elements[j+1].Type() == STRING_OBJ {
						if elements[j].(*String).Value > elements[j+1].(*String).Value {
							newElements[j], newElements[j+1] = newElements[j+1], newElements[j]
						}
					}
				}
			}

			// Add a 10% chance to randomly swap two elements, because Trump is unpredictable
			if e.rand.Float64() < 0.1 && length > 1 {
				i := e.rand.Intn(length)
				j := e.rand.Intn(length)
				newElements[i], newElements[j] = newElements[j], newElements[i]
			}

			return &Array{Elements: newElements}
		},
	}

	e.builtins["AMERICA_FIRST"] = &Builtin{
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError("wrong number of arguments for AMERICA_FIRST. got=%d, want=1", len(args))
			}
			if args[0].Type() != ARRAY_OBJ {
				return newError("argument to AMERICA_FIRST must be ARRAY, got %s", args[0].Type())
			}

			// This prioritizes certain elements in an array
			array := args[0].(*Array)
			elements := array.Elements
			length := len(elements)

			if length == 0 {
				return array
			}

			// Make a copy to avoid modifying the original
			newElements := make([]Object, length)
			copy(newElements, elements)

			// Priority value: 45 is highest priority
			for i := 0; i < length; i++ {
				if elements[i].Type() == INTEGER_OBJ && elements[i].(*Integer).Value == 45 {
					// Move to the front
					newElements[0], newElements[i] = newElements[i], newElements[0]
					break
				}
			}

			return &Array{Elements: newElements}
		},
	}
}
