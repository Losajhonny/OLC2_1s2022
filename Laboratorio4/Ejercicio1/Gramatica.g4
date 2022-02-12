grammar Gramatica;

options {
    language='Go';
}

start
    : e EOF         # stat
    ;

e
    : t ep          # exp
    ;

ep
    : '+' t ep      # Sum
    |               # EpExp
    ;

t
    : f tp          # term
    ;

tp
    : '*' f tp      # Mul
    |               # EpTerm
    ;

f
    : '(' e ')'     # parentesis
    | NUM           # numero
    ;

// Tokens
MAS: '+';
POR: '*';
PARI: '(';
PARD: ')';
NUM: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
