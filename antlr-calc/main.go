package main

import (
	"antlr-calc/parser"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type exprListener struct {
	*parser.BaseExprListener
	stack []int
}

func (l *exprListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *exprListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty ubable to pop")
	}
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

func (l *exprListener) ExitMulDiv(ctx *parser.MulDivContext) {
	right, left := l.pop(), l.pop()
	switch ctx.GetOp().GetTokenType() {
	case parser.ExprParserMUL:
		l.push(left * right)
	case parser.ExprParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (l *exprListener) ExitAddSub(ctx *parser.AddSubContext) {
	right, left := l.pop(), l.pop()
	switch ctx.GetOp().GetTokenType() {
	case parser.ExprParserADD:
		l.push(left + right)
	case parser.ExprParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", ctx.GetOp().GetText()))
	}
}

func (l *exprListener) ExitNumber(ctx *parser.NumberContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err.Error())
	}
	l.push(i)
}
func calc(input string) int {
	is := antlr.NewInputStream(input)
	lexer := parser.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewExprParser(stream) // 事件驱动 API

	listener := &exprListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Expr())
	return listener.pop()
}

func main() {
	// lexer
	res := calc("142857*2+1")
	fmt.Printf("%d\n", res)
}
