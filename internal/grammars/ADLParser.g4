parser grammar ADLParser;

tokens {
    DOWN, UP, ROOT, ERROR,
    Module, Import, Annotation, Struct, Union, Newtype
}

options {
  tokenVocab = adlLexer;
}

adl  
    : module LCUR top_level_statement* RCUR SEMI EOF
;
module
    : kw_mod=ID name+=ID (DOT name+=ID)*                        #ModuleStatement
;
top_level_statement
    : kw_impt=ID a+=ID (DOT a+=ID)* (DOT s=STAR)? SEMI          #ImportStatement
    | kw_soru=ID a=ID typeParam? LCUR soruBody* RCUR SEMI        #StructOrUnion
    | kw_tnew=ID a=ID typeParam? EQ b=ID typeParam? SEMI        #TypeOrNewtype
    | kw_anno=ID a=ID jsonValue SEMI                            #ModuleAnnotation
    | kw_anno=ID a=ID b=ID  jsonValue SEMI                      #DeclAnnotation
    | kw_anno=ID a=ID DCOLON b=ID jsonValue SEMI                #FieldAnnotation
;
typeParam
    : LCHEVR typep+=ID typeParam? (COMMA typep+=ID typeParam?)* RCHEVR
;
soruBody
    : ID typeParam? ID SEMI
;
jsonValue
    : s=STR                                                     #StringStatement
    | kw_tfn=ID                                                 #TrueFalseNull
    | INT                                                       #NumberStatement
    | FLT                                                       #FloatStatement
    | LSQ jsonValue (COMMA jsonValue)* RSQ                      #ArrayStatement
    | LCUR STR COLON jsonValue (COMMA STR COLON jsonValue)* RCUR    #ObjStatement
;