package operators

import (
	"ast-operation-parser/ast/nodes"
	"ast-operation-parser/lexer/token"
)

var OperatorAcceptedTypes = map[string][]nodes.DataType{
	token.TOKEN_PLUS:          {nodes.DataNumber, nodes.DataString},
	token.TOKEN_MINUS:         {nodes.DataNumber},
	token.TOKEN_MUL:           {nodes.DataNumber},
	token.TOKEN_DIV:           {nodes.DataNumber},
	token.TOKEN_MODULO:        {nodes.DataNumber},
	token.TOKEN_EQ:            {nodes.DataNumber, nodes.DataString, nodes.DataBool},
	token.TOKEN_NEQ:           {nodes.DataNumber, nodes.DataString, nodes.DataBool},
	token.TOKEN_LT:            {nodes.DataNumber},
	token.TOKEN_LTE:           {nodes.DataNumber},
	token.TOKEN_GT:            {nodes.DataNumber},
	token.TOKEN_GTE:           {nodes.DataNumber},
	token.TOKEN_AND:           {nodes.DataBool},
	token.TOKEN_OR:            {nodes.DataBool},
	token.TOKEN_BIT_AND:       {nodes.DataNumber},
	token.TOKEN_BIT_OR:        {nodes.DataNumber},
	token.TOKEN_BIT_XOR:       {nodes.DataNumber},
	token.TOKEN_BIT_NOT:       {nodes.DataNumber},
	token.TOKEN_LEFT_SHIFT:    {nodes.DataNumber},
	token.TOKEN_RIGHT_SHIFT:   {nodes.DataNumber},
	token.TOKEN_INCREMENT:     {},
	token.TOKEN_DECREMENT:     {},
	token.TOKEN_ASSIGN:        {nodes.DataNumber, nodes.DataBool, nodes.DataString},
	token.TOKEN_ASSIGN_ADD:    {nodes.DataNumber, nodes.DataString},
	token.TOKEN_ASSIGN_SUB:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_MUL:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_DIV:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_MOD:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_AND:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_OR:     {nodes.DataNumber},
	token.TOKEN_ASSIGN_XOR:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_LSHIFT: {nodes.DataNumber},
	token.TOKEN_ASSIGN_RSHIFT: {nodes.DataNumber},
}
var OperatorReturnedTypes = map[string][]nodes.DataType{
	token.TOKEN_PLUS:          {nodes.DataNumber, nodes.DataString},
	token.TOKEN_MINUS:         {nodes.DataNumber},
	token.TOKEN_MUL:           {nodes.DataNumber},
	token.TOKEN_DIV:           {nodes.DataNumber},
	token.TOKEN_MODULO:        {nodes.DataNumber},
	token.TOKEN_EQ:            {nodes.DataBool},
	token.TOKEN_NEQ:           {nodes.DataBool},
	token.TOKEN_LT:            {nodes.DataBool},
	token.TOKEN_LTE:           {nodes.DataBool},
	token.TOKEN_GT:            {nodes.DataBool},
	token.TOKEN_GTE:           {nodes.DataBool},
	token.TOKEN_AND:           {nodes.DataBool},
	token.TOKEN_OR:            {nodes.DataBool},
	token.TOKEN_BIT_AND:       {nodes.DataNumber},
	token.TOKEN_BIT_OR:        {nodes.DataNumber},
	token.TOKEN_BIT_XOR:       {nodes.DataNumber},
	token.TOKEN_BIT_NOT:       {nodes.DataNumber},
	token.TOKEN_LEFT_SHIFT:    {nodes.DataNumber},
	token.TOKEN_RIGHT_SHIFT:   {nodes.DataNumber},
	token.TOKEN_INCREMENT:     {},
	token.TOKEN_DECREMENT:     {},
	token.TOKEN_ASSIGN:        {},
	token.TOKEN_ASSIGN_ADD:    {},
	token.TOKEN_ASSIGN_SUB:    {},
	token.TOKEN_ASSIGN_MUL:    {},
	token.TOKEN_ASSIGN_DIV:    {},
	token.TOKEN_ASSIGN_MOD:    {},
	token.TOKEN_ASSIGN_AND:    {},
	token.TOKEN_ASSIGN_OR:     {},
	token.TOKEN_ASSIGN_XOR:    {},
	token.TOKEN_ASSIGN_LSHIFT: {},
	token.TOKEN_ASSIGN_RSHIFT: {},
}
var OperatorPrecedence = map[string]int{
	token.TOKEN_BIT_NOT:     9, // Unary bitwise NOT (~), highest precedence
	token.TOKEN_MUL:         7, // Multiplication (*)
	token.TOKEN_DIV:         7, // Division (/)
	token.TOKEN_MODULO:      7, // Modulo (%)
	token.TOKEN_PLUS:        6, // Addition (+)
	token.TOKEN_MINUS:       6, // Subtraction (-)
	token.TOKEN_BIT_AND:     5, // Bitwise AND (&)
	token.TOKEN_LEFT_SHIFT:  5, // Bitwise left shift (<<)
	token.TOKEN_RIGHT_SHIFT: 5, // Bitwise right shift (>>)
	token.TOKEN_BIT_XOR:     4, // Bitwise XOR (^)
	token.TOKEN_BIT_OR:      3, // Bitwise OR (|)
	token.TOKEN_EQ:          2, // Equal to (==)
	token.TOKEN_NEQ:         2, // Not equal to (!=)
	token.TOKEN_LT:          2, // Less than (<)
	token.TOKEN_LTE:         2, // Less than or equal to (<=)
	token.TOKEN_GT:          2, // Greater than (>)
	token.TOKEN_GTE:         2, // Greater than or equal to (>=)
	token.TOKEN_AND:         1, // Logical AND (&&)
	token.TOKEN_OR:          0, // Logical OR (||), lowest precedence
}

// AssigningOperators maps assignment token constants to their respective slices of nodes.DataType.
var AssigningOperators = map[string][]nodes.DataType{
	token.TOKEN_INCREMENT:     {},
	token.TOKEN_DECREMENT:     {},
	token.TOKEN_ASSIGN:        {nodes.DataNumber, nodes.DataBool, nodes.DataString},
	token.TOKEN_ASSIGN_ADD:    {nodes.DataNumber, nodes.DataString},
	token.TOKEN_ASSIGN_SUB:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_MUL:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_DIV:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_MOD:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_AND:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_OR:     {nodes.DataNumber},
	token.TOKEN_ASSIGN_XOR:    {nodes.DataNumber},
	token.TOKEN_ASSIGN_LSHIFT: {nodes.DataNumber},
	token.TOKEN_ASSIGN_RSHIFT: {nodes.DataNumber},
}

// Example function to get precedence of an operator
func GetOperatorPrecedence(op string) int {
	if precedence, exists := OperatorPrecedence[op]; exists {
		return precedence
	}
	return -1 // Return -1 if the operator does not exist
}
