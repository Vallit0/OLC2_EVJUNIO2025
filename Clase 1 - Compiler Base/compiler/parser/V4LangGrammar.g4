parser grammar V4LangGrammar;

options { tokenVocab=V4LangLexer; }

prog : expr EOF ;

expr
    : expr op=(MUL|DIV) expr     # MulDiv
    | expr op=(PLUS|MINUS) expr  # AddSub
    | INT                        # Int
    | LPAREN expr RPAREN         # Parens
    ;
