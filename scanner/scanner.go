package scanner

import (
	"strconv"
	"xolog/error"
	"xolog/token"
)

type Scanner struct {
	source   string
	start    int
	current  int
	line     int
	tokens   []token.Token
	HadError bool
}

// NewScanner accepts a string, and returns a pointer to the initialized Scanner struct.
func NewScanner(source string) *Scanner {
	return &Scanner{source: source, start: 0, current: 0, line: 1, tokens: []token.Token{}}
}

// ScanTokens will return an array of source length tokens of type Token.
func (s *Scanner) ScanTokens() []token.Token {
	for s.current < len(s.source) {
		s.start = s.current
		s.scanToken()
	}
	eof := token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: s.line}
	s.tokens = append(s.tokens, eof)
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case "(":
		s.addToken(token.LEFT_PAREN, nil)
	case ")":
		s.addToken(token.RIGHT_PAREN, nil)
	case "{":
		s.addToken(token.LEFT_BRACE, nil)
	case "}":
		s.addToken(token.RIGHT_BRACE, nil)
	case ",":
		s.addToken(token.COMMA, nil)
	case ".":
		s.addToken(token.DOT, nil)
	case "-":
		s.addToken(token.MINUS, nil)
	case "+":
		s.addToken(token.PLUS, nil)
	case ";":
		s.addToken(token.SEMICOLON, nil)
	case "*":
		s.addToken(token.STAR, nil)
	case "!":
		if s.match("=") {
			s.addToken(token.BANG_EQUAL, nil)
		} else {
			s.addToken(token.BANG, nil)
		}
	case "=":
		if s.match("=") {
			s.addToken(token.EQUAL_EQUAL, nil)
		} else {
			s.addToken(token.EQUAL, nil)
		}
	case "<":
		if s.match("=") {
			s.addToken(token.LESS_EQUAL, nil)
		} else {
			s.addToken(token.LESS, nil)
		}
	case ">":
		if s.match("=") {
			s.addToken(token.GREATER_EQUAL, nil)
		} else {
			s.addToken(token.GREATER, nil)
		}
	case `\`:
		if s.match(`\`) {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH, nil)
		}
	case " ":
	case "\t":
	case "\r":
	case "\n":
		s.line++
	case "'":
		s.string()
	case `"`:
		s.string()
	default:
		error.Error(s.line, "Unexpected character: "+c)
		s.HadError = true
	}
}

func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	text := string(s.source[s.current])
	if text != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) string() {
	for s.peek() != `"` && s.peek() != "'" && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		error.Error(s.line, "Unterminated string.")
		return
	}
	s.advance()
	val := s.source[s.start+1 : s.current-1]
	s.addToken(token.STRING, val)
}

func (s *Scanner) isAtEnd() bool {
	if s.current >= len(s.source) {
		return true
	}
	return false
}

// advance will consume the current token, and return the consumed token.
func (s *Scanner) advance() string {
	s.current++
	return string(s.source[s.current-1])
}

// peek will return the current token, without consuming
func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\000"
	}
	c := string(s.source[s.current])
	return c
}

func (s *Scanner) peekNext() string {
	if s.current+1 >= len(s.source) {
		return "\000"
	}
	c := string(s.source[s.current+1])
	return c
}

func (s *Scanner) isDigit(c string) bool {
	i, err := strconv.Atoi(c)
	if err != nil {
		return false
	}
	return i >= 0 && i <= 9
}

// addToken will add token type and lexeme to returned tokens array.
func (s *Scanner) addToken(tokenType token.TokenType, literal interface{}) []token.Token {
	text := s.source[s.start:s.current]
	token := token.Token{Type: tokenType, Lexeme: text, Literal: literal, Line: s.line}
	s.tokens = append(s.tokens, token)
	return s.tokens
}
