package evaluator

import (
	"fmt"
	"sepia/ast"
	"sepia/objects"
)

func newError(format string, a ...interface{}) *objects.Error {
	return &objects.Error{Message: fmt.Sprintf(format, a...)}
}

var (
	TRUE  = &objects.Boolean{Value: true}
	FALSE = &objects.Boolean{Value: false}
	NULL  = &objects.Null{}
)

func Eval(node ast.Node, machine *objects.Machine) objects.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements, machine)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, machine)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, machine)
		if isError(val) {
			return val
		}
		return &objects.ReturnValue{Value: val}
	case *ast.BlockStatement:
		return evalBlockStatement(node, machine)
	case *ast.LetStatement:
		val := Eval(node.Value, machine)
		if isError(val) {
			return val
		}
		machine.Set(node.Name.Value, val)

	// Expressions
	case *ast.IntegerLiteral:
		return &objects.Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return toBool(node.Value)
	case *ast.CallExpression:
		function := Eval(node.Function, machine)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, machine)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)
	case *ast.StringLiteral:
		return &objects.String{Value: node.Value}

	case *ast.PrefixExpression:
		right := Eval(node.Right, machine)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		right := Eval(node.Right, machine)
		if isError(right) {
			return right
		}
		left := Eval(node.Left, machine)
		if isError(left) {
			return left
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.IfExpression:
		return evalIfExpression(node, machine)
	case *ast.Identifier:
		return evalIdentifier(node, machine)
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &objects.Function{Parameters: params, Machine: machine, Body: body}
	default:
		return newError("I literally have no clue wtf that is. RTFM pls.")
	}
	return nil
}

func applyFunction(fn objects.Object, args []objects.Object) objects.Object {
	switch fn := fn.(type) {
	case *objects.Function:
		extendedLocMachine := extendLocalMachine(fn, args)
		evaluated := Eval(fn.Body, extendedLocMachine)
		return unwrapReturnValue(evaluated)
	case *objects.Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
}

func extendLocalMachine(fn *objects.Function, args []objects.Object,
) *objects.Machine {
	machine := objects.NewLocalMachine(fn.Machine)
	for paramIdx, param := range fn.Parameters {
		machine.Set(param.Value, args[paramIdx])
	}
	return machine
}
func unwrapReturnValue(obj objects.Object) objects.Object {
	if returnValue, ok := obj.(*objects.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalExpressions(
	expressions []ast.Expression,
	machine *objects.Machine,
) []objects.Object {
	var result []objects.Object

	for _, e := range expressions {
		evaluated := Eval(e, machine)
		if isError(evaluated) {
			return []objects.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement, machine *objects.Machine) objects.Object {
	var result objects.Object

	for _, statement := range block.Statements {
		result = Eval(statement, machine)

		if result != nil {
			rt := result.Type()
			if rt == objects.RETURN_VALUE_OBJ || rt == objects.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func isError(obj objects.Object) bool {
	if obj != nil {
		return obj.Type() == objects.ERROR_OBJ
	}

	return false
}
func evalProgram(program *ast.Program, machine *objects.Machine) objects.Object {
	var result objects.Object

	for _, statement := range program.Statements {
		result = Eval(statement, machine)
		switch result := result.(type) {
		case *objects.ReturnValue:
			return result.Value
		case *objects.Error:
			return result
		}
	}

	return result
}

func evalIdentifier(node *ast.Identifier, machine *objects.Machine) objects.Object {
	if val, ok := machine.Get(node.Value); ok {
		return val
	}
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}
	return newError("identifier not found: " + node.Value)
}

func evalStatements(stmts []ast.Statement, machine *objects.Machine) objects.Object {
	var result objects.Object
	for _, statement := range stmts {
		result = Eval(statement, machine)

		if returnValue, ok := result.(*objects.ReturnValue); ok {
			return returnValue.Value
		}
	}
	return result
}

func toBool(input bool) *objects.Boolean {
	if input {
		return TRUE
	}

	return FALSE
}

func evalPrefixExpression(operator string, right objects.Object) objects.Object {
	switch operator {
	case "!":
		return evalNegationOpExpression(right)
	case "-":
		return evalMinusOpExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalNegationOpExpression(right objects.Object) objects.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusOpExpression(right objects.Object) objects.Object {
	if right.Type() != objects.INTEGER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*objects.Integer).Value

	return &objects.Integer{Value: -value}
}

func evalInfixExpression(
	operator string,
	left, right objects.Object,
) objects.Object {
	switch {
	case left.Type() == objects.INTEGER_OBJ && right.Type() == objects.INTEGER_OBJ:
		return evalIntInfixExpression(operator, left, right)
	case operator == "==":
		return toBool(left == right)
	case operator == "!=":
		return toBool(left != right)

	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",
			left.Type(), operator, right.Type())
	case left.Type() == objects.STRING_OBJ && right.Type() == objects.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)

	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())

	}
}

func evalStringInfixExpression(operator string,
	left, right objects.Object,
) objects.Object {
	if operator != "+" {
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
	leftVal := left.(*objects.String).Value
	rightVal := right.(*objects.String).Value
	return &objects.String{Value: leftVal + rightVal}
}

func evalIntInfixExpression(
	operator string,
	left, right objects.Object,
) objects.Object {
	leftVal := left.(*objects.Integer).Value
	rightVal := right.(*objects.Integer).Value

	switch operator {
	case "+":
		return &objects.Integer{Value: leftVal + rightVal}
	case "-":
		return &objects.Integer{Value: leftVal - rightVal}
	case "*":
		return &objects.Integer{Value: leftVal * rightVal}
	case "/":
		return &objects.Integer{Value: leftVal / rightVal}
	case "<":
		return toBool(leftVal < rightVal)
	case ">":
		return toBool(leftVal > rightVal)
	case "==":
		return toBool(leftVal == rightVal)
	case "!=":
		return toBool(leftVal != rightVal)

	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIfExpression(ifExp *ast.IfExpression, machine *objects.Machine) objects.Object {
	condition := Eval(ifExp.Condition, machine)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ifExp.Consequence, machine)
	} else if ifExp.Alternative != nil {
		return Eval(ifExp.Alternative, machine)
	} else {
		return NULL
	}
}

func isTruthy(obj objects.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
