package parser

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pskp-95/calc/lexer/lexer"
	"github.com/pskp-95/calc/lexer/token"
)

type Parser struct {
	l    *lexer.Lexer
	curr token.Token
}

func New(l *lexer.Lexer) Parser {
	return Parser{l: l, curr: l.NextToken()}
}

func (p *Parser) Consume() {
	p.curr = p.l.NextToken()
}

func (p *Parser) Run() int {
	return p.Exp()
}

func (p *Parser) Exp() int {
	a := p.Term()

	var op func(int) int
	for p.curr.Type != token.EOF {
		if p.curr.Type == token.PLUS {
			op = add(a)
		} else if p.curr.Type == token.MINUS {
			op = sub(a)
		} else {
			return a
		}

		p.Consume()
		b := p.Term()
		a = op(b)
	}

	return a
}

func (p *Parser) Term() int {
	a := p.Factor()

	var op func(int) int
	for p.curr.Type != token.EOF {
		if p.curr.Type == token.MUL {
			op = mul(a)
		} else if p.curr.Type == token.DIV {
			op = div(a)
		} else {
			return a
		}

		p.Consume()
		b := p.Factor()
		a = op(b)
	}
	return a
}

func (p *Parser) Factor() int {
	if p.curr.Type == token.LPAREN {
		p.Consume()
		a := p.Exp()

		if p.curr.Type != token.RPAREN {
			fmt.Printf("TokenType mismatch, Expected=RPAREN, Got=%q\n", p.curr.Type)
			os.Exit(1)
		}

		p.Consume()
		return a
	} else {
		return p.Integer(p.curr)
	}
}

func (p *Parser) Integer(t token.Token) int {
	if t.Type != token.INT {
		fmt.Printf("TokenType mismatch, Expected=INT, Got=%q\n", t.Type)
		os.Exit(1)
	}

	val, err := strconv.Atoi(t.Literal)
	if err != nil {
		fmt.Printf("Could not parse string to int. value=%q\n", t.Literal)
		os.Exit(1)
	}
	p.Consume()
	return val
}
