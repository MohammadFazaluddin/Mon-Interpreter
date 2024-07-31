# Mon Interpreter for Monkey Language

Monkey programming language is a high level programming language, Mon (short for Monkey) Interpreter is written in Go (Golang). 
The language uses REPL (Read-Evaluate-Print-Loop) to interpret the syntax, similar to Python or JavaScript. 
REPL can also be called a console.

## Get Started 🚀

Clone the repository

```bash
git clone https://github.com/MohammadFazaluddin/Mon-Interpreter.git
```

Run the following command to start the Monkey Language Console.

```bash
go run main.go
```

That's It, You Have The Monkey and The Mon-Interpreter! 
You can write, execute and enjoy The Monkey. Take it away.  

## How it Works 🤖
The Monkey source code is tokenized and then parsed into a REPL, building up an internal representation of the code called Abstract Syntax Tree
and then evaluates this tree. 

Following are the major parts:
1. The Lexer
2. The Parser
3. The Abstract Syntax Tree (AST)
4. The Internal Object System
5. The Evaluator


1. **The Lexer:** Tokenizes the Monkey Source code to the tokens also known as "Lexical Analysis" or "Lexing".

2. **The Parser:** The Tokens transformed from the lexer is fed to the parser, which does the second transformation and turn the token into "Abstract Syntax Tree".

3. **Evaluation:** Evaluation is the main process while processing the source code. This is where code becomes meaningful. Without evaluation an
expression like 1 + 2 is just a series of characters, tokens, or a tree structure that represents this expression.
