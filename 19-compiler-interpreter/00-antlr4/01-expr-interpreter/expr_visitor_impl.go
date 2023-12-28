package main

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pskp-95/antlr-expr/parser"
)

type exprVisitor struct {
	parser.BaseExprVisitor
	memory map[string]int
}

func NewExprVisitor(memory map[string]int) exprVisitor {
	return exprVisitor{
		memory: memory,
	}
}

func (v *exprVisitor) Visit(tree antlr.ParseTree) int {
	switch val := tree.(type) {
	case *parser.MulDivContext:
		return v.VisitMulDiv(val)
	case *parser.AssignContext:
		return v.VisitAssign(val)
	case *parser.AddSubContext:
		return v.VisitAddSub(val)
	case *parser.PrintExprContext:
		return v.VisitPrintExpr(val)
	case *parser.IntContext:
		return v.VisitInt(val)
	case *parser.IdContext:
		return v.VisitId(val)
	case *parser.ParensContext:
		return v.VisitParens(val)
	case *parser.ProgContext:
		return v.VisitProg(val)
	default:
		panic("unknown context")
	}
}

func (e *exprVisitor) VisitProg(ctx *parser.ProgContext) int {
	for _, v := range ctx.AllStat() {
		e.Visit(v)
	}
	return 0
}

func (e *exprVisitor) VisitAssign(ctx *parser.AssignContext) int {
	id := ctx.ID().GetText()
	value := e.Visit(ctx.Expr())
	e.memory[id] = value
	return value
}

func (e *exprVisitor) VisitMulDiv(ctx *parser.MulDivContext) int {
	left := e.Visit(ctx.Expr(0))
	right := e.Visit(ctx.Expr(1))

	if ctx.GetOp().GetTokenType() == parser.ExprLexerMUL {
		return left * right
	}

	return left / right
}

func (e *exprVisitor) VisitAddSub(ctx *parser.AddSubContext) int {
	left := e.Visit(ctx.Expr(0))
	right := e.Visit(ctx.Expr(1))

	if ctx.GetOp().GetTokenType() == parser.ExprLexerADD {
		return left + right
	}

	return left - right
}

func (e *exprVisitor) VisitPrintExpr(ctx *parser.PrintExprContext) int {
	value := e.Visit(ctx.Expr())
	fmt.Println(value)
	return 0
}

func (e *exprVisitor) VisitInt(ctx *parser.IntContext) int {
	value, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		panic(err.Error())
	}

	return value
}

func (e *exprVisitor) VisitId(ctx *parser.IdContext) int {
	id := ctx.ID().GetText()
	value, exist := e.memory[id]
	if exist {
		return value
	}
	return 0
}

func (e *exprVisitor) VisitParens(ctx *parser.ParensContext) int {
	return e.Visit(ctx.Expr())
}
