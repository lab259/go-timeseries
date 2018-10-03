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

func NewOperationInc(field string, value interface{}) *Operation {
	return &Operation{
		Type:  OperationTypeInc,
		Field: field,
		Value: value,
	}
}

func NewOperationSet(field string, value interface{}) *Operation {
	return &Operation{
		Type:  OperationTypeSet,
		Field: field,
		Value: value,
	}
}
