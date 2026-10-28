package ast

import "go-monkey/interpreter/token"

// every node in ASt has to implement Node and give a tokenLiteral that returns the value of that token
// TokenLiteral( ) is only used for testing and debugging
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifer
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifer struct {
	Token token.Token
	Value string
}

func (i *Identifer) expressionNode() {}
func (i *Identifer) TokenLiteral() string {
	return string(i.Token.Literal)
}
