package object

import "fmt"


type ObjType string

const(
    INTEGER_OBJ = "INTIGET_TYPE"
    BOOLEAN_OBJ = "BOOLEAN"
)

type Object interface{
    Type() ObjType
    Inspect() string
}


type Integer struct{
    Value int
}

func (i *Integer) Type() ObjType{
    return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
    return fmt.Sprintf("%d", i.Value)
}


type Boolean struct{
    Value bool
}

func (b *Boolean) Inspect() string{
    return fmt.Sprintf("%v", b.Value)
}
func (b *Boolean) Type() ObjType {
    return BOOLEAN_OBJ
}




























