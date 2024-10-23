package parser

import (
	"testing"

	"github.com/medragneel/lex/ast"
	"github.com/medragneel/lex/lexer"
)

func TestLetStatement(t *testing.T) {
	// Consider moving this to a separate function or file if used in multiple tests
	input := `
    let x = 5;
    let y = 5;
    let foo = 23332;
    `
	l := lexer.New(input)
	p := New(l)
	program := p.parseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("parseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program statements does not contain 3 statements. got %d", len(program.Statements))
	}

	// Consider using a map for O(1) lookup instead of slice for large datasets
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	// Consider using t.Run() for subtests, allowing for better parallelization
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() not 'let' got %q", s.TokenLiteral())
		return false
	}
	letstmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.Statement got=%T", s)
		return false // Missing return statement here
	}
	if letstmt.Name.Value != name {
		t.Errorf("letstmt.Name not %s, got=%s", name, letstmt.Name.Value)
		return false
	}
	if letstmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s', got=%s", name, letstmt.Name.TokenLiteral())
		return false
	}
	return true
}

func TestReturnStatement(t *testing.T) {
	// Consider moving this to a separate function or file if used in multiple tests
	input := `
    return 5;
    return 5;
    return 23332;
    `
	l := lexer.New(input)
	p := New(l)
	program := p.parseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("parseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program statements does not contain 3 statements. got %d", len(program.Statements))
	}

	// Consider using a map for O(1) lookup instead of slice for large datasets
	// tests := []struct {
	// 	expectedIdentifier string
	// }{
	// 	{"x"},
	// 	{"y"},
	// 	{"foo"},
	// }

	// Consider using t.Run() for subtests, allowing for better parallelization
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'retrun' got=%q", returnStmt.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has  %d Errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser Error: %q", msg)
	}
	t.FailNow()

}
