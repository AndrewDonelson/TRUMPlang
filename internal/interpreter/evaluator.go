// file: internal/interpreter/evaluator.go
// description: Main evaluator for the TRUMP programming language

package interpreter

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AndrewDonelson/trumplang/internal/parser"
)

// Create a new error
func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

// Unwrap a return value
func unwrapReturnValue(obj Object) Object {
	if returnValue, ok := obj.(*ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

// Evaluator evaluates a TRUMP program
type Evaluator struct {
	env *Environment
	// Cache the TRUE and FALSE values for efficiency
	TRUE  *Boolean
	FALSE *Boolean
	NULL  *Null

	// Random number generator for Trump-like behavior
	rand *rand.Rand

	// Built-in functions
	builtins map[string]Object
}

// NewEvaluator creates a new Evaluator
func NewEvaluator() *Evaluator {
	e := &Evaluator{
		env:      NewEnvironment(),
		TRUE:     &Boolean{Value: true},
		FALSE:    &Boolean{Value: false},
		NULL:     &Null{},
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
		builtins: make(map[string]Object),
	}

	// Register built-in functions
	e.registerBuiltins()

	return e
}

// Eval evaluates a node
func (e *Evaluator) Eval(node parser.Node) Object {
	switch node := node.(type) {
	// Statements
	case *parser.Program:
		return e.evalProgram(node)
	case *parser.ExpressionStatement:
		return e.Eval(node.Expression)
	case *parser.BlockStatement:
		return e.evalBlockStatement(node)
	case *parser.LetStatement:
		val := e.Eval(node.Value)
		if IsError(val) {
			return val
		}
		e.env.Set(node.Name.Value, val)
	case *parser.ReturnStatement:
		val := e.Eval(node.ReturnValue)
		if IsError(val) {
			return val
		}
		return &ReturnValue{Value: val}
	case *parser.IfStatement:
		return e.evalIfStatement(node)
	case *parser.WhileStatement:
		return e.evalWhileStatement(node)
	case *parser.ForStatement:
		return e.evalForStatement(node)
	case *parser.TweetStatement:
		return e.evalTweetStatement(node)
	case *parser.RallyStatement:
		return e.evalRallyStatement(node)
	case *parser.ExecutiveOrderStatement:
		return e.evalExecutiveOrderStatement(node)

	// Expressions
	case *parser.IntegerLiteral:
		return &Integer{Value: node.Value}
	case *parser.FloatLiteral:
		return &Float{Value: node.Value}
	case *parser.StringLiteral:
		return &String{Value: node.Value}
	case *parser.BooleanLiteral:
		return e.nativeBoolToBooleanObject(node.Value)
	case *parser.PrefixExpression:
		right := e.Eval(node.Right)
		if IsError(right) {
			return right
		}
		return e.evalPrefixExpression(node.Operator, right)
	case *parser.InfixExpression:
		left := e.Eval(node.Left)
		if IsError(left) {
			return left
		}

		right := e.Eval(node.Right)
		if IsError(right) {
			return right
		}

		return e.evalInfixExpression(node.Operator, left, right)
	case *parser.Identifier:
		return e.evalIdentifier(node)
	case *parser.ArrayLiteral:
		elements := e.evalExpressions(node.Elements)
		if len(elements) == 1 && IsError(elements[0]) {
			return elements[0]
		}
		return &Array{Elements: elements}
	case *parser.IndexExpression:
		left := e.Eval(node.Left)
		if IsError(left) {
			return left
		}
		index := e.Eval(node.Index)
		if IsError(index) {
			return index
		}
		return e.evalIndexExpression(left, index)
	case *parser.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		rating := node.Rating
		return &Function{Parameters: params, Body: body, Env: e.env, Rating: rating}
	case *parser.CallExpression:
		function := e.Eval(node.Function)
		if IsError(function) {
			return function
		}

		args := e.evalExpressions(node.Arguments)
		if len(args) == 1 && IsError(args[0]) {
			return args[0]
		}

		return e.applyFunction(function, args)
	}

	return nil
}
