package evaluator

import (
	"fmt"
	"sepia/objects"
)

var builtins = map[string]*objects.Builtin{
	"len": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("how are you this stupid. it's the length function, ya idiot. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *objects.String:
				return &objects.Integer{Value: int64(len(arg.Value))}
			case *objects.Array:
				return &objects.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		}},
	"typeof": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments supplied. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *objects.Boolean:
				return &objects.String{Value: fmt.Sprintf("%T", arg.Value)}

			case *objects.Integer:
				return &objects.String{Value: "int"}
			case *objects.String:
				return &objects.String{Value: fmt.Sprintf("%T", arg.Value)}
			default:
				return newError("argument to `typeof` not supported, got %s", args[0].Type())
			}
		},
	},
	"print": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
	"string": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments supplied. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *objects.String:
				return arg
			case *objects.Integer:
				return &objects.String{Value: fmt.Sprintf("%v", arg.Value)}
			case *objects.Boolean:
				return &objects.String{Value: fmt.Sprintf("%v", arg.Value)}
			default:
				return newError("argument to `string` not supported, got %s.", args[0].Type())
			}
		},
	},
	"bool": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments supplied. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *objects.String:
				return &objects.Boolean{Value: arg.Value != ""}
			case *objects.Integer:
				return &objects.Boolean{Value: arg.Value != 0}
			case *objects.Boolean:
				return arg
			default:
				return newError("argument to `len` not supported, got %s.", args[0].Type())
			}
		},
	},
	"int": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments supplied. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *objects.String:
				return &objects.Integer{Value: int64(len(arg.Value))}
			case *objects.Integer:
				return arg
			case *objects.Boolean:
				bitSet := arg.Value
				bitSetVar := int64(0)
				if bitSet {
					bitSetVar = 1
				}
				return &objects.Integer{Value: bitSetVar}
			default:
				return newError("argument to `len` not supported, got %s.", args[0].Type())
			}
		},
	},
	"first": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != objects.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*objects.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return NULL
		},
	},

	"last": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != objects.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*objects.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},

	"append": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != objects.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*objects.Array)
			length := len(arr.Elements)
			newElements := make([]objects.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &objects.Array{Elements: newElements}
		},
	},

	"rest": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != objects.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*objects.Array)
			length := len(arr.Elements)

			if length > 0 {
				newElements := make([]objects.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &objects.Array{Elements: newElements}
			}
			return NULL
		},
	},
}
