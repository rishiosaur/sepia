package evaluator

import (
	"sepia/ast"
	"sepia/objects"

	"github.com/go-git/go-git/v5/plumbing/object"
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
		return &object.Integer{Value: node.Value}

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
