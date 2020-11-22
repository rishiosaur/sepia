package evaluator

import (
	"sepia/ast"
	"sepia/objects"
)

var (
	TRUE  = &objects.Boolean{Value: true}
	FALSE = &objects.Boolean{Value: false}
	NULL  = &objects.Null{}
)

func Eval(node ast.Node) objects.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// Expressions
	case *ast.IntegerLiteral:
		return &objects.Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return toBool(node.Value)

	}
	return nil
}

func evalStatements(stmts []ast.Statement) objects.Object {
	var result objects.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}

func toBool(input bool) *objects.Boolean {
	if input {
		return TRUE
	}

	return FALSE
}
