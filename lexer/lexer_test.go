package lexer

import (
	"testing"

	"github.com/ytakaya/nasparse/token"
)

func TestNextToken(t *testing.T) {
	input := `<html>
	<body>
		(#%Message|sample#)
	</body>
</html>
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.StartTagToken, "<html>"},
		{token.StartTagToken, "<body>"},
		{token.TextToken, `
		(#%Message|sample#)
	`},
		{token.EndTagToken, "</body>"},
		{token.EndTagToken, "</html>"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
