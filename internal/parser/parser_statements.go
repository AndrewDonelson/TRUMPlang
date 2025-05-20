// file: internal/parser/parser_statements.go
// description: Statement parsing for the TRUMP programming language

package parser

import (
	"github.com/AndrewDonelson/trumplang/internal/errors"
	"github.com/AndrewDonelson/trumplang/internal/lexer/token"
)

// Parse a statement
func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case token.YUGE, token.TREMENDOUS:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.BUILD:
		return p.parseIfStatement()
	case token.MAKE:
		if p.peekTokenIs(token.DEALS) {
			return p.parseWhileStatement()
		} else if p.peekTokenIs(token.AMERICA) {
			return p.parseForStatement()
		}
		fallthrough // If doesn't match special cases, treat as expression statement
	case token.TWEET:
		return p.parseTweetStatement()
	case token.RALLY:
		return p.parseRallyStatement()
	case token.EXECUTIVE_ORDER:
		return p.parseExecutiveOrderStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// Parse a variable declaration statement
func (p *Parser) parseLetStatement() *LetStatement {
	stmt := &LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		p.addError(errors.EXPECTED_IDENTIFIER, "Expected identifier after YUGE/TREMENDOUS")
		return nil
	}

	stmt.Name = &Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected '=' after identifier")
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	// Allow optional semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Parse a return statement
func (p *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{Token: p.curToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	// Allow optional semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Parse a tweet statement (print)
func (p *Parser) parseTweetStatement() *TweetStatement {
	stmt := &TweetStatement{Token: p.curToken}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	// Allow optional semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Parse a rally statement (emphasized print)
func (p *Parser) parseRallyStatement() *RallyStatement {
	stmt := &RallyStatement{Token: p.curToken}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	// Allow optional semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Parse an executive order statement (error/warning print)
func (p *Parser) parseExecutiveOrderStatement() *ExecutiveOrderStatement {
	stmt := &ExecutiveOrderStatement{Token: p.curToken}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	// Allow optional semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Parse an if statement
func (p *Parser) parseIfStatement() *IfStatement {
	stmt := &IfStatement{Token: p.curToken}

	// Expect BUILD WALL IF
	if !p.expectPeek(token.WALL) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected WALL after BUILD")
		return nil
	}

	if !p.expectPeek(token.IF) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected IF after BUILD WALL")
		return nil
	}

	// Parse condition
	if !p.expectPeek(token.LPAREN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected '(' after IF")
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected ')' after condition")
		return nil
	}

	// Parse consequence
	if !p.expectPeek(token.LBRACE) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected '{' after condition")
		return nil
	}

	stmt.Consequence = p.parseBlockStatement()

	// Parse optional else
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			p.addError(errors.UNEXPECTED_TOKEN, "Expected '{' after ELSE")
			return nil
		}

		stmt.Alternative = p.parseBlockStatement()
	}

	return stmt
}

// Parse a while statement
func (p *Parser) parseWhileStatement() *WhileStatement {
	stmt := &WhileStatement{Token: p.curToken}

	// Expect MAKE DEALS WHILE
	if !p.expectPeek(token.DEALS) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected DEALS after MAKE")
		return nil
	}

	if !p.expectPeek(token.WHILE) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected WHILE after MAKE DEALS")
		return nil
	}

	// Parse condition
	if !p.expectPeek(token.LPAREN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected '(' after WHILE")
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected ')' after condition")
		return nil
	}

	// Parse body
	if !p.expectPeek(token.LBRACE) {
		p.addError(errors.UNEXPECTED_TOKEN, "Expected '{' after condition")
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}
