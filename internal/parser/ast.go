// file: internal/parser/ast.go
// description: Abstract Syntax Tree (AST) for the TRUMP programming language

package parser

import (
	"bytes"

	"github.com/AndrewDonelson/trumplang/internal/lexer/token"
)

// Node is the interface that all AST nodes implement
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is a Node that represents a statement
type Statement interface {
	Node
	statementNode()
}

// Expression is a Node that represents an expression
type Expression interface {
	Node
	expressionNode()
}

// Program represents the entire program
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement represents a variable declaration statement
// e.g., "YUGE x = 5;" or "TREMENDOUS y = 10;"
type LetStatement struct {
	Token token.Token // YUGE or TREMENDOUS
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement represents a return statement
// e.g., "RETURN x;"
type ReturnStatement struct {
	Token       token.Token // the 'RETURN' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement represents a statement consisting solely of an expression
// e.g., "x + y;"
type ExpressionStatement struct {
	Token      token.Token // The first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// BlockStatement represents a block of statements enclosed in braces
// e.g., "{ x = 5; y = 10; }"
type BlockStatement struct {
	Token      token.Token // the '{' token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString("{ ")

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	out.WriteString(" }")

	return out.String()
}

// IfStatement represents an if conditional
// e.g., "BUILD WALL IF (x > 5) { ... } ELSE { ... }"
type IfStatement struct {
	Token       token.Token // the 'BUILD' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {
	var out bytes.Buffer

	out.WriteString("BUILD WALL IF ")
	out.WriteString("(")
	out.WriteString(is.Condition.String())
	out.WriteString(") ")
	out.WriteString(is.Consequence.String())

	if is.Alternative != nil {
		out.WriteString(" ELSE ")
		out.WriteString(is.Alternative.String())
	}

	return out.String()
}
