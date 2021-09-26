// pkisauce grammar v1
// i'm lazy for fuck sake!
// wanna try? debug?
// install antlr4, then:
// $ antlr4 pksc.g4
// $ javac pksc*.java
// $ grun pksc config -tree
// (DUMP a CONF FILE IN STDIN)
//
// generate the go files for this package:
// antlr4 -Dlanguage=Go -package config pksc.g4
//
// TODO: need to add debug level..
grammar pksc;

WS    : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines

COMMENT      : '/*' .*? '*/'      -> channel(HIDDEN);
LINE_COMMENT : '//' ~[\r\n]*      -> channel(HIDDEN);
TERMINATOR   : [\r\n]+            -> channel(HIDDEN);

LCURLY : '{';
RCURLY : '}';
COMMA  : ',';
QUOTE  : '"';
SEMI   : ';';
KWANY    : 'any';
KWDBG    : 'debug';
KWPASS   : 'pass';
KWON     : 'on';
KWRPC    : 'rpc';
KWDAEMON : 'daemon';
KWPATH   : 'path';
KWUSE    : 'use';
KWGOGRPC : 'go-grpc';
KWGOTLS  : 'go-tls';
KWPEM    : 'pem';

HNAME : QUOTE IDENTIFIER? QUOTE;

IDENTIFIER: LETTER ( LETTER | DECIMAL_DIGIT )*;
fragment LETTER: [A-Za-z_];
fragment DECIMAL_DIGIT: [0-9];

PATH: ( '\'' ( CHAR_VALUE )*? '\'' ) |  ( '"' ( CHAR_VALUE )*? '"' );
fragment CHAR_VALUE: CHAR_ESCAPE | ~[\u0000\n\\];
fragment CHAR_ESCAPE: '\\' ( 'a' | 'b' | 'f' | 'n' | 'r' | 't' | 'v' | '\\' | '\'' | '"' );

gen_type  : KWGOGRPC
          | KWGOTLS
	  | KWPEM ;

name_elt_list : HNAME? (COMMA HNAME)* ;

name_list : LCURLY name_elt_list RCURLY ;

name_block : HNAME
	   | name_list
	   | KWANY ;

pass_stmt : KWPASS name_block KWON name_block KWRPC name_block SEMI  // e.g.: pass { "client1", "client2" } on { "Public", "Private } rpc { "Call1", "Call2" };
	  | KWPASS name_block SEMI  // pass { "client1", "client2" };
	  | KWDBG SEMI ;

pass_block   : LCURLY (pass_stmt)* RCURLY ;

daemon_stmt :
      KWDAEMON HNAME KWPATH PATH KWUSE gen_type pass_block
      | KWDAEMON HNAME KWPATH PATH KWUSE gen_type ;


config : COMMENT*
	| LINE_COMMENT*
	| daemon_stmt+ ;
