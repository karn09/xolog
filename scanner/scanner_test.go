package scanner

import (
	"testing"
	"xolog/token"
)

func TestAdvance(t *testing.T) {
	scanner := NewScanner("{(")
	leftBrace := scanner.advance()
	if leftBrace != "{" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", leftBrace, "{")
	}
	leftParen := scanner.advance()
	if leftParen != "(" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", leftParen, "(")
	}
}

func TestScanTokens(t *testing.T) {
	scanner := NewScanner("({{(")
	tokens := scanner.ScanTokens()
	if len(tokens) != 5 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", len(tokens), 5)
	}
	if tokens[0].Lexeme != "(" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[0].Lexeme, "(")
	}
	if tokens[1].Lexeme != "{" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[1].Lexeme, "{")
	}
	if tokens[3].Lexeme != "(" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[3].Lexeme, "(")
	}
	if tokens[4].Type != 38 {
		t.Errorf("Type was incorrect, got: %d, want: %d.", tokens[4].Type, 38)
	}
}
func TestAddToken(t *testing.T) {
	scanner := NewScanner("(")
	scanner.advance()
	tokens := scanner.addToken(token.LEFT_PAREN, nil)
	if len(tokens) != 2 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", len(tokens), 2)
	}
	if tokens[0].Lexeme != "(" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[0].Lexeme, "(")
	}
	if tokens[1].Lexeme != "" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[1].Lexeme, "")
	}
	if tokens[1].Type != 0 {
		t.Errorf("Character was incorrect, got: %d, want: %d.", tokens[1].Type, 0)
	}
}
func TestDoubleToken(t *testing.T) {
	scanner := NewScanner("!=")
	tokens := scanner.ScanTokens()
	if len(tokens) != 3 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", len(tokens), 3)
	}
	if tokens[1].Lexeme != "!=" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[1].Lexeme, "!=")
	}
}
func TestNewLineToken(t *testing.T) {
	scanner := NewScanner("\n{\n}")
	tokens := scanner.ScanTokens()
	if len(tokens) != 5 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", len(tokens), 5)
	}
	if tokens[1].Lexeme != "{" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[1].Lexeme, "{")
	}
	if tokens[1].Line != 2 {
		t.Errorf("Line was incorrect, got: %d, want: %d.", tokens[1].Line, 2)
	}
	if tokens[3].Line != 3 {
		t.Errorf("Line was incorrect, got: %d, want: %d.", tokens[1].Line, 3)
	}
	if tokens[3].Lexeme != "}" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[3].Lexeme, "}")
	}
}
func TestCommentTokens(t *testing.T) {
	scanner := NewScanner("\\\\ comment\n{}")
	tokens := scanner.ScanTokens()
	if tokens[0].Lexeme == "\\" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[0].Lexeme, "")
	}
	if tokens[1].Lexeme == "\\" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[1].Lexeme, "")
	}
	if len(tokens) != 14 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", len(tokens), 14)
	}
	if tokens[11].Lexeme != "{" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", tokens[11].Lexeme, "{")
	}
	if tokens[13].Line != 2 {
		t.Errorf("Line was incorrect, got: %d, want: %d.", tokens[13].Line, 2)
	}
}

func TestPeek(t *testing.T) {
	scanner := NewScanner("test")
	c := scanner.peek()
	if c != "t" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", c, "t")
	}
	c2 := scanner.peek()
	if c2 != "t" {
		t.Errorf("Character was incorrect, got: %s, want: %s.", c2, "t")
	}
}
