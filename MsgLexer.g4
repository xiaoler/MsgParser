lexer grammar MsgLexer;

TEXT: ~'<'+;

SEA_WS: (' ' | '\t' | '\r'? '\n')+;

TAG_OPEN: '<x-' -> pushMode(TAG);

mode TAG;

TAG_TYPE:
	'mention'
	| 'image'
	| 'video'
	| 'file'
	| 'audio'
	// | 'location'
	| 'link'
	| 'sticker'
	| 'contact'
	| 'document'
	| 'meeting';

TAG_CLOSE: '/>' -> popMode;

ATTR_NAME: [a-zA-Z_] [0-9a-zA-Z_]*;

ATTR_VALUE: STRING | INT | BOOL;

EQ: '=';

// whitespace
WS: [ \t\r\n]+ -> skip;

// attr data types
fragment STRING: '"' (ESC | .)*? '"';
fragment ESC: '\\"' | '\\\\';
// 匹配字符\"和\\
fragment INT: '0' | '-'? [1-9] [0-9]*;
fragment BOOL: 'true' | 'false';
