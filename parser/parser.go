package parser

import (
	"fmt"

	"github.com/medragneel/lex/ast"
	"github.com/medragneel/lex/lexer"
	"github.com/medragneel/lex/token"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken token.Token
	// This represents the next token in the input stream, allowing the parser to look ahead without consuming the token.
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := Parser{l: l, errors: []string{}}

	p.nextToken()
	p.nextToken()

	return &p

}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()

}
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t

}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t

}
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}

}

// Relevant parts of the parser implementation

func (p *Parser) ParseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// NOTE: This loop skips the entire expression without parsing it
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
func (p *Parser) ParseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}

func (p *Parser) parseStatment() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.ParseLetStatement()
	case token.RETURN:
		return p.ParseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatment()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}
