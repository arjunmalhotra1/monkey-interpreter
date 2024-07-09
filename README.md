# Monkey-Interpreter
The purpose of this project was to learn `How to write an Interpreter in Go`.
This is the part-1 of the book written by Thorsten Ball.

Using Go programming language, I created an interpreter for a new language named `Monkey Programming Language`.


Money supports:
- Integers
- Booleans
- Strings
- Arrays
- Hashes
- Prefix-, infix- & index operators
- Conditionals
- Global & Local bindings
- First-class functions
- Return statements
- Closures



## Lexer
First I built the Lexer that turns strings entered in the REPL into tokens.
The lexer is defined in the lexer package and the tokens it generates can be found in the token package.
Next I built the Parser.

## Parser
Next, I built a top-down recursive descent parser, also knows as `Pratt Parser`.
That turns the tokens into an abstract syntax tree, `AST`.
The output of the Parser - representation of the program as a tree in memory, is given to the Evaluator.

## Evaluator
Evaluator evaluates the program. There's a function called `Eval` defined in the Evaluator package. `Eval` recursively walks down the `AST` & evaluates it using object system defined in the `object` package.

It would turn for example an `AST` node representing 

```1+2``` into 
`object.Integer{Value:3}`


With that the life cycle of Monkey code would complete and the result printed to the REPL.


## Future
### In the next part I will work on turning the tree-walking & on-the-fly-evaluating interpreter into a bytecode compiler and a virtual machine that executes the bytecode.
