parser grammar ADLParser;

tokens {
    DOWN, UP, ROOT, ERROR,
    Module, Import, Annotation, Struct, Union, Newtype
}

options {
  tokenVocab = adlLexer;
}

adl  
    : module LCUR imports* top_level_statement* RCUR SEMI EOF
;
module
    : annon* kw_mod=ID name+=ID (DOT name+=ID)*                         #ModuleStatement
;
imports
    : kw_impt=ID a+=ID (DOT a+=ID)* (DOT s=STAR)? SEMI                  #ImportStatement
;
annon
    : AT ID jsonValue                                                   #LocalAnno
    | LINE_DOC                                                          #DocAnno
;
top_level_statement
    : annon* kw_soru=ID a=ID typeParam? LCUR soruBody* RCUR SEMI                #StructOrUnion
    | annon* kw_tnew=ID a=ID typeParam? EQ b=ID typeExpr? (EQ jsonValue)? SEMI  #TypeOrNewtype
    | kw_anno=ID a=ID jsonValue SEMI                                            #ModuleAnnotation
    | kw_anno=ID a=ID b=ID  jsonValue SEMI                                      #DeclAnnotation
    | kw_anno=ID a=ID DCOLON b=ID c=ID jsonValue SEMI                           #FieldAnnotation
;
typeParam
    : LCHEVR typep+=ID (COMMA typep+=ID)* RCHEVR                                #TypeParameter
;
typeExpr
    : LCHEVR typep+=typeExprElem (COMMA typep+=typeExprElem)* RCHEVR            #TypeExpression
;
typeExprElem
    : ID typeExpr?                                                      #TypeExpressionElem
;
soruBody
    : annon* a=ID typeExpr? b=ID (EQ jsonValue)? SEMI                   #FieldStatement
;
jsonValue
    : s=STR                                                             #StringStatement
    | kw_tfn=ID                                                         #TrueFalseNull
    | INT                                                               #NumberStatement
    | FLT                                                               #FloatStatement
    | LSQ (jsonValue (COMMA jsonValue)*)? RSQ                           #ArrayStatement
    | LCUR (STR COLON jsonValue (COMMA STR COLON jsonValue)*)? RCUR     #ObjStatement
;
