grammar Ejercicio1;

/*
S	::=	‘int’ L

L	::=	L [ num ]
	|	[ num ]
*/

@parser::header {
    import "Laboratorio3/Node"
}

@parser::members{
    func mostrar(out string) {
        fmt.Println(out)
    }
}

// Rules
// Método 3
start
    : t='int' na=lista EOF  {
                                mostrar($na.n.Cad + "integer" + $na.n.Aux);
                            }
    ;

lista returns [node.Nodo n]
    @init { $n = node.NewNodo("", ""); }
    :   ('[' num=NUMBER ']' {
                                $n.Cad += "arreglo(" + $num.text + ",";
                                $n.Aux += ")";
                            } )+
    ;

/*
// Metodo 2
start returns [node.Nodo n]
    : 'int' na=lista EOF  {
                                mostrar($na.n.Cad + $na.n.Aux);
                            }
    ;

lista returns [node.Nodo n]
    :   '[' num=NUMBER ']' na=lista     {
                                            cad := "arreglo(" + $num.text + "," + $na.n.Cad;
                                            aux := $na.n.Aux + ")";
                                            $n = node.NewNodo(cad, aux);
                                        }
    |   '[' num=NUMBER ']'              {
                                            $n = node.NewNodo("arreglo(" + $num.text + ", integer", ")");
                                        }
    ;
*/

/* Metodo 1
start returns [node.Nodo n]
    :   t='int' un=lista EOF {
                            mostrar($un.n.Cad + "integer" + $un.n.Aux)
                        }
    ;

lista returns [node.Nodo n]
    :   na=lista '[' num=NUMBER ']' {
                                        cad := $na.n.Cad + "arreglo(" + $num.text + ","
                                        aux := $na.n.Aux + ")"
                                        $n = node.NewNodo(cad, aux)
                                    }
    |   '[' num=NUMBER ']'      {
                                    cad := "arreglo(" + $num.text + ","
                                    aux := ")"
                                    $n = node.NewNodo(cad, aux)
                                }
    ;
*/

// Tokens
CORI: '[';
CORD: ']';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
