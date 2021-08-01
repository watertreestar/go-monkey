package parser

import (
	"github.com/watertreestar/go-monkey/ast"
	"github.com/watertreestar/go-monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {

	t.Run("valid let statement", func(t *testing.T) {
		input := `
let x = 5;
let u = 10;
let foobar = 76983;
`
		l := lexer.NewLexer(input)
		parser := NewParse(l)
		program := parser.parseProgram()
		checkErrors(t, parser)
		if program == nil {
			t.Fatalf("Prarse program return nil")
		}
		if len(program.Statements) != 3 {
			t.Fatalf("Program does not contians 3 statements, actully got %v", len(program.Statements))
		}

		tests := []struct {
			expectedIdentifier string
		}{
			{"x"},
			{"u"},
			{"foobar"},
		}

		for i, tt := range tests {
			stmt := program.Statements[i]
			if !testLetStatement(t, stmt, tt.expectedIdentifier) {
				return
			}
		}
	})

	t.Run("invalid let statement", func(t *testing.T) {
		input := `
let x = 5;
let u = 10;
let 76983`
		l := lexer.NewLexer(input)
		parser := NewParse(l)
		program := parser.parseProgram()
		if len(parser.Errors()) == 0 {
			t.FailNow()
		}
		if program == nil {
			t.Fatalf("Prarse program return nil")
		}

	})

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func checkErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, err := range errors {
		t.Errorf("parse error : %q", err)
	}
	t.FailNow()
}

func TestReturnStatement(t *testing.T) {
	input := `
return 5;
return 10;
return 988989;
`
	l := lexer.NewLexer(input)
	p := NewParse(l)

	program := p.parseProgram()
	checkErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program statement does not contians 3 statements.got %d", len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStmt, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement not *ast.ReturnStatement")
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not be 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}
