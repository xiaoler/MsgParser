package main

import (
	"fmt"
	"os"

	parser "msgparser/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseMsgParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewMsgLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMsgParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Msg()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
