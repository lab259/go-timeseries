package timesrs

type OperationType = int

const (
	OperationTypeInc = iota
	OperationTypeSet
)

type Operation struct {
	Type  OperationType
	Field string
	Value interface{}
}
