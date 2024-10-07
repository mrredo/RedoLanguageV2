package nodes

import (
	"ast-operation-parser/lexer/token"
	"errors"
)

type NodeType int8

const (
	VariableDeclaration NodeType = iota // 0
	Identifier                          // 1
	Assignment                          // 2
	Literal                             // 3
	BinaryOperator
	NodeFunctionCall
	Variable
	VariableAssign
)

type ASTStructure struct {
	FileName string
	Children []ASTNode
}
type ASTNode interface {
	Type() NodeType // Method to return the type of the AST node
}

// VariableDefinition represents a variable declaration/definition in the AST.
type VariableDefinition struct {
	Identifier string  // The name of the variable
	Value      ASTNode // The assigned value (could be a literal, expression, etc.)
}

// Implement the Type method for VariableDefinition
func (v VariableDefinition) Type() NodeType {
	return VariableDeclaration
}

// FunctionCall represents a function call in the AST.
type FunctionCall struct {
	FunctionName string    // The name of the function being called
	Arguments    []ASTNode // A list of arguments for the function
}

// Implement the Type method for FunctionCall
func (f FunctionCall) Type() NodeType {
	return NodeFunctionCall
}

func NewBinaryOperation(left, right ASTNode, operator string, AcceptedDataTypes []DataType, ReturnedDataTypes []DataType) (BinaryOperation, error) {
	if _, ok := token.Operators[operator]; !ok {
		return BinaryOperation{}, errors.New("invalid operator")
	}

	return BinaryOperation{
		AcceptedDataTypes: AcceptedDataTypes,
		ReturnedDataTypes: ReturnedDataTypes,
		Operator:          operator,
		Left:              left,
		Right:             right,
	}, nil
}

// BinaryOperation represents a binary operation (e.g., addition, subtraction).
type BinaryOperation struct {
	AcceptedDataTypes []DataType
	ReturnedDataTypes []DataType
	Operator          string  // The operator (e.g., "+", "-", "*", "/")
	Left              ASTNode // Left operand
	Right             ASTNode // Right operand
}

// Implement the Type method for BinaryOperation
func (b BinaryOperation) Type() NodeType {
	return BinaryOperator
}

// Literal represents a constant value like a number or a string.
type LiteralNode struct {
	DataType DataType
	Value    string // The literal value (e.g., "10", "true", "hello")
}

// Implement the Type method for Literal
func (l LiteralNode) Type() NodeType {
	return Literal
}

type VariableNode struct {
	//operators types are only checked in the runtime/execution of the program
	// DataType structs.DataType
	Value string // The literal value (e.g., "10", "true", "hello")
}

// Implement the Type method for Literal
func (l VariableNode) Type() NodeType {
	return Variable
}

type VariableAssigning struct {
	Identifier string // The name of the variable
	Operator   string
	Value      ASTNode // The assigned value (could be a literal, expression, etc.)
}

// Implement the Type method for VariableDefinition
func (v VariableAssigning) Type() NodeType {
	return VariableAssign
}
