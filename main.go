package main

import (
	"ast-operation-parser/ast"
	"ast-operation-parser/lexer"
	"fmt"
)

func main() {
	input := `var ābols = true
a += lol(10, 10+10+10 + lol(10));
`
	tokens := lexer.Tokenize(input)
	//rpnE, err := expressions.InfixToRPNAST(tokens)
	rpn, err := ast.BuildAST(tokens)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(rpn.Children)
	//fmt.Println(rpnE)

	// Build the AST
	// astRoot := ast.BuildAST(token)

	// Print the AST
	// ast.PrintAST(astRoot, 0)
}
