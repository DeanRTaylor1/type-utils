grammar HCLLikeDSL;

file: hclconfig? block*;

hclconfig: 'Type_Config' '{' configAttribute* '}';

configAttribute: IDENTIFIER '=' STRING;

// importStatement: 'import' STRING;

block:
	IDENTIFIER '{' blockBody '}'
	| 'repeated' IDENTIFIER '{' blockBody '}'
	| 'optional' IDENTIFIER '{' blockBody '}'
	| 'optional repeated' IDENTIFIER '{' blockBody '}';

blockBody: (attribute | block)*;

attribute:
	IDENTIFIER ('=' value)
	| 'repeated' IDENTIFIER '=' value
	| 'optional' IDENTIFIER '=' value
	| 'optional repeated' IDENTIFIER '=' value;

value: STRING | NUMBER | BOOLEAN | IDENTIFIER;

STRING: '"' (~["\r\n])* '"';
NUMBER: [0-9]+ ('.' [0-9]+)?;
BOOLEAN: 'true' | 'false';
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;
COMMENT: '#' ~[\r\n]* -> skip;
