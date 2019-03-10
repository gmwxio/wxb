parser grammar TronWalker;

options {
  tokenVocab = TronParser;
}

proto
    : DOWN ROOT syntax? (DOWN tld* UP)? UP EOF
;
syntax
    : SYNTAX
;
tld
    : Package
    | Import
    | Message (DOWN msgBody* UP)?
    | Enum (DOWN enumBody* UP)?
    | Service (DOWN svcBody* UP)?
    | Extend (DOWN extendBody* UP)?
;
msgBody
    : Option
    | Field
    | Message (DOWN msgBody* UP)?
    | Enum (DOWN enumBody* UP)?
;
enumBody
    : Option
    | EnumValue EnumNum?
;
svcBody
    : Option
    | Rpc
;
extendBody
    : Option
    | Field
;