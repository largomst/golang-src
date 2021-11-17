/** Grammar from tour chapter augmented with actions */
grammar Expr;
prog:
	stat+;
stat:
	expr NEWLINE
	| ID '=' expr NEWLINE
	| NEWLINE;
expr:
	expr op = ('*' | '/') expr		# MulDiv
	| expr op = ('+' | '-') expr	# AddSub
	| INT							# Number
	| ID							# ID
	| '(' expr ')'					# BRACKET;

NEWLINE:
	'\r'? '\n';
INT:
	[0-9]+;
ID:
	[a-zA-Z]+;
WS:
	[ \t\r\n]+ -> skip;
MUL:
	'*';
DIV:
	'/';
ADD:
	'+';
SUB:
	'-';