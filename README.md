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

## Breakdown of Features 👀

1. **The Lexer:** The first transformation from source code to Tokens, called as "Lexical Analysis", or "Lexing" for short. Done by Lexer, also called tokenizer or scanner. Tokens are small, easily categorizable data structures.

2. **The Parser:** The Tokens transformed from the lexer is fed to the parser, which does the second transformation and turn the token into "Abstract Syntax Tree".

3. **Evaluation:** Evaluation is the main process while processing the source code. This is where code becomes meaningful. Without evaluation an
expression like 1 + 2 is just a series of characters, tokens, or a tree structure that represents this expression.

## Supported Features and Syntax of Monkey Language
1. Data Type & Built-in Functions:
- Integers:
```
>> 54
54
>> -3252
-3252
```

- Booleans:
```
>> true == false
false
>> false != true
true
```

- Strings:
```
>> "Hello World"
Hello World
>> len("Monkey")
6
```

- Arrays:
```
>> let a = [1, 2 * 2, 10 - 5, 8 / 2];
>> a[1]
4
```
```
>> let myArray = [true, "Banana", fn(x) { x * x }]
>> myArray[1]
Banana
>> myArray[0]
true
>> myArray[2](3)
9
```
```
>> len(myArray)
3
```

- Hash / Hashmap:
```
>> let bob = {"name": "Bob", "age": 99};
>> bob["name"]
Bob
```
```
>> let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
>> people[0]["name"];
Alice
>> people[1]["age"];
28
>> people[1]["age"] + people[0]["age"];
52
>> let getName = fn(person) { person["name"]; };
>> getName(people[0]);
Alice
>> getName(people[1]);
Anna
```

1. Mathematical Expressions, Variable binding:
```
>> let exp = 3 + 8 * 7 / 2 - 1;
>> exp 
30
```
```
>> (10 + 2) * 30 == 300 + 20 * 3
true
```
```
>> (5 > 5 == true) != false
false
```

1. Conditionals:
```
>> if (5 * 5 + 10 > 34) { 99 } else { 100 }
99
```
```
>> if ((1000 / 2) + 250 * 2 == 1000) { 9999 }
9999
```

1. Functions and application of Functions:
```
>> let addThree = fn(x) { return x + 3 };
>> addThree(3);
6
```

```
>> let max = fn(x, y) { if (x > y) { x } else { y } };
>> max(5, 10)
10
```
```
>> let callTwoTimes = fn(x, func) { func(func(x)) };
>> callTwoTimes(3, addThree);
9
>> callTwoTimes(3, fn(x) { x + 1 });
5
>> let newAdder = fn(x) { fn(n) { x + n } };
>> let addTwo = newAdder(2);
>> addTwo(2);
4
```



