package evaluator

import (
	"testing"

	"github.com/dmarticus/monkey/ast"
	"github.com/dmarticus/monkey/lexer"
	"github.com/dmarticus/monkey/object"
	"github.com/dmarticus/monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	// two types of program: one with content, one that's nil
	tests := []struct {
		program *ast.Program
	}{
		{
			program: p.ParseProgram(),
		},
		{
			program: nil,
		},
	}
	for _, test := range tests {
		return Eval(test.program)
	}

	return nil
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer, got=%T (%v) instead", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value; got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}
