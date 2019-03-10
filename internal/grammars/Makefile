
tronantlr:
	java -jar ../../.wxb/antlr-4.7.1-complete.jar -Dlanguage=Go -o tron tronLexer.g4 TronParser.g4
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' tron/*.go
	java -jar ../../.wxb/antlr-4.7.1-complete.jar -Dlanguage=Go -o tronwalker -lib ./tron TronWalker.g4 -o walker -package walker
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' walker/*.go
