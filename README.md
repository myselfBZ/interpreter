# An interpreter for the monkey language


> [!IMPORTANT]
> There are some external dependencies for testing different components

For example:
[to make the AST look pretty in a json doc](https://github.com/tidwall/pretty) 


## Structs have `json` tag, because **I** wanted to inspect tokens and nodes in a json document



# User manual

variable declaration
`let x = 12;`
> [!NOTE]
> Semicolons are required at the end of lines! 

if statements
`if (expression true or false) {
    ...
} else {
    ...
}`

sorry `else if` is not supported

arithmetic expressions
`1 + 3 * 34`

floating point numbers are not supported either.
developer is super sad about that (skill issues)
