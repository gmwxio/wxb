parser grammar tron_parser;

options {
  tokenVocab = tron_lexer;
}

proto  
    : syntax top_level_statement* EOF
;
top_level_statement
    : importStatement
    | packageStatement
    | extend
    | option
    | message 
    | enumDefinition 
    | service
    | emptyStatement
;
syntax
    : SYNTAX EQ PROTO3 SEMI
;
extend
    : EXTEND fullIdent LCUR field RCUR
;
importStatement
    : IMPORT (WEAK | PUBLIC)? StrLit SEMI
;
packageStatement
    : PACKAGE fullIdent SEMI
;
option
    : OPTION optionName EQ ( constant | pron ) SEMI
;
optionName
    : (Ident | LPAREN fullIdent RPAREN) (DOT Ident)*
;
message
    : MESSAGE messageName LCUR messageBody* RCUR
;
messageBody:
    field
    | enumDefinition
    | message
    | option
    | oneof
    | mapField
    | reserved
    | emptyStatement
;
enumDefinition
    : ENUM enumName LCUR enumBody+ RCUR
;
enumBody
    : option 
    | enumField
    | emptyStatement
;
enumField:
  Ident EQ DASH? Int_lit (
    LSB enumValueOption (COMMA enumValueOption)* RSB
  )? SEMI
;
enumValueOption
    : optionName EQ constant
;
service
    : SERVICE serviceName LCUR serviceBody* RCUR
;
serviceBody
    : option
    | rpc
    // not defined in the protobuf specification | stream
    | emptyStatement
;
rpc
    : RPC rpcName rpcParam RETURNS rpcParam 
      (
        (LCUR (option | emptyStatement)* RCUR)
        | SEMI
      )
;
rpcParam
    : LPAREN STREAM? messageType RPAREN
;
reserved
    : RESERVED (ranges | fieldNames) SEMI
;
ranges
    : rangee (COMMA rangee)*
;
rangee
    : Int_lit | Int_lit TO Int_lit
;
fieldNames
    : StrLit (COMMA StrLit)*
;
typer
    : types
    | messageOrEnumType // TODO is this correct??
;
types
    : DOUBLE
    | FLOAT
    | INT32
    | INT64
    | UINT32
    | UINT64
    | SINT32
    | SINT64
    | FIXED32
    | FIXED64
    | SFIXED32
    | SFIXED64
    | BOOL
    | STRING
    | BYTES
;
field
    : REPEATED? typer fieldName EQ Int_lit fieldOptions? SEMI
;
fieldOptions
    : LSB fieldOption (COMMA fieldOption)* RSB
;
fieldOption
    : optionName EQ constant
;
oneof
    : ONEOF oneofName LCUR (oneofField | emptyStatement)* RCUR
;
oneofField:
  typer fieldName EQ Int_lit fieldOptions? SEMI
;
mapField
    : MAP LCHEVR keyType COMMA typer RCHEVR mapName EQ Int_lit fieldOptions? SEMI
;
keyType
    : INT32
    | INT64
    | UINT32
    | UINT64
    | SINT32
    | SINT64
    | FIXED32
    | FIXED64
    | SFIXED32
    | SFIXED64
    | BOOL
    | STRING
;
fullIdent
    : Ident (DOT Ident)*
;
messageName: Ident;

enumName: Ident;

messageOrEnumName: Ident;

fieldName: Ident;

oneofName: Ident;

mapName: Ident;

serviceName: Ident;

rpcName: Ident;

messageType: DOT? (Ident DOT)* messageName;

messageOrEnumType: DOT? (Ident DOT)* messageOrEnumName;

bool_lit: TRUE | FALSE;

emptyStatement: SEMI;

constant:
  fullIdent
  | (DASH | PLUS)? Int_lit
  | (DASH | PLUS)? Float_lit
  | ( StrLit | bool_lit);

pron:
  LCUR id=Ident ':' pronElem RCUR 
;

pronElem:
  StrLit	#PronSTR 
  | pron  #PronOBJ
;