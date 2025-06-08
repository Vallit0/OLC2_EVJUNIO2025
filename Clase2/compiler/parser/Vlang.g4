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
     | 'print(' expresion ')' 
     ;

assign_stmt:
	id_pattern EQ expresion  	# DirectAssign
    ; 

decl_stmt: 
    var_type ID EQ expresion 
    ; 

var_type: MUT #mutType
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
	| left = expresion op = (MUL | DIV | MOD) right = expresion	# BinaryExp // a * b, a / b, a % b
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
RBRACK  : ']' ;
DOT     : '.' ;
COMMA   : ',' ;

// === Espacios y comentarios ===
WS : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;