grammar HCLLikeDSL;

file: hclconfig? block*;

hclconfig: 'HCLCONFIG' '{' configAttribute* '}';

configAttribute: IDENTIFIER '=' STRING;

// importStatement: 'import' STRING;

block:
	IDENTIFIER '{' blockBody '}'
	| 'repeated' IDENTIFIER '{' blockBody '}'
	| 'optional' IDENTIFIER '{' blockBody '}'
	| 'optional repeated' '{' blockBody '}';

blockBody: (attribute | block)*;

attribute:
	IDENTIFIER ('=' value)
	| 'repeated' IDENTIFIER '=' value
	| 'optional' IDENTIFIER '=' value
	| 'optional repeated' '=' value;

value: STRING | NUMBER | BOOLEAN | IDENTIFIER;

STRING: '"' (~["\r\n])* '"';
NUMBER: [0-9]+ ('.' [0-9]+)?;
BOOLEAN: 'true' | 'false';
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;
COMMENT: '#' ~[\r\n]* -> skip;
