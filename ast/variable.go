package ast

import (
	"ast-operation-parser/ast/expressions"
	"ast-operation-parser/ast/nodes"
	"ast-operation-parser/lexer/operators"
	"ast-operation-parser/lexer/token"
	"errors"
	"fmt"
)

func MakeVariableDefinition(tokens *[]token.Token, i *int) (*nodes.VariableDefinition, error) {
	// Check if there are enough token for a variable declaration
	if !(*i+1 < len(*tokens)) || !(*i+2 < len(*tokens)) {
		return nil, errors.New("not enough arguments to define variable")
	}

	// Get the variable token and the assignment operator token
	varT := &(*tokens)[*i+1]
	opT := &(*tokens)[*i+2]

	// Debugging output

	// Validate the variable declaration format
	if varT.Type != token.TOKEN_IDENT || opT.Type != token.TOKEN_ASSIGN {
		return nil, errors.New("invalid Variable declaration format")
	}

	// Move the index forward to skip the variable and assignment token
	*i += 3

	// expression for thing
	var valueTokens []token.Token

	// Expecting an expression after the assignment operator
	for *i < len(*tokens) {
		t := (*tokens)[*i]
		if TokenEnd(t) { // Stop if we reach the end token
			break
		}
		valueTokens = append(valueTokens, t) // Collect the token
		*i++                                 // Move to the next token
	}
	//rpn, err := expressions.InfixToRPN(valueTokens)
	//if err != nil {
	//
	//	return nil, err
	//}
	//expression, err := expressions.ParseRPN(rpn)
	fmt.Println(valueTokens)
	expression, err := expressions.InfixToRPNAST(valueTokens)
	if err != nil {
		return nil, err
	}
	for {
		if *i >= len(*tokens) {
			break
		}
		t := (*tokens)[*i]
		if TokenEnd(t) {
			break
		}
		*i++

	}

	// Create the VariableDefinition and return it
	return &nodes.VariableDefinition{
		Identifier: varT.Value, // The variable name
		Value:      expression,
	}, nil
}

func MakeVariableAssigning(tokens *[]token.Token, i *int) (*nodes.VariableAssigning, error) {
	// Check if there are enough token for a variable declaration
	if !(*i < len(*tokens)) || !(*i+1 < len(*tokens)) {
		return nil, errors.New("not enough arguments to define variable")
	}

	// Get the variable token and the assignment operator token
	varT := &(*tokens)[*i]
	opT := &(*tokens)[*i+1]

	// Debugging output

	// Validate the variable declaration format
	if _, ok := operators.AssigningOperators[opT.Type]; varT.Type != token.TOKEN_IDENT || !ok {
		return nil, errors.New("invalid Variable assigning format")
	}
	if v := operators.AssigningOperators[opT.Type]; len(v) == 0 {
		if *i+3 >= len(*tokens) {
			return nil, errors.New("invalid Variable assigning format, expected 0 arguments")
		}
		return &nodes.VariableAssigning{
			Identifier: varT.Value,
			Operator:   opT.Value,
			Value:      nil,
		}, nil
	}
	// Move the index forward to skip the variable and assignment token
	*i += 2
	if *i >= len(*tokens) {
		return nil, errors.New("invalid Variable assigning format, expected an expression after the assignment operator")
	}

	if !expressions.IsValue(tokens, i) || TokenEnd((*tokens)[*i]) {
		return nil, errors.New("invalid Variable assigning format, expected an expression after the assignment operator")
	}
	// expression for thing
	var valueTokens []token.Token

	// Expecting an expression after the assignment operator
	for *i < len(*tokens) {
		t := (*tokens)[*i]
		if TokenEnd(t) { // Stop if we reach the end token
			break
		}
		valueTokens = append(valueTokens, t) // Collect the token
		*i++                                 // Move to the next token
	}
	expression, err := expressions.InfixToRPNAST(valueTokens)

	if err != nil {
		return nil, err
	}

	for {
		if *i >= len(*tokens) {
			break
		}
		t := (*tokens)[*i]
		if TokenEnd(t) {
			break
		}
		*i++

	}

	// Create the VariableDefinition and return it
	return &nodes.VariableAssigning{
		Identifier: varT.Value, // The variable name
		Value:      expression,
		Operator:   opT.Value,
	}, nil
}
