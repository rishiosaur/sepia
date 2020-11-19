package parser

import (
	"fmt"
	"monkey-go/ast"
	"monkey-go/lexer"
	"monkey-go/token"
)

type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	readToken    token.Token

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

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) consumeToken() {
	p.currentToken = p.readToken
	p.readToken = p.lexer.NextToken()
}

func (p *Parser) currentTokenIs(tok token.Type) bool {
	return p.currentToken.Type == tok
}

func (p *Parser) peekTokenIs(tok token.Type) bool {
	return p.readToken.Type == tok
}

func (p *Parser) addPeekError(tok token.Type) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead", tok, p.readToken)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(tok token.Type) bool {
	if p.peekTokenIs(tok) {
		p.consumeToken()
		return true
	} else {
		p.addPeekError(tok)
		return false
	}
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
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// PRECEDENCES
const (
	_ int = iota
	LOWEST
	EQUALS  // == LESSGREATER // > or <
	SUM     //+
	PRODUCT //*
	PREFIX  //-Xor!X
	CALL    // myFunction(X)
)

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.consumeToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]

	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: parse expressions
	for !p.currentTokenIs(token.SEMICOLON) {
		p.consumeToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.consumeToken()

	for !p.currentTokenIs(token.SEMICOLON) {
		p.consumeToken()
	}

	return stmt
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)
