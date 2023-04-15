package compiler

import (
	"fmt"
	"strings"
)

type TealOp interface {
	String() string
}

type Teal struct {
	Ops []TealOp
}

func (t *Teal) Write(op TealOp) {
	t.Ops = append(t.Ops, op)
}

func (t *Teal) String() string {
	res := Lines{}
	for _, op := range t.Ops {
		res.WriteLine(op.String())
	}
	return res.String()
}

type Teal_pragma_version struct {
	Version int
}

func (t *Teal_pragma_version) String() string {
	return fmt.Sprintf("#pragma version %d", t.Version)
}

type Teal_intcblock_fixed struct {
	UINT1 []uint64
}

func (a *Teal_intcblock_fixed) String() string {
	res := strings.Builder{}
	res.WriteString("intcblock")

	for _, v := range a.UINT1 {
		res.WriteString(" ")
		res.WriteString(fmt.Sprintf("%d", v))
	}

	return res.String()
}

type Teal_bytecblock_fixed struct {
	BYTES1 [][]byte
}

func (a *Teal_bytecblock_fixed) String() string {
	res := strings.Builder{}
	res.WriteString("bytecblock")

	for _, v := range a.BYTES1 {
		res.WriteString(" ")
		// TODO: may need other formatting than string for non-printable bytes
		res.WriteString(string(v))
	}

	return res.String()
}

type Teal_b_fixed struct {
	TARGET1 string
}

func (a *Teal_b_fixed) String() string {
	return fmt.Sprintf("b %s", a.TARGET1)
}

type Teal_bnz_fixed struct {
	s1      AstStatement
	TARGET1 string
}

func (a *Teal_bnz_fixed) String() string {
	return fmt.Sprintf("%s\nbnz %s", a.s1.String(), a.TARGET1)
}

type Teal_bz_fixed struct {
	s1      AstStatement
	TARGET1 string
}

func (a *Teal_bz_fixed) String() string {
	return fmt.Sprintf("%s\nbz %s", a.s1.String(), a.TARGET1)
}

type Teal_label struct {
	Name string
}

func (a *Teal_label) String() string {
	return fmt.Sprintf("%s:", a.Name)
}

type Teal_match_fixed struct {
	TARGET1 []string
}

func (a *Teal_match_fixed) String() string {
	res := strings.Builder{}
	res.WriteString("match")

	for _, v := range a.TARGET1 {
		res.WriteString(" ")
		res.WriteString(v)
	}

	return res.String()
}

type Teal_int struct {
	V uint64
}

func (a *Teal_int) String() string {
	return fmt.Sprintf("int %d", a.V)
}

type Teal_seq struct {
	Ops []TealOp
}

func (a *Teal_seq) String() string {
	res := Teal{}
	for _, op := range a.Ops {
		res.Write(op)
	}
	return res.String()
}
