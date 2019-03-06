parser grammar tronParser;

options {
  tokenVocab = tronLexer;
}

proto  
    : syntax top_level_statement* EOF
;
syntax
    : SYNTAX EQ_PRE PROTO3 ENDPRE
;
top_level_statement
    : importStatement
    | packageStatement
    | extend
    | optionFile
    | message 
    | enumDefinition 
    | service
    | emptyStatement
;
optionFile
    : OPTION optionName EQ ( constant | pron ) SEMI #Option_File
;
optionMessage
    : OPTION optionName EQ ( constant | pron ) SEMI #Option_Msg
;
optionEnum
    : OPTION optionName EQ ( constant | pron ) SEMI #Option_Enum
;
optionService
    : OPTION optionName EQ ( constant | pron ) SEMI #Option_Service
;
optionRpc
    : OPTION optionName EQ ( constant | pron ) SEMI #Option_Rpc
;
optionName
    : (Ident | LPAREN Ident (DOT Ident)* RPAREN) (DOT Ident)*
;
extend
    : EXTEND Ident (DOT Ident)* LCUR field RCUR
;
importStatement
    : IMPORT (WEAK | PUBLIC)? StrLit SEMI
;
packageStatement
    : PACKAGE Ident (DOT Ident)* SEMI
;
message
    : MESSAGE Ident LCUR messageBody* RCUR
;
messageBody
    : field
    | enumDefinition
    | message
    | optionMessage
    | oneof
    | mapField
    | reserved
    | emptyStatement
;
enumDefinition
    : ENUM Ident LCUR enumBody+ RCUR
;
enumBody
    : optionEnum
    | enumField
    | emptyStatement
;
enumField
    : Ident EQ DASH? Int_lit enumValueOption? SEMI
;
enumValueOption
    : LSB optionName EQ constant (COMMA enumValueOption)* RSB
;
service
    : SERVICE Ident LCUR serviceBody* RCUR
;
serviceBody
    : optionService
    | rpc
    // not defined in the protobuf specification | stream
    | emptyStatement
;
rpc
    : RPC Ident rpcParam RETURNS rpcParam 
      (
        (LCUR (optionRpc | emptyStatement)* RCUR)
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
    : REPEATED? typer Ident EQ Int_lit fieldOptions? SEMI
;
fieldOptions
    : LSB fieldOption (COMMA fieldOption)* RSB
;
fieldOption
    : optionName EQ constant
;
oneof
    : ONEOF Ident LCUR (oneofField | emptyStatement)* RCUR
;
oneofField:
  typer Ident EQ Int_lit fieldOptions? SEMI
;
mapField
    : MAP LCHEVR keyType COMMA typer RCHEVR Ident EQ Int_lit fieldOptions? SEMI
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

messageType: DOT? (Ident DOT)* Ident;

messageOrEnumType: DOT? (Ident DOT)* Ident;

bool_lit: TRUE | FALSE;

emptyStatement: SEMI;

constant
  : Ident (DOT Ident)*
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

