parser grammar ADLWalker;
options {  tokenVocab = ADLParser; }

adl
    : DOWN ADL DOWN module UP UP EOF
;
module
    : Module (DOWN Annotation* Import* tld* UP)?
;
tld
    : Struct  (DOWN Annotation* TypeParam? nameBody* UP)?            #Struct
    | Union   (DOWN Annotation* TypeParam? nameBody* UP)?            #Union
    | Type    (DOWN Annotation* TypeParam? TypeExpr? jsonVal* UP)?   #Type
    | Newtype (DOWN Annotation* TypeParam? TypeExpr? jsonVal* UP)?   #Newtype
    | ModuleAnno                                                     #ModAnno
    | DeclAnno                                                       #DeclAnno
    | FieldAnno                                                      #FieldAnno
;
nameBody
    : Field (DOWN Annotation* TypeExpr? jsonVal? UP)?               #Field
;
jsonVal
    : Json DOWN JsonStr UP                                          #JsonStr
    | Json DOWN JsonBool UP                                         #JsonBool
    | Json DOWN JsonNull UP                                         #JsonNull
    | Json DOWN JsonInt UP                                          #JsonInt
    | Json DOWN JsonFloat UP                                        #JsonFloat
    | Json DOWN JsonArray (DOWN jsonVal+ UP)? UP                    #JsonArray
    | Json DOWN JsonObj (DOWN jsonVal+ UP)? UP                      #JsonObj
;