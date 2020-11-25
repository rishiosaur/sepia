package objects

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"sepia/ast"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	MAP_OBJ          = "MAP"
)

type BuiltinFunc func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunc
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Machine struct {
	store map[string]Object
	outer *Machine
}

func NewMachine() *Machine {
	s := make(map[string]Object)
	return &Machine{store: s, outer: nil}
}

func NewLocalMachine(outer *Machine) *Machine {
	env := NewMachine()
	env.outer = outer
	return env
}

func (e *Machine) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok := e.outer.Get(name)
		return obj, ok
	}
	return obj, ok
}

func (e *Machine) Set(name string, val Object) Object {

	e.store[name] = val
	return val

}

func (e *Machine) Update(name string, val Object) Object {
	_, ok := e.store[name]

	if !ok && e.outer != nil {
		e.outer.Update(name, val)
	} else if !ok && e.outer == nil {
		return &Error{Message: "Could not find identitier `" + name + "` in program."}
	} else {
		e.store[name] = val
	}

	return val

}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Machine    *Machine
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type MapKey struct {
	Type  ObjectType
	Value uint
}

type Mappable interface {
	MapKey() MapKey
}

func (b *Boolean) MapKey() MapKey {
	var value uint

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return MapKey{Type: b.Type(), Value: value}
}

func (i *Integer) MapKey() MapKey {
	return MapKey{Type: i.Type(), Value: uint(i.Value)}
}

func (s *String) MapKey() MapKey {
	h := fnv.New64a()
	_, ok := h.Write([]byte(s.Value))
	if ok != nil {
		panic("could not hash string `" + s.Value + "`")
	}
	return MapKey{Type: s.Type(), Value: uint(h.Sum64())}
}

type MapPair struct {
	Key   Object
	Value Object
}

type Map struct {
	Pairs map[MapKey]MapPair
}

func (h *Map) Type() ObjectType { return MAP_OBJ }

func (h *Map) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
