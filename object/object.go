package object

import (
	"bytes"
	"fmt"
	"javascript_interpreter/ast"
	"strings"
	"writingincompiler/code"
)

type ObjectType string

const (
	INTEGER_OBJ           = "INTEGER"
	BOOLEAN_OBJ           = "BOOLEAN"
	NULL_OBJ              = "NULL"
	RETURN_VALUE_OBJ      = "RETURN_VALUE"
	ERROR_OBJ             = "ERROR"
	FUNCTION_OBJ          = "FUNCTION"
	ARRAY_OBJ             = "ARRAY_OBJ"
	HASH_OBJ              = "HASH_OBJ"
	STRING_OBJ            = "STRING_OBJ"
	COMPILED_FUNCTION_OBJ = "COMPILED_FUNCTION_OBJ"
)

type CompiledFunction struct {
	Instructions code.Instructions
}

func (cf *CompiledFunction) Type() ObjectType { return COMPILED_FUNCTION_OBJ }

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Hash struct {
	Hash map[string]Object
}

func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

func (h *Hash) Inspect() string {
	var out bytes.Buffer
	element := []string{}

	// for k, v := range h.Hash {
	// 	eleString := k.(string) + ":" + v.(string)
	// 	element = append(element, eleString)
	// }

	out.WriteString("{")
	out.WriteString(strings.Join(element, ", "))
	out.WriteString("}")

	return out.String()
}

type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType {
	return ARRAY_OBJ
}

func (a *Array) Inspect() string {

	var out bytes.Buffer
	element := []string{}

	for _, p := range a.Elements {
		element = append(element, p.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(element, ", "))
	out.WriteString("]")

	return out.String()

}

type Number struct {
	Value float64
}

type Boolean struct {
	Value bool
}

func (i *Number) Inspect() string {
	return fmt.Sprintf("%v", i.Value)
}

func (i *Number) Type() ObjectType { return INTEGER_OBJ }

type String struct {
	Value string
}

func (i *String) Inspect() string {
	return fmt.Sprintf("%s", i.Value)
}

func (i *String) Type() ObjectType { return STRING_OBJ }

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())

	return out.String()
}
