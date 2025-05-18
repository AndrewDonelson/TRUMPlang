// file: internal/interpreter/evaluator_identifiers.go
// description: Identifier and compound expression evaluation for the TRUMP language

package interpreter

import (
	"github.com/AndrewDonelson/trumplang/internal/parser"
)

// Evaluate an identifier
func (e *Evaluator) evalIdentifier(node *parser.Identifier) Object {
	// Check for built-in functions
	if builtin, ok := e.builtins[node.Value]; ok {
		return builtin
	}

	// Check for variables in the environment
	val, ok := e.env.Get(node.Value)
	if !ok {
		// Easter egg: Undefined variables are "covfefe"
		if e.rand.Float64() < 0.1 {
			return newError("Nobody knows what this '%s' covfefe means, but it's provocative!", node.Value)
		}
		return newError("identifier not found: " + node.Value)
	}

	return val
}

// Evaluate an index expression
func (e *Evaluator) evalIndexExpression(left, index Object) Object {
	switch {
	case left.Type() == ARRAY_OBJ:
		return e.evalArrayIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

// Evaluate an array index expression
func (e *Evaluator) evalArrayIndexExpression(array, index Object) Object {
	arrayObject := array.(*Array)
	idx, ok := index.(*Integer)
	if !ok {
		return newError("array index must be INTEGER, got %s", index.Type())
	}

	if idx.Value < 0 || idx.Value >= int64(len(arrayObject.Elements)) {
		return e.NULL
	}

	return arrayObject.Elements[idx.Value]
}

// Evaluate expressions
func (e *Evaluator) evalExpressions(exps []parser.Expression) []Object {
	var result []Object

	for _, exp := range exps {
		evaluated := e.Eval(exp)
		if IsError(evaluated) {
			return []Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

// Apply a function
func (e *Evaluator) applyFunction(fn Object, args []Object) Object {
	switch fn := fn.(type) {
	case *Function:
		extendedEnv := e.extendFunctionEnv(fn, args)
		oldEnv := e.env
		e.env = extendedEnv

		evaluated := e.Eval(fn.Body)
		e.env = oldEnv

		return unwrapReturnValue(evaluated)
	case *Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
}

// Extend the environment with function parameters
func (e *Evaluator) extendFunctionEnv(fn *Function, args []Object) *Environment {
	env := NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		if paramIdx < len(args) {
			env.Set(param.Value, args[paramIdx])
		} else {
			env.Set(param.Value, e.NULL)
		}
	}

	return env
}
