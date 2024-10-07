package nodes

import "ast-operation-parser/lexer/token"

type DataType int

const (
	DataNumber DataType = iota
	DataString
	DataBool
)

var TokenToDataType = map[string]DataType{
	token.TOKEN_STRING: DataString,
	token.TOKEN_NUMBER: DataNumber,
	token.TOKEN_BOOL:   DataBool,
}

func GetDataType(t ASTNode) []DataType {
	switch t.Type() {
	case Literal:
		return []DataType{t.(LiteralNode).DataType}
	case BinaryOperator:
		return t.(BinaryOperation).ReturnedDataTypes
	case Variable:
		return []DataType{}
	}
	return nil
}
