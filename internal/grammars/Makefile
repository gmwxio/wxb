adlantlr:
	java -jar ../../.wxb/antlr-4.7.1-complete.jar -Dlanguage=Go -o adl adlLexer.g4 ADLParser.g4
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' adl/*.go
	java -jar ../../.wxb/antlr-4.7.1-complete.jar -Dlanguage=Go -o adlwalker -lib ./adl ADLWalker.g4 -o adlwalker -package walker
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' adlwalker/*.go
	java -jar ../../.wxb/antlr-4.7.1-complete.jar -Dlanguage=Go -o dnaca -lib ./adlwalker DNAC_01_Walker.g4 -package dnaca
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' dnaca/*.go

tronantlr:
	java -jar ../../.wxb/antlr-4.7.1-complete.jar -Dlanguage=Go -o tron tronLexer.g4 TronParser.g4
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' tron/*.go
	java -jar ../../.wxb/antlr-4.7.1-complete.jar -Dlanguage=Go -o tronwalker -lib ./tron TronWalker.g4 -o walker -package walker
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' walker/*.go

