package ast

import "github.com/dmarticus/monkey/token"

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
	Token token.Token // the token.LET token Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {
	// TODO
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	// the token.IDENT token Value string
	Token token.Token
}

func (i *Identifier) expressionNode() {
	// TODO
}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
