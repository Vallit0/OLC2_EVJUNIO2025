grammar Vlang; 


// === Axioma principal ===

programa: (stmt)+ EOF;

/*
Desde el codigo vamos a recorrer Stmt y vamos 
a ver de que tipo es cada uno entonces, no podemos 
darle una etiqueta a cada uno 
 */

stmt:
	 decl_stmt                
	 | assign_stmt
     | if_stmt
     | while_stmt
     | for_stmt
     | transfer_stmt           
     | println 
     | print
     ;

if_stmt: if_chain (ELSE_KW if_chain)* else_stmt? # IfStmt;

if_chain: IF_KW expresion  LCOR stmt* RCOR # IfChain;

else_stmt: ELSE_KW LCOR stmt* RCOR # ElseStmt;

println: 'println'  LPAREN expresion RPAREN # PrintlnStmt;
print: 'print'  LPAREN expresion (COMMA expresion)* RPAREN # PrintStmt;

while_stmt: WHILE_KW expresion LCOR stmt* RCOR # WhileStmt;

for_stmt:
	FOR_KW ID IN_KW (expresion | range) LCOR stmt* RCOR # ForStmt;

range: expresion DOT DOT DOT expresion # NumericRange;


transfer_stmt:
	RETURN_KW expresion?	# ReturnStmt
	| BREAK_KW		# BreakStmt
	| CONTINUE_KW	# ContinueStmt;


assign_stmt:
	id_pattern ASSIGN expresion  	# DirectAssign
    ; 

decl_stmt: 
    MUT ID ASSIGN expresion # DeclAssign
    | MUT id_pattern ASSIGN expresion # DeclAssignPattern
    | MUT id_pattern # DeclPattern
    ; 


id_pattern: ID (DOT ID)* # IdPattern;
// === Reglas de expresiones ===
expresion
    : valor                                                #valorexpresion        
    | LPAREN expresion RPAREN                              #parentesisexpre
    | LBRACK expresion RBRACK                              #corchetesexpre
    | op=(NOT | MINUS) expresion                           #unario

// ------------- >  Reglas de Expresiones Aritmeticas
// por facilidad la llamaremos a nuestra regla #BinaryExpr
/*¿Que puede tener una expresion binaria? |

una expresion binaria puede tener dos operandos y un operador.
el operando sera -> expresion 
los operadores -> (*| + | )

Entonces, nos ahorraremos trabajo en el BinaryExpr,
porque lo definiremos en el visit de BinaryExpr
 */
	| left = expresion op = (MUL | DIV | MOD) right = expresion	# BinaryExp// a * b, a / b, a % b
	| left = expresion op = (PLUS | MINUS) right = expresion    # BinaryExp // a + b, a - b
	| left = expresion op = ( LT | LE | GT | GE) right = expresion	# BinaryExp // a < b, a <= b, a > b, a >= b
	| left = expresion op = (EQ | NEQ ) right = expresion	# BinaryExp // a == b, a != b
	| left = expresion op = AND right = expresion		    # BinaryExp // a && b
	| left = expresion op = OR right = expresion 			# BinaryExp // a || b

    | ID                                                   #id              
    | incredecre                                           #incredecr      
    | ID DOT ID                                            #expdotexp1             
    | ID DOT expresion                                     #expdotexp      
    | ID ASSIGN expresion                                  #asignacionfor
    ;

// === Parámetros en llamadas ===
parametros : expresion (COMMA expresion)* ;

// === Tipos de valores simples ===
valores : valor ;

// === Subcontextos para d ===
valor
    : ENTERO    #valorEntero
    | DECIMAL   #valorDecimal
    | CADENA    #valorCadena
    | BOOLEANO  #valorBooleano
    | CARACTER  #valorCaracter
    | FLOAT     #valorFloat
    ;



// === Incremento / Decremento ===
incredecre
    : ID INC    #incremento
    | ID DEC    #decremento
    ;

// === Tokens de palabras clave ===
LEN     : 'len' ;
CAP     : 'cap' ;
APPEND  : 'append' ;
MUT     : 'mut' ;
// === Literales ===
BOOLEANO : 'true' | 'false' ;
ENTERO   : [0-9]+ ;
FLOAT    : [0-9]+ '.' [0-9]* | [0-9]* '.' [0-9]+ ;
DECIMAL  : [0-9]+ '.' [0-9]+ ;
CADENA   : '"' (~["\\] | '\\' .)* '"' ;
CARACTER : '\'' . '\'' ;

// === Identificadores ===
ID : [a-zA-Z_][a-zA-Z0-9_]* ;


IF_KW : 'if' ;
ELSE_KW : 'else' ;
WHILE_KW : 'while' ;
FOR_KW : 'for' ;
IN_KW : 'in' ;
RETURN_KW : 'return' ;
BREAK_KW : 'break' ;
CONTINUE_KW : 'continue' ;
// === Operadores ===
PLUS    : '+' ;
MINUS   : '-' ;
MUL     : '*' ;
DIV     : '/' ;
MOD     : '%' ;
NOT     : '!' ;
OR      : '||' ;
AND     : '&&' ;
EQ      : '==' ;
NEQ     : '!=' ;
LT      : '<' ;
LE      : '<=' ;
GT      : '>' ;
GE      : '>=' ;
ASSIGN  : '=' ;
INC     : '++' ;
DEC     : '--' ;

// === Símbolos ===
LPAREN  : '(' ;
RPAREN  : ')' ;
LBRACK  : '[' ;
LCOR  : '{' ;
RBRACK  : ']' ;
RCOR  : '}' ;
DOT     : '.' ;
COMMA   : ',' ;

// === Espacios y comentarios ===
WS : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;