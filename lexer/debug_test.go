package lexer

import (
	"fmt"
	"go-monkey/token"
	"testing"
)

func TestDebugTokens(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
  return true;
} else {
  return false;
}
10 == 10;
10 != 9;
`

	l := New(input)

	for i := 0; i < 120; i++ {
		tok := l.NextToken()
		fmt.Printf("%d: type=%q literal=%q\n", i, tok.Type, tok.Literal)
		if tok.Type == token.EOF {
			break
		}
	}
}
