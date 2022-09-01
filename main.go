package main

import (
	"fmt"

	"github.com/ytakaya/nasparse/lexer"
	"github.com/ytakaya/nasparse/token"
)

func main() {
	input := `<html>
	<body>
		(#%Message|sample1#)
		(#%Message|sample2#)
		(#%Message|sample3#)
	</body>
</html>
`

	l := lexer.New(input)
	for {
		tok := l.NextToken()
		if tok.Type == token.NaspToken {
			fmt.Printf("found nasp tag in line %d: %s\n", tok.LinePosition, tok.Literal)
		}
		if tok.Type == token.ErrorToken {
			break
		}
	}
}
