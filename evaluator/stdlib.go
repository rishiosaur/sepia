package evaluator

import "sepia/objects"

var builtins = map[string]*objects.Builtin{
	"len": &objects.Builtin{
		Fn: func(args ...objects.Object) objects.Object {
			if len(args) != 1 {
				return newError("how are you this stupid. it's the length function, ya idiot. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *objects.String:
				return &objects.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		}},
}
