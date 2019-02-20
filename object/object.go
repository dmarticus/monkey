package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJECT = "INTEGER"
	BOOLEAN_OBJECT = "BOOLEAN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Boolean struct {
	Value bool
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJECT
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJECT
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%d", b.Value)
}
