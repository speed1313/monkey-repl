# monkey-repl
monkey interpreter


# Note
## Lexer
Lexer reads the source code and returns the token sequense. Lexer is easy to implement because it simply converts words into tokens one by one, ignoring white space.

- attention
Sometimes you have to look at the next character, like == and =.

## Parser
Parser reads token sequenses and construct the AST.
AST is the data structure that represents the source code.
While reading, parser checks if the token sequense is valid.
If the current token or the next token is the same as the expected token, then the parser can step forward.

### Points of Parser
-  If an infinite loop occurs, it may be because next() is not called

- the important point is that the parser is able to step forward if the current token or the next token is valid. To check this, expect(), that check if the next token is expected and step forward is used.
-