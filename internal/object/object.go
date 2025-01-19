package object

import "fmt"

type ObjType string

const (
	INTEGER_OBJ = "INTIGER_TYPE"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL        = "NULL"
    RETURN_VALUE = "RETURN_VALUE"
    ERROR_OBJ = "ERROR"
)

type Object interface {
	Type() ObjType
	Inspect() string
}

type Integer struct {
	Value int
}

func (i *Integer) Type() ObjType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%v", b.Value)
}
func (b *Boolean) Type() ObjType {
	return BOOLEAN_OBJ
}

type Null struct{}

func (n *Null) Type() ObjType {
	return NULL
}
func (n *Null) Inspect() string {
	return "NULL"
}


type ReturnValue struct{
    Value Object
}

func (r *ReturnValue) Type() ObjType{
    return RETURN_VALUE
}
func (r *ReturnValue) Inspect() string{
    return fmt.Sprintf("%s", r.Value.Inspect())
}


type Error struct{
    Message string
}

func (e *Error) Type() ObjType{
    return ERROR_OBJ
}
func (e *Error) Inspect() string{
    return e.Message
}


