lexer grammar tron_lexer;

LCUR    : '{';
RCUR    : '}';
EQ      : '=' ;
DQ      : '"' ;
SQ      : '\'';
SEMI    : ';';
COLON   : ':';
LPAREN  : '(';
RPAREN  : ')';
DOT     : '.';
LSB     : '[';
RSB     : ']';
COMMA   : ',';
LCHEVR  : '<';
RCHEVR  : '>';
DASH    : '-';
PLUS    : '+';

PROTO3  : '"proto3"' | '\'proto3\'';
SYNTAX  : 'syntax';
EXTEND  : 'extend';
IMPORT  : 'import' ;
WEAK    : 'weak';
PUBLIC  : 'public';
RETURNS : 'returns';
PACKAGE : 'package';
OPTION  : 'option';
ENUM    : 'enum';
SERVICE : 'service';
RPC     : 'rpc';
STREAM  : 'stream';
RESERVED : 'reserved';
TO       : 'to';
REPEATED : 'repeated';
ONEOF    : 'oneof';
MAP      : 'map';
INF      : 'inf';
NAN      : 'nan';
TRUE     : 'true';
FALSE    : 'false';

BOOL    : 'bool';
BYTES   : 'bytes';
DOUBLE  : 'double';
FIXED32 : 'fixed32';
FIXED64 : 'fixed64';
FLOAT   : 'float';
INT32   : 'int32';
INT64   : 'int64';
MESSAGE : 'message';
SFIXED32 : 'sfixed32';
SFIXED64 : 'sfixed64';
SINT32  : 'sint32';
SINT64  : 'sint64';
STRING  : 'string';
UINT32  : 'uint32';
UINT64  : 'uint64';

// Letters and digits
fragment Letter        : [A-Za-z_];
fragment DecimalDigit  : [0-9];
fragment OctalDigit    : [0-7];
fragment HexDigit      : [0-9A-Fa-f];

fragment DecimalLit    : [1-9] DecimalDigit*;
fragment OctalLit      : '0' OctalDigit*;
fragment HexLit        : '0' ('x' | 'X') HexDigit+;

// Floating-point literals
fragment Decimals      : DecimalDigit+;
fragment Exponent      : ('e' | 'E') ('+' | '-')? Decimals;

StrLit: '\'' CharValue* '\'' | '"' CharValue* '"';

Ident
    : Letter (Letter | DecimalDigit)*
;

Int_lit: DecimalLit | OctalLit | HexLit;

Float_lit: (
    Decimals DOT Decimals? Exponent?
    | Decimals Exponent
    | DOT Decimals Exponent?
  )
  | INF
  | NAN
;


fragment CharValue
    : HexEscape
    | OctEscape
    | CharEscape
    | ~[\u0000\n\\]
;

fragment HexEscape  : '\\' ('x' | 'X') HexDigit HexDigit;
fragment OctEscape  : '\\' OctalDigit OctalDigit OctalDigit;
fragment CharEscape : '\\' [abfnrtv\\'"];

WS           : [ \t\r\n\u000C]+ -> skip;
COMMENT      : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;