parser grammar tronParser;

options {
  tokenVocab = tronLexer;
}

proto  
    : syntax? top_level_statement* EOF
;
syntax
    : SYNTAX EQ_PRE PROTO3 ENDPRE
;
top_level_statement
    : imp=ID weakORpublic=ID? STR SEMI        #ImportStatement
    | pkg=ID ID (DOT ID)* SEMI                #PackageStatement
    | ext=ID ID (DOT ID)* LCUR field RCUR     #Extend
    | opt=ID optionName EQ ( constant ) SEMI  #OptionDef                                    
    | msg=ID ID LCUR messageBody* RCUR        #Message 
    | en=ID ID LCUR enumBody+ RCUR            #EnumDefinition 
    | svc=ID ID LCUR serviceBody* RCUR        #Service
    | SEMI                                    #EmptyStm
;
messageBody
    : field
    | en=ID ID LCUR enumBody+ RCUR     //EnumDef
    | msg=ID ID LCUR messageBody* RCUR  //Message
    | option
    | oneof
    | mapField
    | reserved
    | emptyStatement
;
enumBody
    : option
    | enumField
    | reserved
    | emptyStatement
;
enumField
    : ID EQ DASH? INT fieldOptions? SEMI
;
serviceBody
    : option
    | rpc
    // not defined in the protobuf specification | stream
    | emptyStatement
;
rpc
    : rpcID=ID ID rpcParam rets=ID rpcParam 
      (
        (LCUR (option | emptyStatement)* RCUR)
        | SEMI
      )
;
rpcParam
    : LPAREN rpcStream? messageType RPAREN
;
rpcStream
    : stream=ID
;
reserved
    : res=ID (ranges | fieldNames) SEMI
;
ranges
    : rangee (COMMA rangee)*
;
rangee
    : INT | INT to=ID (INT | max=ID)
;
fieldNames
    : STR (COMMA STR)*
;
typer
    : DOT? (ID DOT)* ID
;
field
    : fieldRepeat? typer ID EQ INT fieldOptions? SEMI
;
fieldRepeat
    : repeated=ID
;
oneof
    : of=ID ID LCUR (oneofField | emptyStatement)* RCUR
;
oneofField:
  typer ID EQ INT fieldOptions? SEMI
;
mapField
    : mp=ID LCHEVR keyType=ID COMMA typer RCHEVR ID EQ INT fieldOptions? SEMI
;
messageType
    : DOT? (ID DOT)* ID
;
messageOrEnumType
    : DOT? (ID DOT)* ID
;
emptyStatement
    : SEMI
;
option       
    : opt=ID optionName EQ ( constant ) SEMI
;
fieldOptions
    : LSB fieldOption (COMMA fieldOption)* RSB
;
fieldOption
    : optionName EQ ( constant )
;
optionName
        : (ID | LPAREN ID (DOT ID)* RPAREN) (DOT ID)*
;
constant
  : ID (DOT ID)+
  | id=ID     // include bool_lit, nan, inf, max
  | (DASH | PLUS)? INT
  | (DASH | PLUS)? FLT
  | STR
  | constantObj
;
constantObj
    : LCUR constantObjElem ((COMMA|SEMI)? constantObjElem)* RCUR 
;
constantObjElem
    : id=ID COLON constant    #PronSTR 
    | id=ID COLON constantObj        #PronOBJ
;
