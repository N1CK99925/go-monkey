package parser

import (
	"go-monkey/interpreter/ast"
	"go-monkey/interpreter/lexer"
	"go-monkey/interpreter/token"
)

// l = current instance of lexer on which we repeadly call next token
type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//read two tokens so currotken and peektoken are set

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseProgram() *ast.Program {
	return nil
}
