package compiler

import (
	"ceal/parser"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func itoa(v int) string {
	return strconv.Itoa(v)
}

func atoi(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}

type FunctionParam struct {
	t    string
	name string
}

type BuiltinFunction struct {
	op    string
	stack []*FunctionParam
	imm   []*FunctionParam
}

type Function struct {
	t    string
	name string

	builtin *BuiltinFunction
	user    *UserFunction
}

type StructField struct {
	t    string
	name string
	fun  string
}

type StructFunction struct {
	t    string
	name string

	params []*FunctionParam
}

type BuiltinStruct struct {
}

type Struct struct {
	fields    map[string]*StructField
	functions map[string]*StructFunction

	builtin *BuiltinStruct
}

type SimpleType struct {
	empty bool
}

type Type struct {
	name    string
	complex *Struct // struct
	simple  *SimpleType
}

type LocalVariable struct {
	slot int
}

type ParameterVariable struct {
	index int
}

type Variable struct {
	t    string
	name string

	readonly bool

	local *LocalVariable
	param *ParameterVariable

	fields map[string]*Variable
}

type Scope struct {
	types     map[string]*Type
	functions map[string]*Function
	variables map[string]*Variable

	children []*Scope

	function *Function

	read bool
	i    int

	parent *Scope
}

func (vt *SymbolTableVisitor) initVariable(v *Variable) {
	t := vt.scope.resolveType(v.t)

	if t.simple != nil {
		return
	}

	if t.complex.builtin != nil {
		return
	}

	if v.param != nil {
		panic("struct types are not supported for parameters yet")
	}

	fields := map[string]*Variable{}

	for _, f := range t.complex.fields {
		v := &Variable{
			name: f.name,
			t:    f.t,
			local: &LocalVariable{
				slot: vt.slot,
			},
		}

		vt.slot++

		vt.initVariable(v)

		fields[f.name] = v
	}

	v.fields = fields
}

func (s *Scope) resolveType(typeName string) *Type {
	current := s

	for current != nil {
		if t, ok := current.types[typeName]; ok {
			return t
		}

		current = current.parent
	}

	return nil
}

func NewScope(parent *Scope) *Scope {
	s := &Scope{
		parent:    parent,
		types:     map[string]*Type{},
		functions: map[string]*Function{},
		variables: map[string]*Variable{},
	}

	return s
}

func (s *Scope) enter() *Scope {
	if s.read {
		if s.i >= len(s.children) {
			return nil
		}

		scope := s.children[s.i]
		s.i++

		return scope
	}

	child := NewScope(s)
	s.children = append(s.children, child)

	return child
}

func (s *Scope) exit() *Scope {
	return s.parent
}

func (s *Scope) readonly() {
	s.read = true

	for _, f := range s.functions {
		user := f.user
		if user != nil {
			scope := user.scope
			if scope != nil {
				scope.readonly()
			}
		}
	}

	for _, item := range s.children {
		item.readonly()
	}

	s.i = 0
}

type UserFunction struct {
	args    int
	returns int

	sub bool

	scope *Scope
}

type SymbolTableVisitor struct {
	*parser.BaseCVisitor

	global *Scope
	scope  *Scope

	slot int // TODO: make the slot index stack-based
}

func (v *SymbolTableVisitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	id := ctx.ID().GetText()

	if _, ok := v.scope.variables[id]; ok {
		panic(fmt.Sprintf("variable '%s' is already defined", id))
	}

	local := &LocalVariable{
		slot: v.slot,
	}

	v.slot++

	t := ctx.Type_().ID().GetText()

	vr := &Variable{
		t:     t,
		name:  id,
		local: local,
	}

	v.initVariable(vr)
	v.scope.variables[vr.name] = vr

	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitDefinition(ctx *parser.DefinitionContext) interface{} {
	id := ctx.ID().GetText()

	if _, ok := v.scope.variables[id]; ok {
		panic(fmt.Sprintf("variable '%s' is already defined", id))
	}

	local := &LocalVariable{
		slot: v.slot,
	}

	v.slot++

	vr := &Variable{
		t:     ctx.Type_().ID().GetText(),
		name:  id,
		local: local,
	}

	v.initVariable(vr)
	v.scope.variables[vr.name] = vr

	return v.VisitChildren(ctx)
}

func (v *SymbolTableVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.scope = v.scope.enter()
	v.VisitChildren(ctx)
	v.scope = v.scope.exit()

	return nil
}

func (v *SymbolTableVisitor) VisitStruct(ctx *parser.StructContext) interface{} {
	name := ctx.ID().GetText()

	if t := v.scope.resolveType(name); t != nil {
		panic(fmt.Sprintf("type '%s' is already defined", name))
	}

	s := &Struct{
		fields:    map[string]*StructField{},
		functions: map[string]*StructFunction{},
	}

	for _, item := range ctx.AllField() {
		t := item.Type_().ID().GetText()
		name := item.ID().GetText()

		f := &StructField{
			t:    t,
			name: name,
		}

		if _, ok := s.fields[name]; ok {
			panic(fmt.Sprintf("field '%s' is already defined", name))
		}

		s.fields[name] = f
	}

	t := &Type{
		name:    name,
		complex: s,
	}

	v.global.types[t.name] = t

	return nil
}

func (v *SymbolTableVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.scope = v.global
	v.VisitChildren(ctx)
	v.scope = nil

	return nil
}

func (v *SymbolTableVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	id := ctx.ID().GetText()

	if _, ok := v.scope.functions[id]; ok {
		panic(fmt.Sprintf("function '%s' already defined", id))
	}

	user := &UserFunction{
		scope: NewScope(v.scope),
		args:  len(ctx.Params().AllParam()),
	}

	ret := ctx.Type_().ID().GetText()
	t := v.scope.resolveType(ret)

	if t.complex != nil {
		user.returns = len(t.complex.fields)
	} else {
		if t.simple.empty {
			user.returns = 0
		} else {
			user.returns = 1
		}
	}

	user.sub = id != AvmMainName

	fun := &Function{
		t:    ret,
		name: id,
		user: user,
	}

	v.scope.functions[id] = fun

	v.scope = user.scope
	v.scope.function = fun

	index := -len(ctx.Params().AllParam())

	for _, pctx := range ctx.Params().AllParam() {
		id := pctx.ID().GetText()

		if _, ok := v.scope.variables[id]; ok {
			panic(fmt.Sprintf("param '%s' already defined", id))
		}

		param := &ParameterVariable{
			index: index,
		}

		index++

		vr := &Variable{
			t:     pctx.Type_().ID().GetText(),
			name:  id,
			param: param,
		}

		v.initVariable(vr)
		v.scope.variables[id] = vr
	}

	v.VisitChildren(ctx)

	v.scope = v.scope.exit()

	return nil
}

func (v *SymbolTableVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *SymbolTableVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

type AstVisitor struct {
	*parser.BaseCVisitor

	global *Scope
	scope  *Scope

	program *AstProgram

	label int
}

func (v *AstVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *AstVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v)
	}
	return nil
}

type AstIsReturn interface {
	IsReturn()
}

type AstVariable struct {
	v *Variable
}

func (a *AstVariable) String() string {
	if a.v.local != nil {
		ast := avm_load_Ast{i1: itoa(a.v.local.slot)}
		return ast.String()
	}

	if a.v.param != nil {
		ast := avm_frame_dig_Ast{i1: itoa(a.v.param.index)}
		return ast.String()
	}

	if a.v.t == "uint64" {
		return fmt.Sprintf("int %s", a.v.name)
	}

	return fmt.Sprintf("byte %s", a.v.name)
}

type AstBinop struct {
	l  AstStatement
	op string
	r  AstStatement
}

func (a *AstBinop) String() string {
	return fmt.Sprintf("%s\n%s\n%s", a.l.String(), a.r.String(), a.op)
}

type AstMinusOp struct {
	value AstStatement
}

func (a *AstMinusOp) String() string {
	return fmt.Sprintf("int 0\n%s\n-", a.value.String())
}

type AstAssign struct {
	v   *Variable
	t   *Type
	f   *StructField
	fun *Function

	value AstStatement
}

func (a *AstAssign) String() string {
	if a.v.param != nil {
		// TODO: add param var assignment support
		panic("cannot assign param var")
	}

	if a.t.complex != nil {
		if a.t.complex.builtin != nil {
			return fmt.Sprintf("%s %s", a.fun.builtin.op, a.f.name)
		} else {
			if a.v.param != nil {
				panic("accessing struct param fields is not supported yet")
			}

			v := a.v.fields[a.f.name]

			ast := avm_store_Ast{
				s1: a.value,
				i1: itoa(v.local.slot),
			}

			return ast.String()
		}
	} else {
		ast := avm_store_Ast{
			s1: a.value,
			i1: itoa(a.v.local.slot),
		}

		return ast.String()
	}
}

type AstStructField struct {
	v   *Variable
	t   *Type
	f   *StructField
	fun *Function
}

func (a *AstStructField) String() string {
	if a.t.complex.builtin != nil {
		return fmt.Sprintf("%s %s", a.fun.builtin.op, a.f.name)
	}

	if a.v.param != nil {
		panic("accessing struct param fields is not supported yet")
	}

	v := a.v.fields[a.f.name]

	ast := avm_load_Ast{
		i1: itoa(v.local.slot),
	}

	return ast.String()
}

type AstCall struct {
	fun  *Function
	args []AstStatement
}

func (a *AstCall) String() string {
	s := strings.Builder{}

	if a.fun.builtin != nil {
		i := 0

		for ; i < len(a.fun.builtin.stack); i++ {
			arg := a.args[i]
			s.WriteString(arg.String())
			s.WriteString("\n")
		}

		s.WriteString(a.fun.builtin.op)

		for ; i < len(a.fun.builtin.stack)+len(a.fun.builtin.imm); i++ {
			arg := a.args[i]
			s.WriteString(" ")
			s.WriteString(arg.String())
		}
	}

	if a.fun.user != nil {
		for _, arg := range a.args {
			s.WriteString(arg.String())
			s.WriteString("\n")
		}

		s.WriteString(fmt.Sprintf("callsub %s", a.fun.name))
	}

	return s.String()
}

type AstRaw struct {
	value string
}

func (a *AstRaw) String() string {
	return a.value
}

type AstIntConstant struct {
	value string
}

func (a *AstIntConstant) String() string {
	return fmt.Sprintf("int %s", a.value)
}

type AstByteConstant struct {
	value string
}

func (a *AstByteConstant) String() string {
	return fmt.Sprintf("byte %s", a.value)
}

type AstReturn struct {
	value    AstStatement
	function *Function
}

func (a *AstReturn) IsReturn() {
}

func (a *AstReturn) String() string {
	s := strings.Builder{}

	if a.value != nil {
		s.WriteString(a.value.String())
		s.WriteString("\n")
	}

	if a.function != nil {
		if a.function.user.sub {
			s.WriteString("retsub")
		} else {
			s.WriteString("return")
		}
	} else {
		s.WriteString("return")
	}

	return s.String()
}

type AstBlock struct {
	statements []AstStatement
}

func (a *AstBlock) String() string {
	ops := strings.Builder{}

	for _, stmt := range a.statements {
		ops.WriteString(stmt.String())
		ops.WriteString("\n")
	}

	return ops.String()
}

type AstIfAlternative struct {
	condition  AstStatement
	statements []AstStatement
}

type AstIf struct {
	label        int
	alternatives []*AstIfAlternative
}

func (a *AstIf) String() string {
	res := strings.Builder{}

	for i, alt := range a.alternatives {
		if alt.condition != nil {
			res.WriteString(alt.condition.String())
			res.WriteString("\n")

			if i < len(a.alternatives)-1 {
				res.WriteString(fmt.Sprintf("bz skip_%d_%d\n", a.label, i))
			} else {
				res.WriteString(fmt.Sprintf("bz end_%d\n", a.label))
			}
		}

		for _, stmt := range alt.statements {
			res.WriteString(stmt.String())
			res.WriteString("\n")
		}

		if i < len(a.alternatives)-1 {
			res.WriteString(fmt.Sprintf("b end_%d\n", a.label))
		}

		if alt.condition != nil {
			if i < len(a.alternatives)-1 {
				res.WriteString(fmt.Sprintf("skip_%d_%d:\n", a.label, i))
			}
		}
	}

	res.WriteString(fmt.Sprintf("end_%d:", a.label))

	return res.String()
}

type AstFunction struct {
	fun        *Function
	statements []AstStatement
}

func (a *AstFunction) String() string {
	res := strings.Builder{}

	res.WriteString(fmt.Sprintf("// function %s\n", a.fun.name))
	res.WriteString(fmt.Sprintf("%s:\n", a.fun.name))

	if a.fun.user.sub {
		ast := avm_proto_Ast{
			i1: itoa(a.fun.user.args),
			i2: itoa(a.fun.user.returns),
		}

		res.WriteString(ast.String())
		res.WriteString("\n")
	}

	for _, stmt := range a.statements {
		res.WriteString(stmt.String())
		res.WriteString("\n")
	}

	if a.fun.user.sub {
		if len(a.statements) > 0 {
			last := a.statements[len(a.statements)-1]
			if _, ok := last.(AstIsReturn); !ok {
				res.WriteString("retsub")
			}
		}
	}

	return res.String()
}

type AstProgram struct {
	functions map[string]*AstFunction
}

func (a *AstProgram) String() string {
	res := strings.Builder{}

	res.WriteString(fmt.Sprintf("#pragma version %d\n", AvmVersion))

	main := a.functions[AvmMainName]

	if len(a.functions) > 1 {
		res.WriteString(fmt.Sprintf("b %s\n", main.fun.name))
	}

	for name, fun := range a.functions {
		if name == AvmMainName {
			continue
		}

		res.WriteString(fun.String())
		res.WriteString("\n")
	}

	res.WriteString(main.String())

	return res.String()
}

func (v *AstVisitor) visitStatement(tree antlr.ParseTree) AstStatement {
	res := v.Visit(tree)
	if res == nil {
		return nil
	}

	return res.(AstStatement)
}

func (v *AstVisitor) VisitMemberExpr(ctx *parser.MemberExprContext) interface{} {
	vr, f := v.mustResolve(ctx.AllID())
	fun := v.global.functions[f.fun]

	ast := &AstStructField{
		v:   vr,
		t:   v.scope.resolveType(vr.t),
		f:   f,
		fun: fun,
	}

	return ast
}

func (v *AstVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	ids := ctx.AllID()

	vr, f := v.mustResolve(ids)
	if vr.readonly {
		panic(fmt.Sprintf("variable '%s' is read only", vr.name))
	}

	if vr.readonly {
		panic(fmt.Sprintf("variable '%s' is read only", vr.name))
	}

	ast := &AstAssign{
		v:     vr,
		t:     v.scope.resolveType(vr.t),
		f:     f,
		value: v.visitStatement(ctx.Expr()),
	}

	return ast
}

func (v *AstVisitor) VisitDeclaration(ctx *parser.DeclarationContext) interface{} {
	id := ctx.Type_().ID().GetText()
	t := v.scope.resolveType(id)

	if t == nil {
		panic(fmt.Sprintf("type '%s' not found", id))
	}

	return nil
}

func (v *AstVisitor) VisitVariableExpr(ctx *parser.VariableExprContext) interface{} {
	id := ctx.ID().GetText()
	vr := v.mustResolveVariable(id)

	ast := &AstVariable{
		v: vr,
	}

	return ast
}

func (v *AstVisitor) VisitMinusExpr(ctx *parser.MinusExprContext) interface{} {
	return &AstMinusOp{
		value: v.visitStatement(ctx.Expr()),
	}
}

func (v *AstVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: ctx.Addsub().GetText(),
	}
}
func (v *AstVisitor) VisitMulDivExpr(ctx *parser.MulDivExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: ctx.Muldiv().GetText(),
	}
}

func (v *AstVisitor) VisitEqNeqExpr(ctx *parser.EqNeqExprContext) interface{} {
	exprs := ctx.AllExpr()
	return &AstBinop{
		l:  v.visitStatement(exprs[0]),
		r:  v.visitStatement(exprs[1]),
		op: ctx.Eqneq().GetText(),
	}
}

func (v *AstVisitor) VisitConstantExpr(ctx *parser.ConstantExprContext) interface{} {
	var res AstStatement

	if ctx.INT() != nil {
		res = &AstIntConstant{
			value: ctx.INT().GetText(),
		}
	}

	if ctx.STRING() != nil {
		res = &AstByteConstant{
			value: ctx.STRING().GetText(),
		}
	}

	return res
}

func (v *AstVisitor) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	return v.Visit(ctx.Call_expr())
}
func (v *AstVisitor) VisitCall(ctx *parser.CallContext) interface{} {
	return v.Visit(ctx.Call_expr())
}

func (v *AstVisitor) VisitCall_expr(ctx *parser.Call_exprContext) interface{} {
	ids := ctx.AllID()

	id := ids[0].GetText()

	imms := []AstStatement{}

	if len(ids) > 1 {
		vr := v.mustResolveVariable(id)
		t := v.scope.resolveType(vr.t)

		if t.simple != nil {
			panic("cannot call simple type")
		}

		if t.complex.builtin == nil {
			panic("calling struct function is not supported yet")
		}

		id = t.complex.fields[ids[1].GetText()].fun

		// TODO: currently supports just a single level of fields

		imms = append(imms, &AstRaw{
			value: ids[1].GetText(),
		})
	}

	fun := v.global.functions[id]

	ast := &AstCall{
		fun: fun,
	}

	for _, arg := range ctx.Args().AllExpr() {
		stmt := v.visitStatement(arg)
		ast.args = append(ast.args, stmt)
	}

	ast.args = append(ast.args, imms...)

	return ast
}

func (v *AstVisitor) VisitReturn(ctx *parser.ReturnContext) interface{} {
	ast := &AstReturn{
		function: v.scope.function,
	}

	if ctx.Expr() != nil {
		ast.value = v.visitStatement(ctx.Expr())
	}

	return ast
}

func (v *AstVisitor) VisitIf(ctx *parser.IfContext) interface{} {
	alts := []*AstIfAlternative{}

	ast := &AstIf{
		label: v.label,
	}

	v.label++

	alt := &AstIfAlternative{
		condition: v.visitStatement(ctx.Expr()),
	}

	for _, item := range ctx.AllStmt() {
		if stmt := v.visitStatement(item); stmt != nil {
			alt.statements = append(alt.statements, stmt)
		}
	}

	alts = append(alts, alt)

	for _, elif := range ctx.AllElseif() {
		alt := &AstIfAlternative{
			condition: v.visitStatement(elif.Expr()),
		}

		for _, item := range elif.AllStmt() {
			if stmt := v.visitStatement(item); stmt != nil {
				alt.statements = append(alt.statements, stmt)
			}
		}

		alts = append(alts, alt)
	}

	if ctx.Else_() != nil {
		alt := &AstIfAlternative{}

		for _, item := range ctx.Else_().AllStmt() {
			stmt := v.visitStatement(item)
			alt.statements = append(alt.statements, stmt)
		}

		alts = append(alts, alt)
	}

	ast.alternatives = alts

	return ast
}

func (v *AstVisitor) mustResolveVariable(name string) *Variable {
	current := v.scope

	for current != nil {
		if vr, ok := current.variables[name]; ok {
			return vr
		}

		current = current.exit()
	}

	panic(fmt.Sprintf("variable '%s' not found", name))
}

func (v *AstVisitor) mustResolve(ids []antlr.TerminalNode) (*Variable, *StructField) {
	vr := v.mustResolveVariable(ids[0].GetText())
	t := v.scope.resolveType(vr.t)

	if len(ids) == 1 {
		return vr, nil
	}

	if t.simple != nil {
		panic("cannot resolve simple type access")
	}

	if len(ids) > 2 {
		panic("multilevel member access is not supported yet")
	}

	var f *StructField

	nvr := vr

	for i := 1; i < len(ids); i++ {
		vr = nvr
		t = v.scope.resolveType(vr.t)
		id := ids[i].GetText()
		f = t.complex.fields[id]
		nvr = vr.fields[id]
	}

	return vr, f
}

func (v *AstVisitor) VisitDefinition(ctx *parser.DefinitionContext) interface{} {
	id := ctx.ID().GetText()
	vr := v.scope.variables[id]
	t := v.scope.resolveType(vr.t)

	ast := &AstAssign{
		v:     vr,
		t:     t,
		value: v.visitStatement(ctx.Expr()),
	}

	return ast
}

func (v *AstVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	id := ctx.ID().GetText()
	fun := v.global.functions[id]

	ast := &AstFunction{
		fun: fun,
	}

	v.scope = fun.user.scope

	for _, item := range ctx.AllStmt() {
		if stmt := v.visitStatement(item); stmt != nil {
			ast.statements = append(ast.statements, stmt)
		}
	}

	v.scope = v.scope.exit()

	v.program.functions[id] = ast

	return nil
}

func (v *AstVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.scope = v.scope.enter()

	ast := &AstBlock{}

	for _, item := range ctx.AllStmt() {
		stmt := v.visitStatement(item)
		if stmt != nil {
			ast.statements = append(ast.statements, stmt)
		}
	}

	v.scope = v.scope.exit()

	return ast
}

func (v *AstVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.scope = v.global
	v.VisitChildren(ctx)
	v.scope = nil

	return nil
}

func Compile(src string) string {
	input := antlr.NewInputStream(src)
	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCParser(stream)

	global := NewScope(nil)

	global.types["void"] = &Type{
		name: "void",
		simple: &SimpleType{
			empty: true,
		},
	}

	global.types["bytes"] = &Type{
		name:   "bytes",
		simple: &SimpleType{},
	}

	global.types["uint64"] = &Type{
		name:   "uint64",
		simple: &SimpleType{},
	}

	global.types["any"] = &Type{
		name:   "any",
		simple: &SimpleType{},
	}

	for _, item := range builtin_functions {
		f := &Function{
			t:    item.t,
			name: item.name,
			builtin: &BuiltinFunction{
				op: item.op,
			},
		}

		for _, item := range item.stack {
			f.builtin.stack = append(f.builtin.stack, &FunctionParam{
				t:    item.t,
				name: item.name,
			})
		}

		for _, item := range item.imm {
			f.builtin.imm = append(f.builtin.imm, &FunctionParam{
				t:    item.t,
				name: item.name,
			})
		}

		global.functions[f.name] = f
	}

	for _, item := range builtin_structs {
		s := &Struct{
			builtin:   &BuiltinStruct{},
			fields:    map[string]*StructField{},
			functions: map[string]*StructFunction{},
		}

		for _, item := range item.fields {
			s.fields[item.name] = &StructField{
				t:    item.t,
				name: item.name,
				fun:  item.fun,
			}
		}

		for _, item := range item.functions {
			f := &StructFunction{
				t:    item.t,
				name: item.name,
			}

			for _, item := range item.params {
				f.params = append(f.params, &FunctionParam{
					t:    item.t,
					name: item.name,
				})
			}

			s.functions[f.name] = f
		}

		t := &Type{
			name:    item.name,
			complex: s,
		}

		global.types[t.name] = t
	}

	for _, item := range builtin_variables {
		v := &Variable{
			t:    item.t,
			name: item.name,
		}

		global.variables[v.name] = v
	}

	sv := &SymbolTableVisitor{
		BaseCVisitor: &parser.BaseCVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		global: global,
	}

	stream.Seek(0)
	sv.Visit(p.Program())

	program := &AstProgram{
		functions: map[string]*AstFunction{},
	}

	global.readonly()

	av := &AstVisitor{
		BaseCVisitor: &parser.BaseCVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		global:  global,
		program: program,
	}

	stream.Seek(0)
	av.Visit(p.Program())

	return program.String()
}