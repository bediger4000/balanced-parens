package main

/*
Check if strings of all-parentheses ( '{}', '[]', '()' )
are "balanced".
Sort of an LR(1) parser approach.
*/

import (
	"fmt"
	"os"
)

var matches = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func main() {
	var stack []rune
	runes := []rune(os.Args[1])

	mismatch := false
	for _, r := range runes {
		if m, ok := matches[r]; len(stack) > 0 && ok {
			top := stack[len(stack)-1]
			if m != top {
				fmt.Printf("Got %c, found %c on stack\n", r, top)
				mismatch = true
				break
			}
			// pop matching '(', '[' or '{'
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, r)
	}

	if !mismatch {
		if len(stack) > 0 {
			// unmatched parens left on stack
			fmt.Printf("Expression unbalanced\n")
			return
		}
		fmt.Printf("Expression balanced\n")
	}
}
