package evaluator

import (
	"github.com/arjunmalhotra1/monkey-interpreter/ast"

	"github.com/arjunmalhotra1/monkey-interpreter/object"
)

func Eval(node ast.Node) object.Object {

	// "type" is the concrete data type that satisfies the interface. Like, Expression is an interface but
	// "ast.IntegerLiteral" is the data type that satisfies the interface.
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

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}
