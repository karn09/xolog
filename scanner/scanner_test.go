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
			want: &Scanner{source: []byte("TEST"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
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
		source   []byte
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
			name:   "Unexpected characters, Return token array containing EOF only.",
			fields: fields{source: []byte("TE"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "BANG_EQUAL, BANG token parsed with EOF.",
			fields: fields{source: []byte("!= !"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.BANG_EQUAL,
					Lexeme:  "!=",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.BANG,
					Lexeme:  "!",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "COMMA token parsed with EOF.",
			fields: fields{source: []byte(","), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.COMMA,
					Lexeme:  ",",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "DOT token parsed with EOF.",
			fields: fields{source: []byte("."), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.DOT,
					Lexeme:  ".",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "LESS_EQUAL, LESS, GREATER, GREATER_EQUAL token parsed with EOF.",
			fields: fields{source: []byte("<= < > >="), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.LESS_EQUAL,
					Lexeme:  "<=",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.LESS,
					Lexeme:  "<",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.GREATER,
					Lexeme:  ">",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.GREATER_EQUAL,
					Lexeme:  ">=",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "EQUAL_EQUAL token parsed with EOF.",
			fields: fields{source: []byte("=="), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.EQUAL_EQUAL,
					Lexeme:  "==",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "EQUAL_EQUAL,EQUAL token parsed with EOF.",
			fields: fields{source: []byte("==="), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.EQUAL_EQUAL,
					Lexeme:  "==",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EQUAL,
					Lexeme:  "=",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "Will handle comment tokens.",
			fields: fields{source: []byte("//te \n{"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.LEFT_BRACE,
					Lexeme:  "{",
					Literal: nil,
					Line:    2,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    2,
				},
			},
		},
		{
			name:   "LEFT_PAREN, LEFT_BRACE, RIGHT_BRACE, RIGHT_PAREN, EOF",
			fields: fields{source: []byte("({})"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
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
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "LEFT_PAREN, LEFT_BRACE, EOF token array, with line = 2 after newline",
			fields: fields{source: []byte("(\n{"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
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
					Lexeme:  "\000",
					Literal: nil,
					Line:    2,
				},
			},
		},
		{
			name:   "MINUS, PLUS, EOF token array",
			fields: fields{source: []byte("-+"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.MINUS,
					Lexeme:  "-",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.PLUS,
					Lexeme:  "+",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "SEMICOLON, STAR, EOF token array",
			fields: fields{source: []byte(";*"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.SEMICOLON,
					Lexeme:  ";",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.STAR,
					Lexeme:  "*",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
				},
			},
		},
		{
			name:   "SLASH, EOF token array, ignore \\ comment",
			fields: fields{source: []byte("\\,\\\\test"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: []token.Token{
				{
					Type:    token.SLASH,
					Lexeme:  "\\",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.COMMA,
					Lexeme:  ",",
					Literal: nil,
					Line:    1,
				},
				{
					Type:    token.EOF,
					Lexeme:  "\000",
					Literal: nil,
					Line:    1,
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
		source   []byte
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
			fields: fields{source: []byte("TE"), start: 0, current: 0, line: 1, tokens: make([]token.Token, len("TE")+1)},
			want:   Scanner{source: []byte("TE"), start: 0, current: 1, line: 1, tokens: make([]token.Token, len("TE")+1), HadError: true},
		},
		{
			name:   "Will advance, and create matching token within first array index..",
			fields: fields{source: []byte("{"), start: 0, current: 0, line: 1, tokens: []token.Token{}},
			want: Scanner{source: []byte("{"), start: 0, current: 1, line: 1, tokens: []token.Token{
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
		source   []byte
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	type args struct {
		expected rune
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
				source:   []byte(""),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{' '},
			want: false,
		},
		{
			name: "match is false when unexpected",
			fields: fields{
				source:   []byte("!"),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{'&'},
			want: false,
		},
		{
			name: "match is true",
			fields: fields{
				source:   []byte("{"),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{'{'},
			want: true,
		},
		{
			name: "match is true",
			fields: fields{
				source:   []byte("{("),
				start:    0,
				current:  1,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{'('},
			want: true,
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
		source   []byte
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
				source:   []byte(""),
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
				source:   []byte("{"),
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
		source   []byte
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{
			name: "Return empty rune",
			fields: fields{
				source:   []byte(" "),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: ' ',
		},
		{
			name: "Return rune",
			fields: fields{
				source:   []byte("{"),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: '{',
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
			if got := s.advance(); got != tt.want {
				t.Errorf("Scanner.advance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_peek(t *testing.T) {
	type fields struct {
		source   []byte
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{
			name: "Peek at end, returns EOF",
			fields: fields{
				source:   []byte(""),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: '\000',
		},
		{
			name: "Peek on number, returns number",
			fields: fields{
				source:   []byte("1"),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: '1',
		},
		{
			name: "Peek returns string",
			fields: fields{
				source:   []byte("{"),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: '{',
		},
		{
			name: "Peek returns string",
			fields: fields{
				source:   []byte("{("),
				start:    0,
				current:  1,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: '(',
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
		source   []byte
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	type args struct {
		tokenType token.TokenType
		literal   []rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []token.Token
	}{
		{
			name: "Add comma token.",
			fields: fields{
				source:   []byte(","),
				start:    0,
				current:  1,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{token.COMMA, nil},
			want: []token.Token{
				{
					Type:    token.COMMA,
					Lexeme:  ",",
					Literal: nil,
					Line:    1,
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
			if got := s.addToken(tt.args.tokenType, tt.args.literal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.addToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner_string(t *testing.T) {
	type fields struct {
		source   []byte
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
				source:   []byte("'test'"),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: Scanner{
				source:  []byte("'test'"),
				start:   0,
				current: 5,
				line:    1,
				tokens: []token.Token{
					{
						Type:    token.STRING,
						Lexeme:  "test",
						Literal: []rune("test"),
						Line:    1,
					},
				},
				HadError: false,
			},
		},
		{
			name: "Will handle newline within string",
			fields: fields{
				source:   []byte("\"test \n more\""),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: Scanner{
				source:  []byte("\"test \n more\""),
				start:   0,
				current: 12,
				line:    2,
				tokens: []token.Token{
					{
						Type:    token.STRING,
						Lexeme:  "test \n more",
						Literal: []rune("test \n more"),
						Line:    2,
					},
				},
				HadError: false,
			},
		},
		{
			name: "Will handle unterminated string",
			fields: fields{
				source:   []byte("'test"),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: Scanner{
				source:   []byte("'test"),
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
				source:   []byte(`"h,ello"`),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: Scanner{
				source:  []byte(`"h,ello"`),
				start:   0,
				current: 7,
				line:    1,
				tokens: []token.Token{
					{
						Type:    token.STRING,
						Lexeme:  `h,ello`,
						Literal: []rune("h,ello"),
						Line:    1,
					},
				},
				HadError: false,
			},
		},
		{
			name: "Will handle valid token after string",
			fields: fields{
				source:   []byte(`"h,ello"{`),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: Scanner{
				source:  []byte(`"h,ello"{`),
				start:   0,
				current: 7,
				line:    1,
				tokens: []token.Token{
					{
						Type:    token.STRING,
						Lexeme:  `h,ello`,
						Literal: []rune("h,ello"),
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
		source   []byte
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{
			name: "Will peek on upcoming char",
			fields: fields{
				source:   []byte(`>=`),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: '=',
		},
		{
			name: "Will return ending if out of bounds",
			fields: fields{
				source:   []byte(`>`),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: '\000',
		},
		{
			name: "Will peek ahead if current is not on starting position.",
			fields: fields{
				source:   []byte(`>=<`),
				start:    1,
				current:  1,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: '<',
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
		source   []byte
		start    int
		current  int
		line     int
		tokens   []token.Token
		HadError bool
	}
	type args struct {
		c rune
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
				source:   []byte(`1`),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{'1'},
			want: true,
		},
		{
			name: "Will be false if non-number encountered",
			fields: fields{
				source:   []byte(`a`),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			args: args{'a'},
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
		source   []byte
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
				source:   []byte("1.02"),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: &Scanner{
				source:  []byte("1.02"),
				start:   0,
				current: 4,
				line:    1,
				tokens: []token.Token{
					{
						Type:    token.NUMBER,
						Lexeme:  "1.02",
						Literal: []rune("1.02"),
						Line:    1,
					},
					{
						Type:    token.EOF,
						Lexeme:  "\000",
						Literal: nil,
						Line:    1,
					},
				},
				HadError: false,
			},
		},
		{
			name: "Will handle number string conversion to float64",
			fields: fields{
				source:   []byte("102"),
				start:    0,
				current:  0,
				line:     1,
				tokens:   []token.Token{},
				HadError: false,
			},
			want: &Scanner{
				source:  []byte("102"),
				start:   0,
				current: 3,
				line:    1,
				tokens: []token.Token{
					{
						Type:    token.NUMBER,
						Lexeme:  "102",
						Literal: []rune("102"),
						Line:    1,
					},
					{
						Type:    token.EOF,
						Lexeme:  "\000",
						Literal: nil,
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
			s.ScanTokens()
			if got := s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("number() = %v, want %v", got, tt.want)
			}
		})
	}
}
