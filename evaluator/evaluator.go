package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	//Recurisvely evaluate each Statement of the Program
	case *ast.Program:
		return evalStatements(node.Statements)

	//Evaluate an Expression Statement
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	//Expressions
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
