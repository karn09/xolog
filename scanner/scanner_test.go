package scanner

import (
	"reflect"
	"testing"
	"xolog/token"
)

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
			want: &Scanner{source: "TEST", start: 0, current: 0, line: 1, tokens: []token.Token{}},
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
			fields: fields{source: "TE", start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.EOF,
					Lexeme:  "",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "Will handle double tokens.",
			fields: fields{source: "!=", start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
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
			fields: fields{source: "//te \n{", start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
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
		{
			name:   "Will return token array of proper length, containing proper values.",
			fields: fields{source: "({})", start: 0, current: 0, line: 1, tokens: []token.Token{}},
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
			fields: fields{source: "(\n{", start: 0, current: 0, line: 1, tokens: []token.Token{}},
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
			s := Scanner{
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
		want   Scanner
	}{
		{
			name:   "Will advance token, and log error.",
			fields: fields{source: "TE", start: 0, current: 0, line: 1, tokens: make([]token.Token, len("TE")+1)},
			want:   Scanner{source: "TE", start: 0, current: 1, line: 1, tokens: make([]token.Token, len("TE")+1), HadError: true},
		},
		{
			name:   "Will advance, and create matching token within first array index..",
			fields: fields{source: "{", start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: Scanner{source: "{", start: 0, current: 1, line: 1, tokens: []token.Token{
				{
					Type:    token.LEFT_BRACE,
					Lexeme:  "{",
					Literal: nil,
					Line:    1,
				},
			}, HadError: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
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
		{
			name: "match is false at end",
			fields: fields{
				source:   "",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{""},
			want: false,
		},
		{
			name: "match is false when unexpected",
			fields: fields{
				source:   "!",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{"&"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
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
		{
			name: "Return true at end",
			fields: fields{
				source:   "",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: true,
		},
		{
			name: "Return false when not at end end",
			fields: fields{
				source:   "{",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
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
			s := Scanner{
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
		{
			name: "Peek at end, returns EOF",
			fields: fields{
				source:   "",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: "\000",
		},
		{
			name: "Peek returns string",
			fields: fields{
				source:   "{",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: "{",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
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
			s := Scanner{
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

func TestScanner_string(t *testing.T) {
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
		want   Scanner
	}{
		{
			name: "Will add string token",
			fields: fields{
				source:   "'test'",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: Scanner{
				source:  "'test'",
				start:   0,
				current: 6,
				line:    1,
				tokens: []token.Token{
					{
						Type:    token.STRING,
						Lexeme:  "'test'",
						Literal: "test",
						Line:    1,
					},
				},
				HadError: false,
			},
		},
		{
			name: "Will handle newline within string",
			fields: fields{
				source:   "'test \n more'",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: Scanner{
				source:  "'test \n more'",
				start:   0,
				current: 13,
				line:    2,
				tokens: []token.Token{
					{
						Type:    token.STRING,
						Lexeme:  "'test \n more'",
						Literal: "test \n more",
						Line:    2,
					},
				},
				HadError: false,
			},
		},
		{
			name: "Will handle unterminated string",
			fields: fields{
				source:   "'test",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: Scanner{
				source:   "'test",
				start:    0,
				current:  5,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
		},
		{
			name: "Will handle quote strings",
			fields: fields{
				source:   `"h,ello"`,
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: Scanner{
				source:  `"h,ello"`,
				start:   0,
				current: 8,
				line:    1,
				tokens: []token.Token{
					{
						Type:    token.STRING,
						Lexeme:  `"h,ello"`,
						Literal: "h,ello",
						Line:    1,
					},
				},
				HadError: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Scanner{
				source:   tt.fields.source,
				start:    tt.fields.start,
				current:  tt.fields.current,
				line:     tt.fields.line,
				tokens:   tt.fields.tokens,
				HadError: tt.fields.HadError,
			}
			// string is called by scanToken as cases all start with opening quote
			s.scanToken()
			if got := s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("string() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_peekNext(t *testing.T) {
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
		{
			name: "Will peek on upcoming char",
			fields: fields{
				source:   `>=`,
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: "=",
		},
		{
			name: "Will return ending if out of bounds",
			fields: fields{
				source:   `>`,
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: "\000",
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
			if got := s.peekNext(); got != tt.want {
				t.Errorf("Scanner.peekNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_isDigit(t *testing.T) {
	type fields struct {
		source   string
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	type args struct {
		c string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Will be true if number encountered",
			fields: fields{
				source:   `1`,
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{"1"},
			want: true,
		},
		{
			name: "Will be false if non-number encountered",
			fields: fields{
				source:   `a`,
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{"a"},
			want: false,
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
			if got := s.isDigit(tt.args.c); got != tt.want {
				t.Errorf("Scanner.isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_number(t *testing.T) {
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
		// TODO: Add test cases.
		{
			name: "Will handle conversion of decimal string to float64",
			fields: fields{
				source:   "1.02",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: &Scanner{
				source:  "1.02",
				start:   0,
				current: 4,
				line:    1,
				tokens: []token.Token{
					{
						Type:    token.NUMBER,
						Lexeme:  "1.02",
						Literal: 1.02,
						Line:    1,
					},
				},
				HadError: false,
			},
		},
		{
			name: "Will handle number string conversion to float64",
			fields: fields{
				source:   "102",
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: &Scanner{
				source:  "102",
				start:   0,
				current: 3,
				line:    1,
				tokens: []token.Token{
					{
						Type:    token.NUMBER,
						Lexeme:  "102",
						Literal: float64(102),
						Line:    1,
					},
				},
				HadError: false,
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
			s.number()
			if got := s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("number() = %v, want %v", got, tt.want)
			}
		})
	}
}
