package main

/*
 * Determine whether strings of pairs of opening/closing characters
 * are "balanced". Allows arbitrary pairs, not just (,) [,] {,}
 */

import (
	"flag"
	"fmt"
)

func main() {
	matchPairs := flag.String("m", "(){}[]", "pairs of matching characters")
	flag.Parse()

	mismatch := false
	var stack []rune
	matches := setupMatches(*matchPairs)

	for _, r := range []rune(flag.Arg(0)) {
		if m, ok := matches[r]; len(stack) > 0 && ok {
			top := stack[len(stack)-1]
			if m != top {
				mismatch = true
				break
			}
			// pop matching opening "parentheses"
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, r)
	}

	if mismatch || len(stack) > 0 {
		// unmatched parens left on stack, or mismatched character
		fmt.Printf("Expression unbalanced\n")
		return
	}

	fmt.Printf("Expression balanced\n")
}

func setupMatches(matchPairs string) map[rune]rune {
	matches := make(map[rune]rune)
	var last rune
	for i, r := range []rune(matchPairs) {
		if i&0x01 == 1 {
			matches[r] = last // Find a ']', must find '[' on stack top
			continue
		}
		last = r
	}
	return matches
}
