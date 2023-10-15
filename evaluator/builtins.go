package evaluator

import "github.com/arjunmalhotra1/monkey-interpreter/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return NULL
		},
	},
}
