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
    :   { tope = entorno.NewEntorno(nil) }
        instrucciones EOF
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
                                                                    gen.GenGoto($e2.lv[0])
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
    |   ref=lref ']'            {
                                    $dir = gen.NewTemp()
                                    gen.Genln($dir + " = " + $ref.id + "[" + $ref.aux + "]")
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
    :   inst_declaracion
    |   inst_asignacion
    |   inst_if
    |   inst_switch_propuesta2
    |   inst_while
    |   inst_doWhile
    |   inst_loop
    ;

/* instruccion de asignacion */
inst_asignacion
    :   id=ID '=' e=expresion ';' { gen.GenAsignacion($id.text, $e.dir) }
    |   ref=lref ']' '=' e=expresion ';'    {
                                                gen.Genln($ref.id + "[" + $ref.aux + "] = " + $e.dir)
                                            }
    ;

lref returns[ string id, string aux, int dim_ ]
    :   ref=lref ',' e=expresion    {
                                        tmp1 := gen.NewTemp()
                                        gen.Genln(tmp1 + " = " + $e.dir + " - 1")

                                        tmp2 := gen.NewTemp()
                                        gen.Genln(tmp2 + " = " + tmp1 + " * " + entorno.GetTamDim($ref.id, $ref.dim_, tope))
                                        gen.Genln($ref.aux + " = " + $ref.aux + " + " + tmp2)

                                        $id = $ref.id
                                        $aux = $ref.aux
                                        $dim_ = $ref.dim_ + 1
                                    }
    |   id_=ID '[' e=expresion      {
                                        $id = $id_.text
                                        $aux = gen.NewTemp()
                                        $dim_ = 1
                                        gen.Genln($aux + " = " + $e.dir)
                                    }
    ;

/* instrucción de declaracion */
inst_declaracion
    :   t=tipo id=ID ';'                {
                                            s := entorno.NewSimbolo($id.text, $t.cad, desp)
                                            desp = desp + 1
                                            tope.Put($id.text, s)
                                        }
    |   'array' '[' dims=ldims ']' id=ID ';'    {
                                                    s := entorno.NewSimboloArr($id.text, $dims.dim_, desp)
                                                    desp = desp + 1
                                                    tope.Put($id.text, s)
                                                }
    ;

ldims returns[ []*entorno.Dim dim_ ]
    :   n1=NUM '..' n2=NUM  {
                                d := entorno.NewDim($n1.int, $n2.int)
                                $dim_ = append($dim_, d)
                            }
        (',' n3=NUM '..' n4=NUM
                            {
                                d := entorno.NewDim($n3.int, $n4.int)
                                $dim_ = append($dim_, d)
                            }
        )*
    ;

tipo returns [ string cad ]
    :   'int' { $cad = "int" }
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

inst_switch_propuesta1 locals[ string lsalida, string lv ]
    @init {
        $lsalida = gen.NewEti()
    }
    :   'switch' e1=expresion '{'
        (
            'case' e2=expresion ':'     {
                                            $lv = gen.NewEti()
                                            gen.GenIf($e1.dir, "!=", $e2.dir, $lv)
                                        }
            (bloque|bloqueSinLLaves)    {
                                            gen.GenGoto($lsalida)
                                            gen.GenDestino($lv)
                                        }
        )+
        (   'default' ':' (bloque|bloqueSinLLaves)  )? { gen.GenDestino($lsalida) }
        '}'
    ;

inst_switch_propuesta2 locals[ string lprueba, string cad, string lsalida, string lv, bool defecto ]
    @init {
        $lprueba = gen.NewEti()
        $lsalida = gen.NewEti()
        $cad = ""
        $defecto = false
    }
    :   'switch' e1=expresion '{'       {
                                            gen.GenGoto($lprueba)
                                        }
        (
            'case' e2=expresion ':'     {
                                            $lv = gen.NewEti()
                                            gen.GenDestino($lv)
                                            $cad += "if " + $e1.dir + " = " + $e2.dir + " then goto " + $lv + "\n"
                                        }
            (bloque|bloqueSinLLaves)    {
                                            gen.GenGoto($lsalida)
                                        }
        )+
        (   'default' ':'               {
                                            $lv = gen.NewEti()
                                            $defecto = true
                                            gen.GenDestino($lv)
                                        }
            (bloque|bloqueSinLLaves)    {
                                            gen.GenGoto($lsalida)
                                        }
        )?                                  {
                                                gen.GenDestino($lprueba)
                                                gen.Gen($cad)
                                                if $defecto {
                                                    gen.GenGoto($lv)
                                                }
                                                gen.GenDestino($lsalida)
                                            }
        '}'
    ;

/* instrucciones ciclicas */
inst_while locals[ string linicio ]
    @init {
        $linicio = gen.NewEti()
    }
    :   'while'         {   gen.GenDestino($linicio)    }
            e=expresion {   gen.Soltar($e.lv)           }
                bloque  {
                            gen.GenGoto($linicio)
                            gen.Soltar($e.lf)
                        }
    ;

inst_doWhile locals[ string linicio ]
    :   'do' {
                $linicio = gen.NewEti()
                gen.GenDestino($linicio)
            }
            bloque 'while' e=expresion ';'  {
                                                gen.Soltar($e.lv)
                                                gen.GenGoto($linicio)
                                                gen.Soltar($e.lf)
                                            }
    ;

inst_loop locals[ string linicio ]
    :   'loop'  {
                    $linicio = gen.NewEti()
                    gen.GenDestino($linicio)
                }
            bloque  {
                        gen.GenGoto($linicio)
                    }
    ;

/* bloque de instrucciones */
bloque
    :   {
            entorno.Push(tope)
            tope = entorno.NewEntorno(tope)
        }
        '{' instrucciones '}'
        {
            tope = entorno.Pop()
        }
    ;

bloqueSinLLaves
    :
        {
            entorno.Push(tope)
            tope = entorno.NewEntorno(tope)
        }
        instrucciones
        {
            tope = entorno.Pop()
        }
    ;

/* Tokens */
NUM: [0-9]+;
ID: [_A-Za-z][_A-Za-z0-9]*;
COMMENT: '/*' .*? '*/' -> skip;
WHITESPACE: [ \r\n\t]+ -> skip;
