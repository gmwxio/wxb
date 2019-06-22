parser grammar DNAC_B_Walker;
options {  tokenVocab = DNAC_A_Walker; }

dnac
    : DOWN DNAC DOWN tld UP UP EOF
;
tld
    : Name  (DOWN Annotation* TypeParam? tld* UP)?                  #NameNode
    | Type  (DOWN Annotation* TypeParam? TypeExpr? Default? UP)?    #TypeNode
    | Field (DOWN Annotation* TypeExpr? Default? UP)?               #NameBodyNode
    | Exnotation                                                    #ExnotationNode
;
