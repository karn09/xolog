package token

import "testing"

func TestToken_toString(t *testing.T) {
	type fields struct {
		Type    TokenType
		Lexeme  string
		Literal interface{}
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
			want: "0 ( %!s(<nil>)",
		},
		{
			name: "RIGHT_PAREN token to string.",
			fields: fields{
				Type:    RIGHT_PAREN,
				Lexeme:  ")",
				Literal: nil,
				Line:    0,
			},
			want: "1 ) %!s(<nil>)",
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
			if got := token.toString(); got != tt.want {
				t.Errorf("Token.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
