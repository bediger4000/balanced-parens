# Balanced parentheses


This problem was asked by Facebook.

Given a string of round, curly, and square open and closing brackets,
return whether the brackets are balanced (well-formed).

For example, given the string "([])[]({})", you should return true.

Given the string "([)]" or "((()", you should return false.

## Building and running

```sh
$ go build balanced.go
$ ./balanced '(([]{}()))'
Expression balanced
$ go build mangle.go
$ ./mangle '(([]{}()))'
Expression balanced
$ go build balanced2
$ ./balanced2 -m '<>()[]{}' '<{[]}()>'
Expression balanced
```

The `balanced2` program allows specification of which pairs of characters
are corresponding opening and closing "parentheses".
It allows you to specify ')' as opening paren,
and '(' as the corresponding closing paren, if you want to.
It defaults to `(){}[]`, 3 pairs of left and matching right parentheses,
but you can use any characters, like `./balanced2 -m aAbBcC abBcBA`.
You can investigate Dyck-$X language recognizing.

## Computer Science

The programs this problem prompts you to write are recognizers for
[Dyck Languages](https://planetmath.org/dycklanguage),
specifically, Dyck-3.

## Analysis

This problem has been around for a while,
I don't think it's original to Facebook.

This also got sent to me as "Daily Coding Problem: Problem #712 [Easy]",
and Daily Coding Problem: Problem #809 [Easy] 

It appears in the "Daily Coding Problem" book
in chapter 4, "Stacks and Queues",
as "4.2 Determine whether brackets are balanced".
This is one problem that the authors give some hints at
further exploration, in that they note some "balanced"
strings of parens are not palindromes, but some "unbalanced"
strings of parens are palindromes.
Usually, that book's emphasis is on getting a solution ASAP,
without giving any context or ideas for futher exploration.

You can't do the "+1 for a left paren, -1 for a right"
stunt that you can for [balanced parentheses problems](https://github.com/bediger4000/binary-tree-odd-string-rep#iteration-3).
Even if you kept 3 counts, one each for parens, square brackets and braces,
it would pass intermingled pairs, '([){]}' would pass.

The algorithm I used is to push left-{parens,brackets,braces} on a stack
as you encounter them working through the string left-to-right.
If/when you encounter a right-{paren,bracket,brace},
a matching left-hand must be on the top of the stack.
Since you've made a balanced pair, pop the top item off the stack,
and examine the next item in the string.

I think this is a simple [LR(1) parser](https://en.wikipedia.org/wiki/LR_parser).

The grammar has 6 tokens '(', ')', '{', '}', '[', ']'.
The grammar has a few productions:

1. start &rarr; PAIRLIST
2. PAIRLIST &rarr; &epsilon; | PAIR | PAIRLIST PAIR
3. PAIR &rarr; '(' PAIRLIST ')' | '[' PAIRLIST ']' | '{' PAIRLIST '}'

The "tokenizer" is just an index into the input string,
since each rune (this is Go) is a token.
The action on encountering a left-{paren,bracket,brace} is to shift the
token onto the stack.
The action on encountering a right-{paren,bracket,brace} is to reduce
the right token with matching left token into a PAIR.
I will grant that some of the other actions (PAIRLIST &rarr; PAIRLIST PAIR)
are implicit in popping the top token off the stack.
The first PAIR gets implicitly turned into a PAIRLIST,
and each successive PAIR that gets popped make the PAIRLIST one item longer.

There's another algorithm that involves indexing through the runes
of the string, eliminating any simple substrings that match '()', '[]' or '{}'.
After each elimination, you start over at index 1 or 0, as appropriate.
If you end up with a zero-length string,
the original string has balanced parens/brackets/braces.
I wrote a [program](mangle.go) for this algorithm to see how it works.

The interviewer could look at whether the candidate tries test cases.
My original cut at this missed the case of an unmatched '(', '{' or '[' case,
despite my best efforts at testing.

I have mixed feelings about this as an interview question.
I did it in less than 50 lines of Go,
but there's a stack and array in it,
so it hits a few data structures,
even if they look like arrays in my code.
There's a lot of comp sci packed into a small problem.

But it seems like if you haven't seen this before,
you'll be baffled.
At best, you'll end up writing/blackboarding code that splices
out substrings of matching pairs.
Maybe, just maybe, if you're a compiler writer,
you'll write a grammar,
and try to implement a parser for the grammar.

In some cases, the interviewer will get a regurgitation of code
that the candidate wrote in that past.
If I were a candidate that got this problem,
that's what you'd get from me.
In others, the interviewer will see the candidate flail for an algorithm,
and maybe implement a hacky substring-match-and-eliminate.
I think that an interviewer wouldn't see much in the way of programming,
so it's not a suitable question if the interviewer wants to see someone
write a program.

The "easy" rating seems off.
This has to qualify as a "medium" if you do actual matching,
and maybe even if you just cut out matching-pairs-substrings.
