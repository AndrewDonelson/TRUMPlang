// file: internal/parser/ast_statements.go
// description: Statement AST nodes for the TRUMP programming language

package parser

import (
	"bytes"

	"github.com/AndrewDonelson/trumplang/internal/lexer/token"
)

// WhileStatement represents a while loop
// e.g., "MAKE DEALS WHILE (x < 10) { ... }"
type WhileStatement struct {
	Token     token.Token // the 'MAKE' token
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
	var out bytes.Buffer

	out.WriteString("MAKE DEALS WHILE ")
	out.WriteString("(")
	out.WriteString(ws.Condition.String())
	out.WriteString(") ")
	out.WriteString(ws.Body.String())

	return out.String()
}

// ForStatement represents a for loop
// e.g., "MAKE AMERICA GREAT AGAIN FOR (i=0; i<10; i++) { ... }"
type ForStatement struct {
	Token     token.Token // the 'MAKE' token
	Init      Statement
	Condition Expression
	Update    Statement
	Body      *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer

	out.WriteString("MAKE AMERICA GREAT AGAIN FOR ")
	out.WriteString("(")
	if fs.Init != nil {
		out.WriteString(fs.Init.String())
	}
	out.WriteString("; ")

	if fs.Condition != nil {
		out.WriteString(fs.Condition.String())
	}
	out.WriteString("; ")

	if fs.Update != nil {
		out.WriteString(fs.Update.String())
	}
	out.WriteString(") ")
	out.WriteString(fs.Body.String())

	return out.String()
}

// TweetStatement represents a print statement
// e.g., "TWEET x;"
type TweetStatement struct {
	Token token.Token // the 'TWEET' token
	Value Expression
}

func (ts *TweetStatement) statementNode()       {}
func (ts *TweetStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *TweetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ts.TokenLiteral() + " ")

	if ts.Value != nil {
		out.WriteString(ts.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// RallyStatement represents an emphasized print statement
// e.g., "RALLY x;"
type RallyStatement struct {
	Token token.Token // the 'RALLY' token
	Value Expression
}

func (rs *RallyStatement) statementNode()       {}
func (rs *RallyStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *RallyStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExecutiveOrderStatement represents an error/warning output statement
// e.g., "EXECUTIVE_ORDER x;"
type ExecutiveOrderStatement struct {
	Token token.Token // the 'EXECUTIVE_ORDER' token
	Value Expression
}

func (eos *ExecutiveOrderStatement) statementNode()       {}
func (eos *ExecutiveOrderStatement) TokenLiteral() string { return eos.Token.Literal }
func (eos *ExecutiveOrderStatement) String() string {
	var out bytes.Buffer

	out.WriteString(eos.TokenLiteral() + " ")

	if eos.Value != nil {
		out.WriteString(eos.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
