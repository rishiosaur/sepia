package parser

import (
	"fmt"
	"sepia/ast"
	"sepia/lexer"
	"sepia/token"
	"strconv"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token

	errors []string

	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l,
		errors: []string{}}

	p.consumeToken()
	p.consumeToken()

	p.prefixParseFns = make(map[token.Type]prefixParseFn)
	p.registerPrefixFunction(token.IDENT, p.parseIdentifier)
	p.registerPrefixFunction(token.INT, p.parseIntegerLiteral)
	p.registerPrefixFunction(token.BANG, p.parsePrefixExpression)
	p.registerPrefixFunction(token.MINUS, p.parsePrefixExpression)
	p.registerPrefixFunction(token.TRUE, p.parseBoolean)
	p.registerPrefixFunction(token.FALSE, p.parseBoolean)
	p.registerPrefixFunction(token.FALSE, p.parseBoolean)
	p.registerPrefixFunction(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefixFunction(token.IF, p.parseIfExpression)
	p.registerPrefixFunction(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefixFunction(token.STRING, p.parseString)

	p.infixParseFns = make(map[token.Type]infixParseFn)
	p.registerInfixFunction(token.PLUS, p.parseInfixExpression)
	p.registerInfixFunction(token.MINUS, p.parseInfixExpression)
	p.registerInfixFunction(token.SLASH, p.parseInfixExpression)
	p.registerInfixFunction(token.ASTERISK, p.parseInfixExpression)
	p.registerInfixFunction(token.EQ, p.parseInfixExpression)
	p.registerInfixFunction(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfixFunction(token.LT, p.parseInfixExpression)
	p.registerInfixFunction(token.GT, p.parseInfixExpression)
	p.registerInfixFunction(token.LPAREN, p.parseCallExpression)

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) consumeToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currentTokenIs(tok token.Type) bool {
	return p.currentToken.Type == tok
}

func (p *Parser) peekTokenIs(tok token.Type) bool {
	return p.peekToken.Type == tok
}

func (p *Parser) addPeekError(tok token.Type) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead", tok, p.peekToken)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(tok token.Type) bool {
	if p.peekTokenIs(tok) {
		p.consumeToken()
		return true
	}

	p.addPeekError(tok)
	return false

}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.BooleanLiteral{Token: p.currentToken, Value: p.currentTokenIs(token.TRUE)}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.consumeToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	defer untrace(trace("parseInfixExpression"))
	expression := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.consumeToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) registerPrefixFunction(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfixFunction(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		_statement := p.parseStatement()

		if _statement != nil {
			program.Statements = append(program.Statements, _statement)
		}

		p.consumeToken()
	}

	return program

}

func (p *Parser) parseStatement() ast.Statement {
	defer untrace(trace("parseStatement"))
	switch p.currentToken.Type {
	case token.VALUE:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	defer untrace(trace("parseExpressionStatement"))
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.consumeToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	defer untrace(trace("parseExpression"))
	prefix := p.prefixParseFns[p.currentToken.Type]

	if prefix == nil {

		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.consumeToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	defer untrace(trace("parseIdentifier"))
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}
func (p *Parser) parsePrefixExpression() ast.Expression {
	defer untrace(trace("parsePrefixExpression"))
	expression := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal}
	p.consumeToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	defer untrace(trace("parseIntegerLiteral"))
	literal := &ast.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil

	}

	literal.Value = value
	return literal
}

func (p *Parser) noPrefixParseFnError(t token.Type) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	defer untrace(trace("parseLetStatement"))
	stmt := &ast.LetStatement{Token: p.currentToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	p.consumeToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.consumeToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	defer untrace(trace("parseReturnStatement"))
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.consumeToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	for !p.currentTokenIs(token.SEMICOLON) {
		p.consumeToken()
	}

	return stmt
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.currentToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.consumeToken()

	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.OPENBLOCK) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.ELSE) {
		p.consumeToken()

		if !p.expectPeek(token.OPENBLOCK) {
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.currentToken}
	block.Statements = []ast.Statement{}

	p.consumeToken()

	for !p.currentTokenIs(token.CLOSEBLOCK) && !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.consumeToken()
	}

	return block
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	fnLit := &ast.FunctionLiteral{Token: p.currentToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	fnLit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.OPENBLOCK) {
		return nil
	}

	fnLit.Body = p.parseBlockStatement()

	return fnLit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.consumeToken()
		return identifiers
	}

	p.consumeToken()

	ident := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.consumeToken()
		p.consumeToken()
		ident := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
		identifiers = append(identifiers, ident)

	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseCallExpression(functionIdentifier ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.currentToken, Function: functionIdentifier}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RPAREN) {
		p.consumeToken()
		return args
	}

	p.consumeToken()

	args = append(args, p.parseExpression(LOWEST))
	for p.peekTokenIs(token.COMMA) {
		p.consumeToken()
		p.consumeToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}

func (p *Parser) parseString() ast.Expression {
	return &ast.StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)
