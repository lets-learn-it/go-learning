// Code generated from Expr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Expr

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ExprParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by ExprParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by ExprParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by ExprParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by ExprParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by ExprParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by ExprParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by ExprParser#int.
	VisitInt(ctx *IntContext) interface{}
}
