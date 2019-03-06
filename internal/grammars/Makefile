
tronantlr:
	java -jar ../../.wxb/antlr-4.7.1-complete.jar -Dlanguage=Go -o tron tronLexer.g4 tronParser.g4
	sed -ie 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/goantlr!' tron/*.go
