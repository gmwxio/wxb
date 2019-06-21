parser grammar ADLWalker;
tokens {
    Name, Exnotation
}
options {  tokenVocab = ADLParser; }

adl
    : DOWN ADL DOWN module UP UP EOF
;
module
    : Module (DOWN annotation* Import* tld* UP)?
;
tld
    : Struct  (DOWN annotation* TypeParam? nameBody* UP)?             #Struct
    | Union   (DOWN annotation* TypeParam? nameBody* UP)?             #Union
    | Type    (DOWN annotation* TypeParam? typeExpr_? jsonVal* UP)?   #Type
    | Newtype (DOWN annotation* TypeParam? typeExpr_? jsonVal* UP)?   #Newtype
    | ModuleAnno DOWN jsonVal UP                                      #ModAnno
    | DeclAnno   DOWN jsonVal UP                                      #DeclAnno
    | FieldAnno  DOWN jsonVal UP                                      #FieldAnno
    | (Struct|Union) DOWN annotation* ERROR nameBody* UP              #TypeParamError
    | (Type|Newtype) DOWN annotation* ERROR typeExpr_? jsonVal* UP?   #TypeParamError
;
nameBody
    : Field (DOWN annotation* typeExpr_? jsonVal? UP)?               #Field
;
annotation
    : Annotation (DOWN jsonVal UP)?
;
typeExpr_
    : TypeExpr (DOWN typeExprElem_+ UP)?
;
typeExprElem_
    : TypeExprElem (DOWN typeExprElem_+ UP)?                         #TypeParams
;
jsonVal
    : JsonStr                                          #JsonStr
    | JsonBool                                         #JsonBool
    | JsonNull                                         #JsonNull
    | JsonInt                                          #JsonInt
    | JsonFloat                                        #JsonFloat
    | JsonArray (DOWN jsonVal+ UP)?                    #JsonArray
    | JsonObj (DOWN jsonVal+ UP)?                      #JsonObj
;