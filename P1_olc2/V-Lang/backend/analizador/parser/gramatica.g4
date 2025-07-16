grammar gramatica;

    /* SINTACTICO */
init: instrucciones*;

instrucciones: instruccion+;

instruccion
    : print
    | println
    | declaracionMultiple
    | declaracionMultipleSimple
    | declaracionMultipleSinTipo
    | asignacionMultiple
    | structDef
    | structInst
    | accesoStruct
    | funStruct
    | actStruct
    | asigStruct
    | fnSinParametro
    | fnConParametro
    | llamadaFuncionesSinParametro
    | llamadaFuncionesConParametro
    | asigIncre
    | asigDecre
    | asignacion    
    | sif           
    | sfor          
    | sSwitch       
    | break_    
    | continue_
    | retorno
    | sliceDef    
    | declaracionSliceVacio
    | modificacionElementoSlice
;

print: PRINT PAR1 listaExpr PAR2;
println: PRINTLN PAR1 listaExpr PAR2;

declaracionMultiple
    : MUT listaIDS tipos IGUAL listaExpr
    ;

declaracionMultipleSimple
    : MUT listaIDS tipos
    ;

declaracionMultipleSinTipo
    : MUT listaIDS IGUAL listaExpr
    ;

asignacionMultiple
    : listaIDS IGUAL listaExpr
    ;

// structs
structDef : STRUCT IDENTIFICADOR LLAV1 atributos+ LLAV2
;

atributos : tipos IDENTIFICADOR PTCOMA
;

structInst : IDENTIFICADOR IGUAL IDENTIFICADOR structInit
;

structInit : LLAV1 listaStructs LLAV2
;

listaStructs : expresion (COMA expresion)*
;

accesoStruct : IDENTIFICADOR '.' IDENTIFICADOR
;

asigStruct : IDENTIFICADOR '.' IDENTIFICADOR IGUAL expresion
;

funStruct : FUNCIONES PAR1 IDENTIFICADOR '*' IDENTIFICADOR PAR2 IDENTIFICADOR PAR1 listaPar PAR2 caja
;

caja : LLAV1 valores+ LLAV2
;

valores : IDENTIFICADOR '.' IDENTIFICADOR IGUAL IDENTIFICADOR
;

actStruct : IDENTIFICADOR '.' IDENTIFICADOR PAR1 listaStructs PAR2
;

// funciones

bloqueFuncion
    : (instruccion | expresion)*
    ;

llamadaFuncionesSinParametro
    : IDENTIFICADOR '('')'
    ;

llamadaFuncionesConParametro
    : IDENTIFICADOR '(' listaExpr ')'
    ;

fnSinParametro
    : FUNCIONES IDENTIFICADOR '('')' tipoRetorno? '{' bloqueFuncion '}'
    ;

fnConParametro
    : FUNCIONES IDENTIFICADOR '(' listaPar ')' tipoRetorno? '{' bloqueFuncion'}'
    ;

// slices

sliceDef
    : IDENTIFICADOR IGUAL sliceLiteral
    ;

sliceLiteral
    : '[' ']' tipos '{' listaExpr '}'
    | '[' ']' '[' ']' tipos '{' listaExprList '}'
    ;

accesoElementoSlice
    : IDENTIFICADOR '['expresion']'
    | IDENTIFICADOR '['expresion']''['expresion']'
    ;

modificacionElementoSlice
    : IDENTIFICADOR '['expresion']' IGUAL expresion
    | IDENTIFICADOR '['expresion']''['expresion']' IGUAL expresion
    ;

fnIndexOf
    : INDEXOF '('listaExpr')'
    ;
fnJoin
    : JOIN '('listaExpr')'
    ;
fnLen
    : LEN '('listaExpr')'
    ;
fnAppend
    : APPEND '('listaExpr')'
    ;


declaracionSliceVacio
    : MUT IDENTIFICADOR '['']' tipos
    | MUT IDENTIFICADOR '['']''['']' tipos
    ;


tipoRetorno
    : INT
    | FLOAT
    | BOOL
    | STRING
    | RUNE
    | IDENTIFICADOR
    ;

retorno
    : 'return' expresion ';'?
    | 'return' ';'?
    ;

fnTypeOf
    : TYPEOF '(' listaExpr ')'
    ;

fnAtoi
    : ATOI '(' listaExpr ')'
    ;

fnParseToFloat
    : 'parseFloat' '(' listaExpr ')'
    ;

asigIncre : IDENTIFICADOR '+=' expresion
;

asigDecre : IDENTIFICADOR '-=' expresion
;

asignacion
    : IDENTIFICADOR IGUAL expresion  # AsignacionNormal
    | IDENTIFICADOR FINCREMENTO      # AsignacionIncremento
    | IDENTIFICADOR FDECREMENTO      # AsignacionDecremento
    ;


listaIDS
    : IDENTIFICADOR (COMA IDENTIFICADOR)*
    ;

listaPar 
    : parametro (COMA parametro)* 
    ;

parametro
    : IDENTIFICADOR tipos
    ;

listaExpr
    : expresion (COMA expresion)*
    ;
    
listaExprList
    : '{' listaExpr '}' (COMA '{' listaExpr '}')* COMA?
    ;

expresion: MENOS expresion
         | DIFER expresion
         | expresion MODULO expresion
         | expresion DIV expresion
         | expresion POR expresion
         | expresion MENOS expresion
         | expresion MAS expresion
         | expresion DIFERENTE expresion
         | expresion IGUALDAD expresion
         | expresion MENIGUAL expresion
         | expresion MAYIGUAL expresion
         | expresion MENOR expresion
         | expresion MAYOR expresion
         | expresion OR expresion
         | expresion AND expresion
         | ENTERO
         | DECIMAL
         | CADENA
         | RUNE
         | TRUE
         | FALSE
         | accesoElementoSlice
         | '[' listaExpr? ']'
         | llamadaFuncionesSinParametro
         | llamadaFuncionesConParametro
         | IDENTIFICADOR
         | PAR1 expresion PAR2
         | fnAtoi
         | fnParseToFloat
         | fnTypeOf
         | accesoStruct
         | fnAppend
         | fnIndexOf
         | fnJoin
         | fnLen
;

/* IF */
sif
    : IF (PAR1 expresion PAR2 | expresion) bloque (elseifPart)* (ELSE bloque)?
;

elseifPart
    : ELSE IF (PAR1 expresion PAR2 | expresion ) bloque
;

bloque
    : LLAV1 instrucciones* LLAV2
;

/*  FOor Condicional, Clasico y Range */
sfor
    : FOR expresion bloque                                              #ForCondicional
    | FOR asignacion PTCOMA expresion PTCOMA asignacion bloque          #ForClasico
    | FOR IDENTIFICADOR COMA IDENTIFICADOR IGUAL RANGE expresion bloque #ForRange
    ;


/* SWITCH */
sSwitch
    : SWITCH expresion LLAV1 caseBlock* defaultBlock? LLAV2
    ;
break_: BREAK;
continue_: CONTINUE;
caseBlock
    : CASE expresion DOSPTS instrucciones
    ;

defaultBlock
    : DEFAULT DOSPTS instrucciones
    ;

tipos : INT
      | FLOAT
      | STRING
      | BOOL
      | IDENTIFICADOR
;

    /* LEXICO */
// palabras reservadas
PRINT : 'print';
PRINTLN : 'println';
FUNCIONES : 'fn';
MUT : 'mut';
SLICE : 'slice';
IF : 'if';
ELSE : 'else';
SWITCH : 'switch';
CASE : 'case';
DEFAULT : 'default';
FOR : 'for';
BREAK : 'break';
CONTINUE : 'continue';
RETURN : 'return';
INT : 'int';
FLOAT : 'float64';
STRING : 'string';
BOOL : 'bool';
NIL : 'nil';
TRUE : 'true';
FALSE : 'false';
INDEXOF : 'indexOf';
JOIN : 'join';
LEN : 'len';
APPEND : 'append';
STRUCT : 'struct';
FUNC : 'func';
RANGE : 'range'; 
ATOI : 'Atoi';
PARSEFLOAT : 'parseFloat';
TYPEOF : 'typeOf';
// signos
MAS : '+';
MENOS : '-';
POR : '*';
DIV : '/';
MODULO : '%';
PAR1 : '(';
PAR2 : ')';
LLAV1: '{';
LLAV2: '}';
COR1 : '[';
COR2 : ']';
DIFERENTE : '!=';
IGUALDAD : '==';
MENIGUAL : '<=';
MAYIGUAL : '>=';
IGUAL : '=' | ':=';
MENOR : '<';
MAYOR : '>';    
OR : '||';
AND : '&&';
DIFER : '!';
DOSPTS : ':';
PTCOMA : ';';
COMA : ',';
FINCREMENTO : '++';
FDECREMENTO : '--'; 
// patrones
IDENTIFICADOR : [a-zA-Z_][a-zA-Z0-9_]*;
ENTERO : [0-9]+;
DECIMAL : [0-9]+ '.' [0-9]+;
CADENA : '"' ( ESC_SEQ | ~["\\\r\n] )* '"' ;

fragment ESC_SEQ : '\\' [nrt"\\'];
RUNE : '\'' .? '\'';

COMENTARIO : '//' ~[\r\n]* -> skip;
MULTICOMENTARIO : '/*' .*? '*/' -> skip;
WS : [ \t\r\n]+ -> skip;