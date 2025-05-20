// file: internal/interpreter/evaluator_statements.go
// description: Statement evaluation for the TRUMP programming language

package interpreter

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/AndrewDonelson/trumplang/internal/parser"
)

// Evaluate a program
func (e *Evaluator) evalProgram(program *parser.Program) Object {
	var result Object = e.NULL

	for _, statement := range program.Statements {
		result = e.Eval(statement)

		switch result := result.(type) {
		case *ReturnValue:
			return result.Value
		case *Error:
			return result
		}
	}

	return result
}

// Evaluate a block statement
func (e *Evaluator) evalBlockStatement(block *parser.BlockStatement) Object {
	var result Object = e.NULL

	for _, statement := range block.Statements {
		result = e.Eval(statement)

		if result != nil {
			rt := result.Type()
			if rt == RETURN_OBJ || rt == ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

// Evaluate an if statement
func (e *Evaluator) evalIfStatement(is *parser.IfStatement) Object {
	condition := e.Eval(is.Condition)
	if IsError(condition) {
		return condition
	}

	// Add Trump unpredictability: 5% chance to randomly flip the condition
	if e.rand.Float64() < 0.05 {
		if IsTruthy(condition) {
			condition = e.FALSE
		} else {
			condition = e.TRUE
		}
	}

	if IsTruthy(condition) {
		return e.Eval(is.Consequence)
	} else if is.Alternative != nil {
		return e.Eval(is.Alternative)
	} else {
		return e.NULL
	}
}

// Evaluate a while statement
func (e *Evaluator) evalWhileStatement(ws *parser.WhileStatement) Object {
	var result Object = e.NULL
	maxIterations := 10000 // Prevent infinite loops

	condition := e.Eval(ws.Condition)
	if IsError(condition) {
		return condition
	}

	iterations := 0
	for IsTruthy(condition) && iterations < maxIterations {
		iterations++

		result = e.Eval(ws.Body)

		if result != nil {
			rt := result.Type()
			if rt == RETURN_OBJ || rt == ERROR_OBJ {
				return result
			}
		}

		condition = e.Eval(ws.Condition)
		if IsError(condition) {
			return condition
		}
	}

	if iterations >= maxIterations {
		return newError("While loop exceeded maximum iterations (possible infinite loop)")
	}

	return result
}

// Evaluate a for statement
func (e *Evaluator) evalForStatement(fs *parser.ForStatement) Object {
	// Create a new environment for the for loop
	outerEnv := e.env
	e.env = NewEnclosedEnvironment(outerEnv)
	maxIterations := 10000 // Prevent infinite loops

	// Initialize
	initResult := e.Eval(fs.Init)
	if IsError(initResult) {
		e.env = outerEnv // Restore environment
		return initResult
	}

	var result Object = e.NULL
	iterations := 0

	// Check condition
	condition := e.Eval(fs.Condition)
	if IsError(condition) {
		e.env = outerEnv // Restore environment
		return condition
	}

	for IsTruthy(condition) && iterations < maxIterations {
		iterations++

		// Execute body
		result = e.Eval(fs.Body)

		if result != nil {
			rt := result.Type()
			if rt == RETURN_OBJ || rt == ERROR_OBJ {
				e.env = outerEnv // Restore environment
				return result
			}
		}

		// Update
		updateResult := e.Eval(fs.Update)
		if IsError(updateResult) {
			e.env = outerEnv // Restore environment
			return updateResult
		}

		// Check condition again
		condition = e.Eval(fs.Condition)
		if IsError(condition) {
			e.env = outerEnv // Restore environment
			return condition
		}
	}

	if iterations >= maxIterations {
		e.env = outerEnv // Restore environment
		return newError("For loop exceeded maximum iterations (possible infinite loop)")
	}

	e.env = outerEnv // Restore environment
	return result
}

// Evaluate a tweet statement (print)
func (e *Evaluator) evalTweetStatement(ts *parser.TweetStatement) Object {
	val := e.Eval(ts.Value)
	if IsError(val) {
		return val
	}

	fmt.Println("🐦 " + val.Inspect())
	return e.NULL
}

// Evaluate a rally statement (emphasized print)
func (e *Evaluator) evalRallyStatement(rs *parser.RallyStatement) Object {
	val := e.Eval(rs.Value)
	if IsError(val) {
		return val
	}

	// Add random Trump-like emphasis
	result := strings.ToUpper(val.Inspect())
	result = addTrumpEmphasis(result)

	fmt.Println("🔊 " + result + " 👐")
	return e.NULL
}

// Add Trump-like emphasis to a string
func addTrumpEmphasis(s string) string {
	emphases := []string{
		", BELIEVE ME!",
		" - TREMENDOUS!",
		", IT'S TRUE!",
		" - THE BEST!",
		", FOLKS!",
		" - BIGLY!",
		", OKAY?",
		" - SO TRUE!",
		", THAT I CAN TELL YOU!",
		" - EVERYBODY KNOWS IT!",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	emphasis := emphases[r.Intn(len(emphases))]

	return s + emphasis
}

// Evaluate an executive order statement (error/warning print)
func (e *Evaluator) evalExecutiveOrderStatement(eos *parser.ExecutiveOrderStatement) Object {
	val := e.Eval(eos.Value)
	if IsError(val) {
		return val
	}

	fmt.Fprintln(os.Stderr, "⚠️ EXECUTIVE ORDER: "+val.Inspect()+" ⚠️")
	return e.NULL
}

// Convert a native boolean to a Boolean object
func (e *Evaluator) nativeBoolToBooleanObject(input bool) *Boolean {
	if input {
		return e.TRUE
	}
	return e.FALSE
}
