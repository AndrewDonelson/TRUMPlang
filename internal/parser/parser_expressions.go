// file: internal/parser/parser_expressions.go
// description: Expression parsing for the TRUMP programming language

package parser

import (
	"fmt"
	"strconv"

	"github.com/AndrewDonelson/trumplang/internal/errors"
	"github.com/AndrewDonelson/trumplang/internal/lexer/token"
)

// Parse a for statement
func (p *Parser) parseForStatement() *ForStatement {
	stmt := &ForStatement{Token: p.curToken}

	// Expect MAKE AMERICA GREAT AGAIN FOR
	if !p.expectPeek(token.AMERICA) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected AMERICA after MAKE")
		return nil
	}

	if !p.expectPeek(token.GREAT) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected GREAT after MAKE AMERICA")
		return nil
	}

	if !p.expectPeek(token.AGAIN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected AGAIN after MAKE AMERICA GREAT")
		return nil
	}

	if !p.expectPeek(token.FOR) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected FOR after MAKE AMERICA GREAT AGAIN")
		return nil
	}

	// Parse for loop components
	if !p.expectPeek(token.LPAREN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected '(' after FOR")
		return nil
	}

	p.nextToken()

	// Parse initialization
	stmt.Init = p.parseStatement()

	if !p.expectPeek(token.SEMICOLON) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected ';' after initialization")
		return nil
	}

	p.nextToken()

	// Parse condition
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.SEMICOLON) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected ';' after condition")
		return nil
	}

	p.nextToken()

	// Parse update
	stmt.Update = p.parseStatement()

	if !p.expectPeek(token.RPAREN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected ')' after update")
		return nil
	}

	// Parse body
	if !p.expectPeek(token.LBRACE) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected '{' after for loop declaration")
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

// Parse a block of statements
func (p *Parser) parseBlockStatement() *BlockStatement {
	block := &BlockStatement{
		Token:      p.curToken,
		Statements: []Statement{},
	}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

// Parse an expression statement
func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Parse an expression
func (p *Parser) parseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

// Parse an identifier
func (p *Parser) parseIdentifier() Expression {
	return &Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

// Parse an integer literal
func (p *Parser) parseIntegerLiteral() Expression {
	lit := &IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.addError(errors.UNEXPECTED_TOKEN, msg)
		return nil
	}

	lit.Value = value

	return lit
}

// Parse a float literal
func (p *Parser) parseFloatLiteral() Expression {
	lit := &FloatLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Literal)
		p.addError(errors.UNEXPECTED_TOKEN, msg)
		return nil
	}

	lit.Value = value

	return lit
}

// Parse a string literal
func (p *Parser) parseStringLiteral() Expression {
	return &StringLiteral{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

// Parse a boolean literal
func (p *Parser) parseBoolean() Expression {
	return &BooleanLiteral{
		Token: p.curToken,
		Value: p.curTokenIs(token.WINNING),
	}
}