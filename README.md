# Writing An Interpreter In Go
This repo documents my attempt at following the book [Writing An Interpreter In Go](https://interpreterbook.com/).

## The Language - Monkey

> In this book we'll create an interpreter for the Monkey programming language. Monkey is a language especially designed for this book. We will bring it to life by implementing its interpreter.
> 
> Monkey looks like this:
> ```
> // Bind values to names with let-statements
> let version = 1;
> let name = "Monkey programming language";
> let myArray = [1, 2, 3, 4, 5];
> let coolBooleanLiteral = true;
> 
> // Use expressions to produce values
> let awesomeValue = (10 / 2) * 5 + 30;
> let arrayWithValues = [1 + 1, 2 * 2, 3];
> ```
> 
> Monkey also supports function literals and we can use them to bind a function to a name:
>
> ```
> // Define a `fibonacci` function
> let fibonacci = fn(x) {
>   if (x == 0) {
>     0                // Monkey supports implicit returning of values
>   } else {
>     if (x == 1) {
>       return 1;      // ... and explicit return statements
>     } else {
>       fibonacci(x - 1) + fibonacci(x - 2); // Recursion! Yay!
>     }
>   }
> };
> ```
>
> The data types we're going to support in this book are booleans, strings, hashes, integers and arrays. We can combine them!
>
> ```
> // Here is an array containing two hashes, that use strings as keys and integers
> // and strings as values
> let people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];
> 
> // Getting elements out of the data types is also supported.
> // Here is how we can access array elements by using index expressions:
> fibonacci(myArray[4]);
> // => 5
> 
> // We can also access hash elements with index expressions:
> let getName = fn(person) { person["name"]; };
> 
> // And here we access array elements and call a function with the element as
> // argument:
> getName(people[0]); // => "Anna"
> getName(people[1]); // => "Bob"
> ```
> 
> That's not all though. Monkey has a few tricks up its sleeve. In Monkey functions are first-class citizens, they are treated like any other value. Thus we can use higher-order functions and pass functions around as values:
>
> ```
> // Define the higher-order function `map`, that calls the given function `f`
> // on each element in `arr` and returns an array of the produced values.
> let map = fn(arr, f) {
>   let iter = fn(arr, accumulated) {
>     if (len(arr) == 0) {
>       accumulated
>     } else {
>       iter(rest(arr), push(accumulated, f(first(arr))));
>     }
>   };
> 
>   iter(arr, []);
> };
> 
> // Now let's take the `people` array and the `getName` function from above and
> // use them with `map`.
> map(people, getName); // => ["Anna", "Bob"]
> ```
>
> And, of course, Monkey also supports closures:
>
> ```
> // newGreeter returns a new function, that greets a `name` with the given
> // `greeting`.
> let newGreeter = fn(greeting) {
>   // `puts` is a built-in function we add to the interpreter
>   return fn(name) { puts(greeting + " " + name); }
> };
> 
> // `hello` is a greeter function that says "Hello"
> let hello = newGreeter("Hello");
> 
> // Calling it outputs the greeting:
> hello("dear, future Reader!"); // => Hello dear, future Reader!
> ```
>
> So, to summarize: Monkey has a C-like syntax, supports variable bindings, prefix and infix operators, has first-class and higher-order functions, can handle closures with ease and has integers, booleans, arrays and hashes built-in.
> 


