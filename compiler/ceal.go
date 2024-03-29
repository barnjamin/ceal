package compiler

import (
	"ceal/teal"
	"fmt"
	"strings"
)

type CealStatement interface {
	TealAst() teal.TealAst
}

type CealProgram struct {
	ConstInts  []int
	ConstBytes [][]byte

	Functions     map[string]*CealFunction
	FunctionNames []string
}

func (a *CealProgram) registerFunction(f *CealFunction) {
	if _, ok := a.Functions[f.Fun.name]; ok {
		panic(fmt.Sprintf("function '%s' is already defined", f.Fun.name))
	}

	a.Functions[f.Fun.name] = f
	a.FunctionNames = append(a.FunctionNames, f.Fun.name)

}

func (a *CealProgram) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	res.Write(&teal.Teal_pragma_version{Version: AvmVersion})

	if len(a.ConstInts) > 0 {
		items := []uint64{}
		for _, v := range a.ConstInts {
			items = append(items, uint64(v))
		}

		res.Write(&teal.Teal_intcblock_fixed{UINT1: items})
	}

	if len(a.ConstBytes) > 0 {
		res.Write(&teal.Teal_bytecblock_fixed{BYTES1: a.ConstBytes})
	}

	main := a.Functions[AvmMainName]

	if len(a.Functions) > 1 {
		res.Write(&teal.Teal_b_fixed{TARGET1: main.Fun.name})
	}

	for _, name := range a.FunctionNames {
		ast := a.Functions[name]

		if name == AvmMainName {
			continue
		}

		res.Write(&teal.Teal_label{Name: ast.Fun.name})
		res.Write(ast.TealAst())
	}

	if len(a.Functions) > 1 {
		res.Write(&teal.Teal_label{Name: main.Fun.name})
	}

	res.Write(main.TealAst())

	return res.Build()
}

type CealContinue struct {
	Label string
	Index int
}

func (a *CealContinue) TealAst() teal.TealAst {
	return &teal.Teal_b_fixed{TARGET1: fmt.Sprintf("%s_%d_continue", a.Label, a.Index)}
}

type CealBreak struct {
	Label string
	Index int
}

func (a *CealBreak) TealAst() teal.TealAst {
	return &teal.Teal_b_fixed{TARGET1: fmt.Sprintf("%s_%d_end", a.Label, a.Index)}
}

type CealSwitchCase struct {
	Value      CealStatement
	Statements []CealStatement
}

type CealSwitch struct {
	Index int
	Loop  *LoopScopeItem

	Value CealStatement
	Cases []*CealSwitchCase

	Default []CealStatement
}

func (a *CealSwitch) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	labels := []string{}

	for i, c := range a.Cases {
		label := fmt.Sprintf("switch_%d_%d", a.Index, i)
		labels = append(labels, label)

		res.Write(c.Value.TealAst())
	}

	res.Write(a.Value.TealAst())
	res.Write(&teal.Teal_match_fixed{TARGET1: labels})

	if len(a.Default) > 0 {
		for _, stmt := range a.Default {
			res.Write(stmt.TealAst())
		}
	}

	for i, c := range a.Cases {
		label := labels[i]

		res.Write(&teal.Teal_label{Name: label})

		for _, stmt := range c.Statements {
			res.Write(stmt.TealAst())
		}
	}

	if a.Loop.breaks {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("switch_%d_end", a.Index)})
	}

	return res.Build()
}

type CealDoWhile struct {
	Index     int
	Loop      *LoopScopeItem
	Condition CealStatement
	Statement CealStatement
}

func (a *CealDoWhile) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}
	res.Write(&teal.Teal_label{Name: fmt.Sprintf("do_%d", a.Index)})
	res.Write(a.Statement.TealAst())

	if a.Loop.continues {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("do_%d_continue", a.Index)})
	}

	res.Write(&teal.Teal_bnz_fixed{
		S1:      a.Condition.TealAst(),
		TARGET1: fmt.Sprintf("do_%d", a.Index),
	})

	if a.Loop.breaks {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("do_%d_end", a.Index)})
	}

	return res.Build()
}

type CealWhile struct {
	Index     int
	Loop      *LoopScopeItem
	Condition CealStatement
	Statement CealStatement
}

func (a *CealWhile) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	res.Write(&teal.Teal_label{Name: fmt.Sprintf("while_%d", a.Index)})
	res.Write(&teal.Teal_bz_fixed{
		S1:      a.Condition.TealAst(),
		TARGET1: fmt.Sprintf("while_%d_end", a.Index),
	})

	res.Write(a.Statement.TealAst())

	if a.Loop.continues {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("while_%d_continue", a.Index)})
	}

	res.Write(&teal.Teal_b_fixed{TARGET1: fmt.Sprintf("while_%d", a.Index)})
	res.Write(&teal.Teal_label{Name: fmt.Sprintf("while_%d_end", a.Index)})

	return res.Build()
}

type CealFor struct {
	Index     int
	Loop      *LoopScopeItem
	Init      []CealStatement
	Condition CealStatement
	Statement CealStatement
	Iter      []CealStatement
}

func (a *CealFor) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	for _, stmt := range a.Init {
		res.Write(stmt.TealAst())
	}

	res.Write(&teal.Teal_label{Name: fmt.Sprintf("for_%d", a.Index)})
	res.Write(&teal.Teal_bz_fixed{
		S1:      a.Condition.TealAst(),
		TARGET1: fmt.Sprintf("for_%d_end", a.Index),
	})
	res.Write(a.Statement.TealAst())

	if a.Loop.continues {
		res.Write(&teal.Teal_label{Name: fmt.Sprintf("for_%d_continue", a.Index)})
	}

	for _, stmt := range a.Iter {
		res.Write(stmt.TealAst())
	}

	res.Write(&teal.Teal_b_fixed{TARGET1: fmt.Sprintf("for_%d", a.Index)})
	res.Write(&teal.Teal_label{Name: fmt.Sprintf("for_%d_end", a.Index)})

	return res.Build()
}

type CealExpr interface {
	// ToStmt converts the expression to a statement so it does not push a value onto the stack
	ToStmt()
}

type CealValue interface {
	ToValue()
}

type CealPrefix struct {
	V  *Variable
	Op string

	IsStmt bool
}

func (a *CealPrefix) ToStmt() {
	a.IsStmt = true
}

func (a *CealPrefix) TealAst() teal.TealAst {
	if a.V.constant {
		panic("cannot modify const var")
	}

	var op teal.TealAst

	switch a.Op {
	case "++":
		op = &teal.Teal_plus{
			STACK_1: &teal.Teal_load{I1: uint8(a.V.local.slot)},
			STACK_2: &teal.Teal_int{V: 1},
		}
	case "--":
		op = &teal.Teal_minus{
			STACK_1: &teal.Teal_load{I1: uint8(a.V.local.slot)},
			STACK_2: &teal.Teal_int{V: 1},
		}
	default:
		panic(fmt.Sprintf("prefix operator not supported: '%s'", a.Op))
	}

	res := &teal.TealAstBuilder{}
	res.Write(&teal.Teal_store{STACK_1: op, I1: uint8(a.V.local.slot)})

	if !a.IsStmt {
		res.Write(&teal.Teal_load{I1: uint8(a.V.local.slot)})
	}

	return res.Build()
}

type CealPostfix struct {
	V  *Variable
	Op string

	IsStmt bool
}

func (a *CealPostfix) ToStmt() {
	a.IsStmt = true
}

func (a *CealPostfix) TealAst() teal.TealAst {
	if a.V.constant {
		panic("cannot modify const var")
	}

	var s1 teal.TealAst
	s1 = &teal.Teal_load{I1: uint8(a.V.local.slot)}

	if !a.IsStmt {
		s1 = &teal.Teal_dup{STACK_1: s1}
	}

	s2 := &teal.Teal_int{V: 1}

	var op teal.TealAst
	switch a.Op {
	case "++":
		op = &teal.Teal_plus{STACK_1: s1, STACK_2: s2}
	case "--":
		op = &teal.Teal_minus{STACK_1: s1, STACK_2: s2}
	default:
		panic(fmt.Sprintf("postfix operator not supported: '%s'", a.Op))
	}

	return &teal.Teal_store{STACK_1: op, I1: uint8(a.V.local.slot)}
}

type CealLabel struct {
	Name string
}

func (a *CealLabel) TealAst() teal.TealAst {
	return &teal.Teal_label{Name: fmt.Sprintf("label_%s", a.Name)}
}

type CealGoto struct {
	Label string
}

func (a *CealGoto) TealAst() teal.TealAst {
	return &teal.Teal_b_fixed{TARGET1: fmt.Sprintf("label_%s", a.Label)}
}

type CealVariable struct {
	V *Variable
}

func (a *CealVariable) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	if a.V.local != nil {
		ast := &teal.Teal_load{I1: uint8(a.V.local.slot)}
		res.Write(ast)
		return res.Build()
	}

	if a.V.param != nil {
		ast := &teal.Teal_frame_dig{I1: int8(a.V.param.index)}
		res.Write(ast)
		return res.Build()
	}

	if a.V.const_ != nil {
		switch a.V.const_.kind {
		case SimpleTypeInt:
			ast := &teal.Teal_intc{
				I1: uint8(a.V.const_.index),
			}
			res.Write(ast)
			return res.Build()
		case SimpleTypeBytes:
			ast := &teal.Teal_bytec{
				I1: uint8(a.V.const_.index),
			}
			res.Write(ast)
			return res.Build()
		}
	}

	switch a.V.t {
	case "uint64":
		res.Write(&teal.Teal_named_int{V: &teal.Teal_named_int_value{V: a.V.name}})
		return res.Build()
	case "bytes":
		res.Write(&teal.Teal_byte{S: &teal.Teal_byte_string_value{V: a.V.name}})
	default:
		panic(fmt.Sprintf("type '%s' is not supported", a.V.t))
	}

	return res.Build()
}

type CealUnaryOp struct {
	Op        string
	Statement CealStatement
}

func (a *CealUnaryOp) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}
	switch a.Op {
	case "!":
		res.Write(&teal.Teal_not{STACK_1: a.Statement.TealAst()})
	default:
		panic(fmt.Sprintf("unary op '%s' not supported", a.Op))
	}
	return res.Build()
}

type CealAssignSumDiff struct {
	V     *Variable
	F     *StructField
	T     *Type
	Value CealStatement
	Op    string

	IsStmt bool
}

func (a *CealAssignSumDiff) TealAst() teal.TealAst {
	var slot uint8

	if a.T.complex != nil {
		v := a.V.fields[a.F.name]
		slot = uint8(v.local.slot)
	} else {
		slot = uint8(a.V.local.slot)
	}

	s1 := &teal.Teal_load{I1: slot}

	var op teal.TealAst

	switch a.Op {
	case "+=":
		op = &teal.Teal_plus{STACK_1: s1, STACK_2: a.Value.TealAst()}
	case "-=":
		op = &teal.Teal_minus{STACK_1: s1, STACK_2: a.Value.TealAst()}
	}

	if !a.IsStmt {
		op = &teal.Teal_dup{STACK_1: op}
	}

	return &teal.Teal_store{STACK_1: op, I1: slot}
}

type CealAnd struct {
	Index int

	Left  CealStatement
	Right CealStatement
}

func (a *CealAnd) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	res.Write(&teal.Teal_andand{
		STACK_1: &teal.Teal_bz_fixed{
			S1:      &teal.Teal_dup{STACK_1: a.Left.TealAst()},
			TARGET1: fmt.Sprintf("and_%d_end", a.Index),
		},
		STACK_2: a.Right.TealAst(),
	})

	res.Write(&teal.Teal_label{Name: fmt.Sprintf("and_%d_end", a.Index)})

	return res.Build()
}

type CealOr struct {
	Index int
	Left  CealStatement
	Right CealStatement
}

func (a *CealOr) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	res.Write(
		&teal.Teal_oror{
			STACK_1: &teal.Teal_bnz_fixed{
				S1: &teal.Teal_dup{
					STACK_1: a.Left.TealAst(),
				},
				TARGET1: fmt.Sprintf("or_%d_end", a.Index),
			},
			STACK_2: a.Right.TealAst(),
		})

	res.Write(&teal.Teal_label{Name: fmt.Sprintf("or_%d_end", a.Index)})

	return res.Build()
}

type CealBinop struct {
	Left  CealStatement
	Op    string
	Right CealStatement
}

func (a *CealBinop) TealAst() teal.TealAst {
	var op teal.TealAst

	switch a.Op {
	case "+":
		op = &teal.Teal_plus{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "-":
		op = &teal.Teal_minus{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "*":
		op = &teal.Teal_mul{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "/":
		op = &teal.Teal_div{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "==":
		op = &teal.Teal_eqeq{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "!=":
		op = &teal.Teal_noteq{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "&":
		op = &teal.Teal_and{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "^":
		op = &teal.Teal_xor{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "|":
		op = &teal.Teal_or{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "<":
		op = &teal.Teal_lt{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case ">":
		op = &teal.Teal_gt{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "<=":
		op = &teal.Teal_lteq{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case ">=":
		op = &teal.Teal_gteq{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	case "%":
		op = &teal.Teal_mod{STACK_1: a.Left.TealAst(), STACK_2: a.Right.TealAst()}
	default:
		panic(fmt.Sprintf("binary op '%s' is not supported yet", a.Op))
	}

	return op
}

type CealNegate struct {
	Value CealStatement
}

func (a *CealNegate) TealAst() teal.TealAst {
	return &teal.Teal_minus{
		STACK_1: &teal.Teal_int{V: 0},
		STACK_2: a.Value.TealAst(),
	}
}

type CealDefine struct {
	V *Variable
	T *Type

	Value CealStatement
}

func (a *CealDefine) TealAst() teal.TealAst {
	if a.T.complex != nil {
		panic("defining complex variable is not supported yet")
	}

	ast := &teal.Teal_store{
		STACK_1: a.Value.TealAst(),
		I1:      uint8(a.V.local.slot),
	}

	return ast
}

type CealAssign struct {
	V   *Variable
	T   *Type
	F   *StructField
	Fun *Function

	Value CealStatement

	IsStmt bool
}

func (a *CealAssign) ToStmt() {
	a.IsStmt = true
}

func (a *CealAssign) TealAst() teal.TealAst {
	if a.V.constant {
		panic("cannot assign to a const var")
	}

	if a.V.param != nil {
		// TODO: add param var assignment support
		panic("cannot assign param var")
	}

	res := &teal.TealAstBuilder{}

	if a.T.complex != nil {
		if a.T.complex.builtin != nil {
			res.Write(&teal.Teal_call_builtin{
				Name: a.Fun.builtin.op,
				Imms: []teal.TealAst{&teal.Teal_named_int_value{
					V: a.F.name,
				}},
			})
			return res.Build()
		} else {
			if a.V.param != nil {
				panic("accessing struct param fields is not supported yet")
			}

			v := a.V.fields[a.F.name]
			ast := &teal.Teal_store{
				STACK_1: a.Value.TealAst(),
				I1:      uint8(v.local.slot),
			}

			res.Write(ast)
		}
	} else {
		ast := &teal.Teal_store{
			STACK_1: a.Value.TealAst(),
			I1:      uint8(a.V.local.slot),
		}

		res.Write(ast)
	}

	if !a.IsStmt {
		load := &teal.Teal_load{
			I1: uint8(a.V.local.slot),
		}

		res.Write(load)
	}

	return res.Build()
}

type CealStructField struct {
	V   *Variable
	T   *Type
	F   *StructField
	Fun *Function
}

func (a *CealStructField) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	if a.T.complex.builtin != nil {
		res.Write(&teal.Teal_call_builtin{
			Name: a.Fun.builtin.op,
			Imms: []teal.TealAst{&teal.Teal_named_int_value{
				V: a.F.name,
			}},
		})
		return res.Build()
	}

	if a.V.param != nil {
		panic("accessing struct param fields is not supported yet")
	}

	v := a.V.fields[a.F.name]

	ast := &teal.Teal_load{
		I1: uint8(v.local.slot),
	}

	res.Write(ast)

	return res.Build()
}

type CealCall struct {
	Fun  *Function
	Args []CealStatement

	IsStmt bool
}

func (a *CealCall) ToStmt() {
	a.IsStmt = true
}

func (a *CealCall) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	var args []teal.TealAst

	if a.Fun.compiler != nil {
		ast := a.Fun.compiler.handler(a.Args)
		res.Write(ast)
	}

	if a.Fun.builtin != nil {
		i := 0

		for ; i < len(a.Fun.builtin.stack); i++ {
			arg := a.Args[i]
			args = append(args, arg.TealAst())
		}

		var imms []teal.TealAst

		for ; i < len(a.Fun.builtin.stack)+len(a.Fun.builtin.imm); i++ {
			arg := a.Args[i]
			if e, ok := arg.(CealValue); ok {
				e.ToValue()
			}

			ast := arg.TealAst()
			imms = append(imms, ast)
		}

		res.Write(&teal.Teal_call_builtin{
			Args: args,
			Imms: imms,
			Name: a.Fun.builtin.op,
		})
	}

	if a.Fun.user != nil {
		for _, arg := range a.Args {
			args = append(args, arg.TealAst())
		}

		res.Write(&teal.Teal_callsub_fixed{
			Args:   args,
			Target: a.Fun.name,
		})
	}

	if a.IsStmt {
		if a.Fun.returns > 0 {
			res.Write(&teal.Teal_popn{N1: uint8(a.Fun.returns)})
		}
	}

	return res.Build()
}

type CealIsReturn interface {
	IsReturn()
}

type CealIsBreak interface {
	IsBreak()
}

type CealIntConstant struct {
	Value string
	value bool
}

func (a *CealIntConstant) ToValue() {
	a.value = true
}

func (a *CealIntConstant) TealAst() teal.TealAst {
	var v teal.TealAst = &teal.Teal_named_int_value{V: a.Value}
	if !a.value {
		v = &teal.Teal_named_int{V: v}
	}
	return v
}

type CealByteConstant struct {
	Value string
	value bool
}

func (a *CealByteConstant) ToValue() {
	a.value = true
}

func (a *CealByteConstant) TealAst() teal.TealAst {
	var v teal.TealAst = &teal.Teal_byte_string_value{V: a.Value}
	if !a.value {
		v = &teal.Teal_byte{S: v}
	}
	return v
}

type CealReturn struct {
	Value CealStatement
	Fun   *Function
}

func (a *CealReturn) IsReturn() {
}

func (a *CealReturn) TealAst() teal.TealAst {
	var op teal.TealAst

	if a.Fun != nil && a.Fun.user.sub {
		var values []teal.TealAst

		if a.Value != nil {
			values = append(values, a.Value.TealAst())
		}

		op = &teal.Teal_retsub_fixed{
			Values: values,
		}
	} else {
		op = &teal.Teal_return_fixed{
			Value: a.Value.TealAst(),
		}
	}

	return op
}

type CealBlock struct {
	Statements []CealStatement
}

func (a *CealBlock) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	for _, stmt := range a.Statements {
		res.Write(stmt.TealAst())
	}

	return res.Build()
}

type CealConditional struct {
	Index int

	Condition CealStatement

	True  CealStatement
	False CealStatement
}

func (a *CealConditional) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	false_label := fmt.Sprintf("conditional_%d_false", a.Index)
	end_label := fmt.Sprintf("conditional_%d_end", a.Index)

	res.Write(&teal.Teal_bz_fixed{
		S1:      a.Condition.TealAst(),
		TARGET1: false_label,
	})
	res.Write(a.True.TealAst())
	res.Write(&teal.Teal_b_fixed{TARGET1: end_label})
	res.Write(&teal.Teal_label{Name: false_label})
	res.Write(a.False.TealAst())
	res.Write(&teal.Teal_label{Name: end_label})

	return res.Build()
}

type CealIfAlternative struct {
	Condition  CealStatement
	Statements []CealStatement
}

type CealIf struct {
	Index        int
	Alternatives []*CealIfAlternative
}

func (a *CealIf) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	end_label := fmt.Sprintf("if_end_%d", a.Index)

	for i, alt := range a.Alternatives {
		if alt.Condition != nil {
			if i < len(a.Alternatives)-1 {
				res.Write(&teal.Teal_bz_fixed{
					S1:      alt.Condition.TealAst(),
					TARGET1: fmt.Sprintf("if_skip_%d_%d", a.Index, i),
				})
			} else {
				res.Write(&teal.Teal_bz_fixed{
					S1:      alt.Condition.TealAst(),
					TARGET1: end_label,
				})
			}
		}

		for _, stmt := range alt.Statements {
			res.Write(stmt.TealAst())
		}

		if i < len(a.Alternatives)-1 {
			res.Write(&teal.Teal_b_fixed{TARGET1: end_label})
		}

		if alt.Condition != nil {
			if i < len(a.Alternatives)-1 {
				res.Write(&teal.Teal_label{Name: fmt.Sprintf("if_skip_%d_%d", a.Index, i)})
			}
		}
	}

	res.Write(&teal.Teal_label{Name: end_label})

	return res.Build()
}

type CealFunction struct {
	Fun        *Function
	Statements []CealStatement
}

func (a *CealFunction) TealAst() teal.TealAst {
	res := &teal.TealAstBuilder{}

	if a.Fun.user.sub {
		if a.Fun.user.args != 0 || a.Fun.returns != 0 {
			ast := &teal.Teal_proto{
				A1: uint8(a.Fun.user.args),
				R2: uint8(a.Fun.returns),
			}

			res.Write(ast)
		}
	}

	for _, stmt := range a.Statements {
		res.Write(stmt.TealAst())
	}

	if a.Fun.user.sub {
		if len(a.Statements) > 0 {
			last := a.Statements[len(a.Statements)-1]
			if _, ok := last.(CealIsReturn); !ok {
				res.Write(&teal.Teal_retsub{})
			}
		}
	}

	return res.Build()
}

type CealRaw struct {
	Value string
}

func (a *CealRaw) String() string {
	return a.Value
}

func (a *CealRaw) Teal() teal.Teal {
	return teal.Teal{a}
}

func (a *CealRaw) TealAst() teal.TealAst {
	return a
}

type CealSingleLineComment struct {
	Line string
}

func (a *CealSingleLineComment) TealAst() teal.TealAst {
	return &teal.Teal_comment{Lines: []string{a.Line}}
}

type CealMultiLineComment struct {
	Text string
}

func (a *CealMultiLineComment) TealAst() teal.TealAst {
	lines := strings.Split(a.Text, "\n")
	return &teal.Teal_comment{Lines: lines}
}
