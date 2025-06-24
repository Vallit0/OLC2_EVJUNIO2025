grammar V4LangGrammar;

expr
  : expr '+' expr         # AddExpr
  | expr '*' expr         # MulExpr
  | '-' expr              # NegExpr
  | '(' expr ')'          # ParenExpr
  | INT                   # IntLiteral
  | ID                    # VarExpr
  ;

INT: [0-9]+;
ID: [a-zA-Z_][a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;
