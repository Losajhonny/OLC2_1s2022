grammar Gramatica;

/* Opciones, header y members */
options {
    language='Go';
}

@parser::header {
    import "Compilador/gen"
    import "Compilador/entorno"
}

@parser::members{
    // posición relativa de un símbolo
    var desp int = 0

    // entorno actual
    var tope *entorno.Entorno
}

/* Reglas */
start
    :   instrucciones EOF
    ;

/* marcador para soltar etiquetas */
marcador [ []string lista ]
    :   { gen.Soltar(lista) }
    ;

/* expresiones aritmeticas, relacionales y booleanas en conjunto */
expresion returns[string dir, []string lv, []string lf, string cad]
    :   <assoc=right> op='+' e=expresion    {
                                                $dir = gen.NewTemp()
                                                gen.GenOperacionUnaria($dir, $op.text, $e.dir)
                                            }
    |   <assoc=right> op='-' e=expresion    {
                                                $dir = gen.NewTemp()
                                                gen.GenOperacionUnaria($dir, $op.text, $e.dir)
                                            }
    |   <assoc=right> op='!' e=expresion    {
                                                $lv = $e.lf
                                                $lf = $e.lv
                                            }
    |   e1=expresion op=('*'|'/'|'%') e2=expresion  {
                                                        $dir = gen.NewTemp()
                                                        gen.GenOperacion($dir, $e1.dir, $op.text, $e2.dir)
                                                    }
    |   e1=expresion op=('+'|'-') e2=expresion      {
                                                        $dir = gen.NewTemp()
                                                        gen.GenOperacion($dir, $e1.dir, $op.text, $e2.dir)
                                                    }
    |   e1=expresion opr=oprel e2=expresion         {
                                                        $lv = append($lv, gen.NewEti())
                                                        $lf = append($lf, gen.NewEti())
                                                        $cad = gen.GenIf($e1.dir, $opr.op, $e2.dir, $lv[0])
                                                        gen.GenGoto($lf[0])
                                                    }
    |   e1=expresion op='&&' marcador[$e1.lv] e2=expresion      {
                                                                    $lv = $e2.lv
                                                                    $lf = gen.Unir($e1.lf, $e2.lf)
                                                                }
    |   e1=expresion op='^'  marcador[$e1.lf] e2=expresion      {
                                                                    gen.Soltar($e1.lv)
                                                                    gen.GenIfCad($e2.cad, $e2.lf[0])
                                                                    gen.GenGoto("goto " + $e2.lv[0])
                                                                    $lv = $e2.lv
                                                                    $lf = $e2.lf
                                                                }
    |   e1=expresion op='||' marcador[$e1.lf] e2=expresion      {
                                                                    $lf = $e2.lf
                                                                    $lv = gen.Unir($e1.lv, $e2.lv)
                                                                }
    |   '(' e=expresion ')'     {
                                    $dir = $e.dir
                                    $lv = $e.lv
                                    $lf = $e.lf
                                    $cad = $e.cad
                                }
    |   n=NUM                   {
                                    $dir = $n.text
                                    $cad = ""
                                }
    |   id=ID                   {
                                    $dir = $id.text
                                    $cad = ""
                                }
    |   'true'                  {
                                    $dir = ""
                                    $cad = ""
                                    $lv = append($lv, gen.NewEti())
                                    $lf = append($lf, gen.NewEti())
                                    gen.GenGoto($lv[0])
                                }
    |   'false'                 {
                                    $dir = ""
                                    $cad = ""
                                    $lv = append($lv, gen.NewEti())
                                    $lf = append($lf, gen.NewEti())
                                    gen.GenGoto($lf[0])
                                }
    ;

oprel returns[ string op ]
    :   ope='=='    { $op = $ope.text }
    |   ope='!='    { $op = $ope.text }
    |   ope='>'     { $op = $ope.text }
    |   ope='<'     { $op = $ope.text }
    |   ope='>='    { $op = $ope.text }
    |   ope='<='    { $op = $ope.text }
    ;

/* instrucciones */
instrucciones
    :   instruccion*
    ;

instruccion
    :   inst_asignacion
    |   inst_if
    ;

/* instruccion de asignacion */
inst_asignacion
    :   id=ID '=' e=expresion ';' { gen.GenAsignacion($id.text, $e.dir) }
    ;

/* instrucciones de seleccion */
inst_if locals[ string lsalida ]
    :   'if' e1=expresion   {
                                $lsalida = gen.NewEti()
                                gen.Soltar($e1.lv)
                            }
        bloque  {
                    gen.GenGoto($lsalida)
                    gen.Soltar($e1.lf)
                }
        (
            'else' 'if' e2=expresion    {
                                            gen.Soltar($e2.lv)
                                        }
            bloque  {
                        gen.GenGoto($lsalida)
                        gen.Soltar($e2.lf)
                    }
        )*
        ('else' bloque)? {   gen.GenDestino($lsalida) }
    ;

/* instrucciones ciclicas */

/* bloque de instrucciones */
bloque
    :   '{' instrucciones '}'
    ;

/* Tokens */
NUM: [0-9]+;
ID: [_A-Za-z][_A-Za-z0-9]*;
COMMENT: '/*' .*? '*/' -> skip;
WHITESPACE: [ \r\n\t]+ -> skip;
