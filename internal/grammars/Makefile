# ANTLR=../../.wxb/antlr4-4.7.1-SNAPSHOT-complete.jar
# ANTLR=/home/garym/devel/wxio/antlr4/tool/target/antlr4-4.7.1-SNAPSHOT-complete.jar
ANTLR=../../.wxb/antlr4-4.7.1-SNAPSHOT-complete.jar


adlantlr:
	java -jar $(ANTLR) -Dlanguage=Go -visitor -o adl adlLexer.g4 ADLParser.g4
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' adl/*.go
	java -jar $(ANTLR) -Dlanguage=Go -visitor -o adlwalker -lib ./adl ADLWalker.g4 -package walker
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' adlwalker/*.go
	java -jar $(ANTLR) -Dlanguage=Go -visitor -o dnaca -lib ./adlwalker DNAC_A_Walker.g4 -package dnaca
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' dnaca/*.go
	java -jar $(ANTLR) -Dlanguage=Go -visitor -o dnacb -lib ./dnaca DNAC_B_Walker.g4 -package dnacb
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' dnacb/*.go

tronantlr:
	java -jar $(ANTLR) -Dlanguage=Go -o tron tronLexer.g4 TronParser.g4
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' tron/*.go
	java -jar $(ANTLR) -Dlanguage=Go -o tronwalker -lib ./tron TronWalker.g4 -o walker -package walker
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' walker/*.go

