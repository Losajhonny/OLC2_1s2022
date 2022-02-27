grammar Gramatica;

/* Opciones, header y members */
options {
    language='Go';
}

@parser::header {
    import "Laboratorio5/gen"
    import "Laboratorio5/entorno"
}

@parser::members{
    // posición relativa de un símbolo
    var desp int = 0

    // entorno actual
    var tope *entorno.Entorno

    // temporal
    var tmp int = 1
}

/* Reglas */
start
    :   instrucciones EOF
    ;

/* expresiones sin reutilizar temporales */
exp returns[ string dir ]
    :   e1=exp op=('*'|'/') e2=exp  {
                                        $dir = gen.NewTemp()
                                        gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir)
                                    }
    |   e1=exp op=('+'|'-') e2=exp  {
                                        $dir = gen.NewTemp()
                                        gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir)
                                    }
    |   '(' e=exp ')'               {
                                        $dir = $e.dir
                                    }
    |   n=NUM                       {   $dir = $n.text  }
    |   id=ID                       {   $dir = $id.text }
    ;

/* expresiones con reutilizacion de temporales */
exp2 returns[ string dir, int num ]
    :   e1=exp2 op=('*'|'/') e2=exp2    {
                                            tmp = tmp - $e1.num - $e2.num
                                            $num = 1
                                            $dir = "t" + strconv.Itoa(tmp)
                                            gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir)
                                            tmp = tmp + 1
                                        }
    |   e1=exp2 op=('+'|'-') e2=exp2    {
                                            tmp = tmp - $e1.num - $e2.num
                                            $num = 1
                                            $dir = "t" + strconv.Itoa(tmp)
                                            gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir)
                                            tmp = tmp + 1
                                        }
    |   '(' e=exp2 ')'                  {
                                            $num = $e.num
                                            $dir = $e.dir
                                        }
    |   n=NUM                           {
                                            $num = 0
                                            $dir = $n.text
                                        }
    |   id=ID                           {
                                            $num = 0
                                            $dir = $id.text
                                        }
    ;

/* condiciones */
cond returns[ []string lv, []string lf, string cad ]
    :   <assoc=right> op='!' c=cond {
                                        $lv = $c.lf
                                        $lf = $c.lv
                                    }
    |   e1=exp opr=oprel e2=exp     {
                                        $lv = append($lv, gen.NewEti())
                                        $lf = append($lf, gen.NewEti())
                                        $cad = $e1.dir + " " + $opr.op + " " + $e2.dir
                                        gen.Gen("if " + $cad + " then goto " + $lv[0])
                                        gen.Gen("goto " + $lf[0])
                                    }
    |   c1=cond op='&&' marcador[$c1.lv] c2=cond    {
                                                        $lv = $c2.lv
                                                        $lf = gen.Unir($c1.lf, $c2.lf)
                                                    }
    |   c1=cond op='^' marcador[$c1.lf] c2=cond     {
                                                        gen.Soltar($c1.lv)
                                                        gen.Gen("if " + $c2.cad + " then goto " + $c2.lf[0])
                                                        gen.Gen("goto " + $c2.lv[0])
                                                        $lv = $c2.lv
                                                        $lf = $c2.lf
                                                    }
    |   c1=cond op='||' marcador[$c1.lf] c2=cond    {
                                                        $lf = $c2.lf
                                                        $lv = gen.Unir($c1.lv, $c2.lv)
                                                    }
    |   '(' c=cond ')'                              {
                                                        $lv = $c.lv
                                                        $lf = $c.lf
                                                    }
    |   'true'                                      {
                                                        $lv = append($lv, gen.NewEti())
                                                        $lf = append($lf, gen.NewEti())
                                                        gen.Gen("goto " + $lv[0])
                                                    }
    |   'false'                                     {
                                                        $lv = append($lv, gen.NewEti())
                                                        $lf = append($lf, gen.NewEti())
                                                        gen.Gen("goto " + $lf[0])
                                                    }
    ;

/* marcador para soltar etiquetas */
marcador [ []string lista ]
    :   { gen.Soltar(lista) }
    ;

/* expresiones aritmeticas, relacionales y booleanas en conjunto */
expresion returns[string dir, []string lv, []string lf, string cad]
    :   <assoc=right> op='+' e=expresion    {
                                                $dir = gen.NewTemp()
                                                gen.Gen($dir + " = + " + $e.dir)
                                            }
    |   <assoc=right> op='-' e=expresion    {
                                                $dir = gen.NewTemp()
                                                gen.Gen($dir + " = - " + $e.dir)
                                            }
    |   <assoc=right> op='!' e=expresion    {
                                                $lv = $e.lf
                                                $lf = $e.lv
                                            }
    |   e1=expresion op=('*'|'/'|'%') e2=expresion  {
                                                        $dir = gen.NewTemp()
                                                        gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir)
                                                    }
    |   e1=expresion op=('+'|'-') e2=expresion      {
                                                        $dir = gen.NewTemp()
                                                        gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir)
                                                    }
    |   e1=expresion opr=oprel e2=expresion         {
                                                        $lv = append($lv, gen.NewEti())
                                                        $lf = append($lf, gen.NewEti())
                                                        $cad = $e1.dir + " " + $opr.op + " " + $e2.dir
                                                        gen.Gen("if " + $cad + " then goto " + $lv[0])
                                                        gen.Gen("goto " + $lf[0])
                                                    }
    |   e1=expresion op='&&' marcador[$e1.lv] e2=expresion      {
                                                                    $lv = $e2.lv
                                                                    $lf = gen.Unir($e1.lf, $e2.lf)
                                                                }
    |   e1=expresion op='^'  marcador[$e1.lf] e2=expresion      {
                                                                    gen.Soltar($e1.lv)
                                                                    gen.Gen("if " + $e2.cad + " then goto " + $e2.lf[0])
                                                                    gen.Gen("goto " + $e2.lv[0])
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
                                    gen.Gen("goto " + $lv[0])
                                }
    |   'false'                 {
                                    $dir = ""
                                    $cad = ""
                                    $lv = append($lv, gen.NewEti())
                                    $lf = append($lf, gen.NewEti())
                                    gen.Gen("goto " + $lf[0])
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

instrucciones
    :   instruccion*
    ;

instruccion
    :   inst_if
    |   inst_asignacion
    ;

inst_if locals[ string lsalida ]
    @init {
        $lsalida = gen.NewEti()
    }
    :   'if' e1=expresion   {
                                gen.Soltar($e1.lv)
                            }
        bloque  {
                    gen.Gen("goto " + $lsalida)
                    gen.Soltar($e1.lf)
                }
        (
            'else' 'if' e2=expresion    {
                                            gen.Soltar($e2.lv)
                                        }
            bloque  {
                        gen.Gen("goto " + $lsalida)
                        gen.Soltar($e2.lf)
                    }
        )*
        ('else' bloque)? {   gen.Gen($lsalida + ":") }
    ;

inst_asignacion
    :   id=ID '=' e=expresion { gen.Gen($id.text + " = " + $e.dir) }
    ;

bloque
    :   '{' instrucciones '}'
    ;

/* Tokens */
NUM: [0-9]+;
ID: [_A-Za-z][_A-Za-z0-9]*;
COMMENT: '/*' .*? '*/' -> skip;
WHITESPACE: [ \r\n\t]+ -> skip;