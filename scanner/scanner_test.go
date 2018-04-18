package scanner

import (
	"reflect"
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

func TestNewScanner(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want *Scanner
	}{
		{
			name: "Will return struct of proper shape.",
			args: args{"TEST"},
			want: &Scanner{source: "TEST", start: 0, current: 0, line: 1, tokens: make([]token.Token, len("TEST")+1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScanner(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_ScanTokens(t *testing.T) {
	type fields struct {
		source   string
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []token.Token
	}{
		{
			name:   "Will return token array of proper length, containing proper values.",
			fields: fields{source: "TE", start: 0, current: 0, line: 1, tokens: make([]token.Token, len("TE")+1)},
			want:   append(make([]token.Token, len("TE")), token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: 1}),
		},
		{
			name:   "Will handle double tokens.",
			fields: fields{source: "!=", start: 0, current: 0, line: 1, tokens: make([]token.Token, len("!=")+1)},
			want: []token.Token{
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
				{
					Type:    token.BANG_EQUAL,
					Lexeme:  "!=",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "Will handle comment tokens.",
			fields: fields{source: "//te \n{", start: 0, current: 0, line: 1, tokens: make([]token.Token, len("// te \n{")+1)},
			want: []token.Token{
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
				{
					Type:    token.LEFT_BRACE,
					Lexeme:  "{",
					Literal: nil,
					Line:    2,
				},
				{
					Type:    token.EOF,
					Lexeme:  "",
					Literal: nil,
					Line:    2,
				},
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
			},
		},
		{
			name:   "Will return token array of proper length, containing proper values.",
			fields: fields{source: "({})", start: 0, current: 0, line: 1, tokens: make([]token.Token, len("({})")+1)},
			want: []token.Token{
				{
					Type:    token.LEFT_PAREN,
					Lexeme:  "(",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.LEFT_BRACE,
					Lexeme:  "{",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.RIGHT_BRACE,
					Lexeme:  "}",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.RIGHT_PAREN,
					Lexeme:  ")",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "Will return token array of proper length, containing proper values, and handle newline",
			fields: fields{source: "(\n{", start: 0, current: 0, line: 1, tokens: make([]token.Token, len("(\n{")+1)},
			want: []token.Token{
				{
					Type:    token.LEFT_PAREN,
					Lexeme:  "(",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
				{
					Type:    token.LEFT_BRACE,
					Lexeme:  "{",
					Literal: nil,
					Line:    2,
				},
				{
					Type:    token.EOF,
					Lexeme:  "",
					Literal: nil,
					Line:    2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				source:   tt.fields.source,
				start:    tt.fields.start,
				current:  tt.fields.current,
				line:     tt.fields.line,
				tokens:   tt.fields.tokens,
				HadError: tt.fields.HadError,
			}
			if got := s.ScanTokens(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.ScanTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_scanToken(t *testing.T) {
	type fields struct {
		source   string
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Scanner
	}{
		{
			name:   "Will advance token, and log error.",
			fields: fields{source: "TE", start: 0, current: 0, line: 1, tokens: make([]token.Token, len("TE")+1)},
			want:   &Scanner{source: "TE", start: 0, current: 1, line: 1, tokens: make([]token.Token, len("TE")+1), HadError: true},
		},
		{
			name:   "Will advance, and create matching token within first array index..",
			fields: fields{source: "{", start: 0, current: 0, line: 1, tokens: make([]token.Token, len("{")+1)},
			want: &Scanner{source: "{", start: 0, current: 1, line: 1, tokens: []token.Token{
				{
					Type:    token.LEFT_BRACE,
					Lexeme:  "{",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    0,
					Lexeme:  "",
					Literal: nil,
					Line:    0,
				},
			}, HadError: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				source:   tt.fields.source,
				start:    tt.fields.start,
				current:  tt.fields.current,
				line:     tt.fields.line,
				tokens:   tt.fields.tokens,
				HadError: tt.fields.HadError,
			}
			s.scanToken()
			if got := s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanToken() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestScanner_match(t *testing.T) {
	type fields struct {
		source   string
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	type args struct {
		expected string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				source:   tt.fields.source,
				start:    tt.fields.start,
				current:  tt.fields.current,
				line:     tt.fields.line,
				tokens:   tt.fields.tokens,
				HadError: tt.fields.HadError,
			}
			if got := s.match(tt.args.expected); got != tt.want {
				t.Errorf("Scanner.match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_isAtEnd(t *testing.T) {
	type fields struct {
		source   string
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				source:   tt.fields.source,
				start:    tt.fields.start,
				current:  tt.fields.current,
				line:     tt.fields.line,
				tokens:   tt.fields.tokens,
				HadError: tt.fields.HadError,
			}
			if got := s.isAtEnd(); got != tt.want {
				t.Errorf("Scanner.isAtEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_advance(t *testing.T) {
	type fields struct {
		source   string
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				source:   tt.fields.source,
				start:    tt.fields.start,
				current:  tt.fields.current,
				line:     tt.fields.line,
				tokens:   tt.fields.tokens,
				HadError: tt.fields.HadError,
			}
			if got := s.advance(); got != tt.want {
				t.Errorf("Scanner.advance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_peek(t *testing.T) {
	type fields struct {
		source   string
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				source:   tt.fields.source,
				start:    tt.fields.start,
				current:  tt.fields.current,
				line:     tt.fields.line,
				tokens:   tt.fields.tokens,
				HadError: tt.fields.HadError,
			}
			if got := s.peek(); got != tt.want {
				t.Errorf("Scanner.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_addToken(t *testing.T) {
	type fields struct {
		source   string
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	type args struct {
		tokenType token.TokenType
		literal   interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []token.Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				source:   tt.fields.source,
				start:    tt.fields.start,
				current:  tt.fields.current,
				line:     tt.fields.line,
				tokens:   tt.fields.tokens,
				HadError: tt.fields.HadError,
			}
			if got := s.addToken(tt.args.tokenType, tt.args.literal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.addToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
