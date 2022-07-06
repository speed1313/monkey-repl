package token

type TokenType string

type Token struct{
    Type TokenType
    Literal string

}


const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF

    // identifier and literal
    IDENT = "IDENT" // variable name, function name,..
    INT = "INT" // integer

    // operator
    ASSIGN = "="
    PLUS = "+"

    // delimiter
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // keyword
    FUNCTION = "FUNCTION"
    LET = "LET"
)

