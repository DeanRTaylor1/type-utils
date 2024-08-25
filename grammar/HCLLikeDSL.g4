grammar HCLLikeDSL;

file: importStatement* block*;

importStatement: 'import' STRING;

block: IDENTIFIER '{' (attribute | block)* '}';

attribute:
	IDENTIFIER '=' value
	| 'repeated' IDENTIFIER '=' value;

value: STRING | NUMBER | BOOLEAN | array | IDENTIFIER;

array: '[' value (',' value)* ']';

STRING: '"' (~["\r\n])* '"';
NUMBER: [0-9]+ ('.' [0-9]+)?;
BOOLEAN: 'true' | 'false';
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;
COMMENT: '#' ~[\r\n]* -> skip;
