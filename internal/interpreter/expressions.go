// file: internal/interpreter/evaluator_expressions.go
// description: Expression evaluation for the TRUMP programming language

package interpreter

// Evaluate a prefix expression
func (e *Evaluator) evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return e.evalBangOperatorExpression(right)
	case "-":
		return e.evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

// Evaluate a bang operator expression
func (e *Evaluator) evalBangOperatorExpression(right Object) Object {
	switch right {
	case e.TRUE:
		return e.FALSE
	case e.FALSE:
		return e.TRUE
	case e.NULL:
		return e.TRUE
	default:
		return e.FALSE
	}
}

// Evaluate a minus prefix operator expression
func (e *Evaluator) evalMinusPrefixOperatorExpression(right Object) Object {
	switch right.Type() {
	case INTEGER_OBJ:
		value := right.(*Integer).Value
		return &Integer{Value: -value}
	case FLOAT_OBJ:
		value := right.(*Float).Value
		return &Float{Value: -value}
	default:
		return newError("unknown operator: -%s", right.Type())
	}
}

// Evaluate an infix expression
func (e *Evaluator) evalInfixExpression(operator string, left, right Object) Object {
	switch {
	case left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ:
		return e.evalIntegerInfixExpression(operator, left, right)
	case left.Type() == FLOAT_OBJ && right.Type() == FLOAT_OBJ:
		return e.evalFloatInfixExpression(operator, left, right)
	case left.Type() == INTEGER_OBJ && right.Type() == FLOAT_OBJ:
		intValue := float64(left.(*Integer).Value)
		return e.evalFloatInfixExpression(operator, &Float{Value: intValue}, right)
	case left.Type() == FLOAT_OBJ && right.Type() == INTEGER_OBJ:
		intValue := float64(right.(*Integer).Value)
		return e.evalFloatInfixExpression(operator, left, &Float{Value: intValue})
	case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ:
		return e.evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return e.nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return e.nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// Evaluate an integer infix expression
func (e *Evaluator) evalIntegerInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Integer).Value
	rightVal := right.(*Integer).Value

	switch operator {
	case "+":
		return &Integer{Value: leftVal + rightVal}
	case "-":
		return &Integer{Value: leftVal - rightVal}
	case "*":
		return &Integer{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("division by zero")
		}
		return &Integer{Value: leftVal / rightVal}
	case "<":
		return e.nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return e.nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return e.nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return e.nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return e.nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return e.nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// Evaluate a float infix expression
func (e *Evaluator) evalFloatInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Float).Value
	rightVal := right.(*Float).Value

	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("division by zero")
		}
		return &Float{Value: leftVal / rightVal}
	case "<":
		return e.nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return e.nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return e.nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return e.nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return e.nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return e.nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// Evaluate a string infix expression
func (e *Evaluator) evalStringInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*String).Value
	rightVal := right.(*String).Value

	switch operator {
	case "+":
		return &String{Value: leftVal + rightVal}
	case "==":
		return e.nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return e.nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
