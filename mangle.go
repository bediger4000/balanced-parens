package main

import (
	"fmt"
	"os"
)

func main() {
	runes := []rune(os.Args[1])

	i := 0
	for {
		fmt.Printf("%d String: %q\n", i, string(runes))
		if len(runes) < 2 {
			break
		}
		if i == len(runes)-1 {
			break
		}
		switch runes[i] {
		case '(':
			if runes[i+1] == ')' {
				runes = append(runes[:i], runes[i+2:]...)
				i = 0
				continue
			}
		case '[':
			if runes[i+1] == ']' {
				runes = append(runes[:i], runes[i+2:]...)
				i = 0
				continue
			}
		case '{':
			if runes[i+1] == '}' {
				runes = append(runes[:i], runes[i+2:]...)
				i = 0
				continue
			}
		}
		i++
	}

	if len(runes) != 0 {
		fmt.Printf("Expression unbalanced\n")
		return
	}
	fmt.Printf("Expression balanced\n")
}
