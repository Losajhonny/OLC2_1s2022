grammar Ejercicio2;

@parser::header {
}

@parser::members{
}

// Rules
start
    : EOF
    ;

// Tokens
COMA: ',';
NUMBER: [0-9]+;
ID: [A-Za-z]+;
WHITESPACE: [ \r\n\t]+ -> skip;
