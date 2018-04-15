package scanner

import (
	"encoding/json"
	"fmt"
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
	tokens := scanner.scanTokens()
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
	tokens := scanner.scanTokens()
	for _, v := range tokens {
		js, _ := json.Marshal(v)
		fmt.Println(string(js))
	}
	if len(tokens) != 3 {
		t.Errorf("Length was incorrect, got: %d, want: %d.", len(tokens), 3)
	}

}
