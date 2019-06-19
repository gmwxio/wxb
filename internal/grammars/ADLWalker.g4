parser grammar ADLWalker;

options {
  tokenVocab = ADLParser;
}

adl
    : DOWN ADL DOWN module UP UP EOF
;
module
    : Annotation* Module Import* tld*
;
tld
    : Annotation* Struct (DOWN TypeParam? nameBody* UP)?                     #PkgNode
    | Annotation* Union  (DOWN TypeParam? nameBody* UP)?                        #ImportNode
    | Annotation* Type   (DOWN TypeParam? TypeExpr? jsonVal* UP)?   #MsgNode
    | Annotation* Newtype   (DOWN TypeParam? TypeExpr? jsonVal* UP)?   #MsgNode
    | ModuleAnno #ExtNode
    | DeclAnno #DeclAnnoNode
    | FieldAnno #FieldAnnoNode
;
nameBody
    : Annotation* Field (DOWN TypeExpr jsonVal UP)?
;
jsonVal
    : Json DOWN JsonStr UP
    | Json DOWN JsonBool UP
    | Json DOWN JsonNull UP
    | Json DOWN JsonInt UP
    | Json DOWN JsonFloat UP
    | Json DOWN JsonArray DOWN jsonVal+ UP UP
    | Json DOWN JsonObj DOWN jsonVal+ UP UP
;