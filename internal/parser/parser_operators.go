// file: internal/parser/parser_operators.go
// description: Operator expression parsing for the TRUMP programming language

package parser

import (
	"fmt"

	"github.com/AndrewDonelson/trumplang/internal/errors"
	"github.com/AndrewDonelson/trumplang/internal/lexer/token"
)

// Parse a prefix expression
func (p *Parser) parsePrefixExpression() Expression {
	expression := &PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

// Parse an infix expression
func (p *Parser) parseInfixExpression(left Expression) Expression {
	expression := &InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

// Parse a grouped expression
func (p *Parser) parseGroupedExpression() Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected ')'")
		return nil
	}

	return exp
}

// Parse an array literal
func (p *Parser) parseArrayLiteral() Expression {
	array := &ArrayLiteral{Token: p.curToken}

	array.Elements = p.parseExpressionList(token.RBRACKET)

	return array
}

// Parse a list of expressions
func (p *Parser) parseExpressionList(end token.TokenType) []Expression {
	list := []Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()

	// Handle comments inside expression lists
	if p.curTokenIs(token.COMMENT) {
		p.nextToken()
	}

	// Ensure we don't hit the end bracket after skipping comments
	if p.curTokenIs(end) {
		return list
	}

	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		// Handle comments after commas
		if p.curTokenIs(token.COMMENT) {
			p.nextToken()
		}

		// If we reach the end after a comment, stop
		if p.curTokenIs(end) {
			break
		}

		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		p.addError(errors.UNEXPECTED_TOKEN, fmt.Sprintf("Expected '%s'", end))
		return nil
	}

	return list
}

// Parse a function literal
func (p *Parser) parseFunctionLiteral() Expression {
	lit := &FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected '(' after FUNCTION")
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	// Parse optional rating
	if p.peekTokenIs(token.RATED) {
		p.nextToken()

		// Get rating value
		p.nextToken()
		ratingStart := p.curToken.Literal

		// Check for division format (e.g., 10/10)
		if p.peekTokenIs(token.SLASH) {
			p.nextToken() // consume slash
			p.nextToken() // get denominator
			ratingStart = ratingStart + "/" + p.curToken.Literal
		}

		lit.Rating = ratingStart
	}

	if !p.expectPeek(token.LBRACE) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected '{' after function parameters")
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

// Parse function parameters
func (p *Parser) parseFunctionParameters() []*Identifier {
	identifiers := []*Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected ')'")
		return nil
	}

	return identifiers
}

// Parse a call expression
func (p *Parser) parseCallExpression(function Expression) Expression {
	exp := &CallExpression{
		Token:    p.curToken,
		Function: function,
	}

	exp.Arguments = p.parseExpressionList(token.RPAREN)

	return exp
}

// Parse an index expression
func (p *Parser) parseIndexExpression(left Expression) Expression {
	exp := &IndexExpression{
		Token: p.curToken,
		Left:  left,
	}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected ']'")
		return nil
	}

	return exp
}
