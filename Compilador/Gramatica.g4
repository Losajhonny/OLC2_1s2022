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
                                                gen.AddExpresionUnaria($dir, $op.text, $e.dir)
                                            }
    |   <assoc=right> op='-' e=expresion    {
                                                $dir = gen.NewTemp()
                                                gen.AddExpresionUnaria($dir, $op.text, $e.dir)
                                            }
    |   <assoc=right> op='!' e=expresion    {
                                                $lv = $e.lf
                                                $lf = $e.lv
                                            }
    |   e1=expresion op=('*'|'/'|'%') e2=expresion  {
                                                        $dir = gen.NewTemp()
                                                        gen.AddExpresion($dir, $e1.dir, $op.text, $e2.dir)
                                                    }
    |   e1=expresion op=('+'|'-') e2=expresion      {
                                                        $dir = gen.NewTemp()
                                                        gen.AddExpresion($dir, $e1.dir, $op.text, $e2.dir)
                                                    }
    |   e1=expresion opr=oprel e2=expresion         {
                                                        $lv = append($lv, gen.NewLabel())
                                                        $lf = append($lf, gen.NewLabel())
                                                        $cad = gen.AddIf($e1.dir, $opr.op, $e2.dir, $lv[0])
                                                        gen.AddGoto($lf[0])
                                                    }
    |   e1=expresion op='&&' marcador[$e1.lv] e2=expresion      {
                                                                    $lv = $e2.lv
                                                                    $lf = gen.Unir($e1.lf, $e2.lf)
                                                                }
    |   e1=expresion op='^'  marcador[$e1.lf] e2=expresion      {
                                                                    gen.Soltar($e1.lv)
                                                                    gen.AddIfCad($e2.cad, $e2.lf[0])
                                                                    gen.AddGoto($e2.lv[0])
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
                                    $lv = append($lv, gen.NewLabel())
                                    $lf = append($lf, gen.NewLabel())
                                    gen.AddGoto($lv[0])
                                }
    |   'false'                 {
                                    $dir = ""
                                    $cad = ""
                                    $lv = append($lv, gen.NewLabel())
                                    $lf = append($lf, gen.NewLabel())
                                    gen.AddGoto($lf[0])
                                }
    |   ref=lref ']'            {
                                    $dir = gen.NewTemp()
                                    gen.AddGetArray($dir, $ref.id, $ref.aux)
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

/* instrucción de declaracion */
inst_declaracion
    :   'let' id=ID ':' t=tipo ';'            {
                                            s := entorno.NewVar($id.text, $t.cad, desp)
                                            desp = desp + 1
                                            tope.Put($id.text, s)
                                        }
    |   'let' id=ID ':' 'array' '[' dims=ldims ']' 'of' t=tipo ';'
                                                {
                                                    s := entorno.NewVarArray($id.text, $t.cad, desp, 0, $dims.dim_)
                                                    desp = desp + 1
                                                    tope.Put($id.text, s)
                                                }
    ;

ldims returns[ []*entorno.Dimension dim_ ]
    :   n1=NUM '..' n2=NUM  {
                                d := entorno.NewDimension($n1.int, $n2.int)
                                $dim_ = append($dim_, d)
                            }
        (','
            n3=NUM '..' n4=NUM
                            {
                                d := entorno.NewDimension($n3.int, $n4.int)
                                $dim_ = append($dim_, d)
                            }
        )*
    ;

tipo returns [ string cad ]
    :   'int' { $cad = "int" }
    ;

/* instruccion de asignacion */
inst_asignacion
    :   id=ID '=' e=expresion ';'           {   gen.AddAsign($id.text, $e.dir)  }
    |   ref=lref ']' '=' e=expresion ';'    {   gen.AddSetArray($ref.id, $ref.aux, $e.dir)  }
    ;

lref returns[ string id, string aux, int dim_ ]
    :   ref=lref ',' e=expresion    {
                                        tmp1 := gen.NewTemp()
                                        gen.AddExpresion(tmp1, $e.dir, "-", "1")

                                        tmp2 := gen.NewTemp()
                                        gen.AddExpresion(tmp2, tmp1, "*", entorno.GetTamOrdenColumn($ref.id, $ref.dim_, tope))
                                        gen.AddExpresion($ref.aux, $ref.aux, "+", tmp2)

                                        $id = $ref.id
                                        $aux = $ref.aux
                                        $dim_ = $ref.dim_ + 1
                                    }
    |   id_=ID '[' e=expresion      {
                                        $id = $id_.text
                                        $aux = gen.NewTemp()
                                        $dim_ = 1
                                        gen.AddAsign($aux, $e.dir)
                                    }
    ;

/* instrucciones de seleccion */
inst_if locals[ string lsalida ]
    :   'if' e1=expresion   {
                                $lsalida = gen.NewLabel()
                                gen.Soltar($e1.lv)
                            }
        bloque  {
                    gen.AddGoto($lsalida)
                    gen.Soltar($e1.lf)
                }
        (
            'else' 'if' e2=expresion    {
                                            gen.Soltar($e2.lv)
                                        }
            bloque  {
                        gen.AddGoto($lsalida)
                        gen.Soltar($e2.lf)
                    }
        )*
        ('else' bloque)? {   gen.AddLabel($lsalida) }
    ;

inst_switch_propuesta1 locals[ string lsalida, string lv ]
    @init {
        $lsalida = gen.NewLabel()
    }
    :   'switch' e1=expresion '{'
        (
            'case' e2=expresion ':'     {
                                            $lv = gen.NewLabel()
                                            gen.AddIf($e1.dir, "!=", $e2.dir, $lv)
                                        }
            (bloque|bloqueSinLLaves)    {
                                            gen.AddGoto($lsalida)
                                            gen.AddLabel($lv)
                                        }
        )+
        (
            'default' ':'
            (bloque|bloqueSinLLaves)
        )?                              {
                                            gen.AddLabel($lsalida)
                                        }
        '}'
    ;

inst_switch_propuesta2 locals[ string lprueba, string cad, string lsalida, string lv, bool defecto ]
    @init {
        $lprueba = gen.NewLabel()
        $lsalida = gen.NewLabel()
        $cad = ""
        $defecto = false
    }
    :   'switch' e1=expresion '{'       {
                                            gen.AddGoto($lprueba)
                                        }
        (
            'case' e2=expresion ':'     {
                                            $lv = gen.NewLabel()
                                            gen.AddLabel($lv)
                                            $cad += "if " + $e1.dir + " = " + $e2.dir + " then goto " + $lv + "\n"
                                        }
            (bloque|bloqueSinLLaves)    {
                                            gen.AddGoto($lsalida)
                                        }
        )+
        (   'default' ':'               {
                                            $lv = gen.NewLabel()
                                            $defecto = true
                                            gen.AddLabel($lv)
                                        }
            (bloque|bloqueSinLLaves)    {
                                            gen.AddGoto($lsalida)
                                        }
        )?                              {
                                            gen.AddLabel($lprueba)
                                            gen.Gen($cad)
                                            if $defecto {
                                                gen.AddGoto($lv)
                                            }
                                            gen.AddLabel($lsalida)
                                        }
        '}'
    ;

/* instrucciones ciclicas */
inst_while locals[ string linicio ]
    @init {
        $linicio = gen.NewLabel()
    }
    :   'while'         {   gen.AddLabel($linicio)      }
            e=expresion {   gen.Soltar($e.lv)           }
                bloque  {
                            gen.AddGoto($linicio)
                            gen.Soltar($e.lf)
                        }
    ;

inst_doWhile locals[ string linicio ]
    :   'do' {
                $linicio = gen.NewLabel()
                gen.AddLabel($linicio)
            }
            bloque 'while' e=expresion ';'  {
                                                gen.Soltar($e.lv)
                                                gen.AddGoto($linicio)
                                                gen.Soltar($e.lf)
                                            }
    ;

inst_loop locals[ string linicio ]
    :   'loop'  {
                    $linicio = gen.NewLabel()
                    gen.AddLabel($linicio)
                }
            bloque  {
                        gen.AddGoto($linicio)
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
FLOAT: NUM'.'NUM;
STRING: '"'~["]*'"';
ID: [_A-Za-z][_A-Za-z0-9]*;
MULCOMMENT: '/*' .*? '*/' -> skip;
UNICOMMENT: '//'~[\r\n]* -> skip;
WHITESPACE: [ \\\r\n\t]+ -> skip;
