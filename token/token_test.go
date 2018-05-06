package token

import "testing"

func TestToken_toString(t *testing.T) {
	type fields struct {
		Type    TokenType
		Lexeme  string
		Literal []rune
		Line    int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "LEFT_PAREN token to string.",
			fields: fields{
				Type:    LEFT_PAREN,
				Lexeme:  "(",
				Literal: nil,
				Line:    0,
			},
			want: "0 ( []",
		},
		{
			name: "STRING token to string.",
			fields: fields{
				Type:    STRING,
				Lexeme:  "Hello",
				Literal: []rune("Hello"),
				Line:    0,
			},
			want: "20 Hello [H e l l o]",
		},
		{
			name: "Number token to string.",
			fields: fields{
				Type:    NUMBER,
				Lexeme:  "123",
				Literal: []rune("123"),
				Line:    0,
			},
			want: "21 123 [1 2 3]",
		},
		{
			name: "Number token to string.",
			fields: fields{
				Type:    NUMBER,
				Lexeme:  "123",
				Literal: nil,
				Line:    0,
			},
			want: "21 123 []",
		},
		{
			name: "LEFT_PAREN token to string.",
			fields: fields{
				Type:    LEFT_PAREN,
				Lexeme:  "(",
				Literal: nil,
				Line:    0,
			},
			want: "0 ( []",
		},
		{
			name: "RIGHT_PAREN token to string.",
			fields: fields{
				Type:    RIGHT_PAREN,
				Lexeme:  ")",
				Literal: nil,
				Line:    0,
			},
			want: "1 ) []",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := &Token{
				Type:    tt.fields.Type,
				Lexeme:  tt.fields.Lexeme,
				Literal: tt.fields.Literal,
				Line:    tt.fields.Line,
			}
			if got := token.String(); got != tt.want {
				t.Errorf("Token.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
