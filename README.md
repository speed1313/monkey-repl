# monkey-repl
this is a toy interpreter for lerning.


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

### expression
parsing expression is difficult because expression has priority.

### Pratt syntax parser
### Points of Parser
-  If an infinite loop occurs, it may be because next() is not called

- the important point is that the parser is able to step forward if the current token or the next token is valid. To check this, expect(), that check if the next token is expected and step forward is used.


## Evaluator
parserから生成されたASTを, 上から再帰的に評価していく. eval(node)は評価した値を返し, 親nodeは子ノードの評価値を用いて新たに値を計算して返す.

### 束縛と環境
letによる変数への値の束縛は, 環境によって実現される. 環境は変数名と値をmapのラッパーで保存するもの.

## 文字列や配列など
goのデータ型をラップする形で表現すればできる.

