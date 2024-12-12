package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("Parsing the program returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Got %d, which is %d -- statement must be longer than 3 identifiers or expressions", program.Statements, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		st := program.Statements[i]
		if !testLetStatement(t, st, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, st ast.Statement, name string) bool {
	if st.TokenLiteral() != "let" {
		t.Errorf("Expected let token but received %q", st.TokenLiteral())
		return false
	}
	letSt, ok := st.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected let statement but received %T", st)
		return false
	}
	if letSt.Name.Value != name {
		t.Errorf("Expected let statement value to be %s but received %s", letSt.Name.Value, name)
		return false
	}
	if letSt.Name.TokenLiteral() != name {
		t.Errorf("Expected token literal of %s but received %s", letSt.Name.TokenLiteral(), name)
		return false
	}
	return true
}
