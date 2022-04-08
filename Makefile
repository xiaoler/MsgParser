parser: 
	antlr -Dlanguage=Go -o parser/ MsgLexer.g4 MsgParser.g4

clean:
	rm -r ./parser

.PHONY: parser parser2 clean
