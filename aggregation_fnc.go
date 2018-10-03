package timesrs

// AggregatorInc increases a value of a field
type AggregationInc struct {
	Field string
	Value int
}

// NewAggregationInc creates a new instnace of the aggregation inc.
func NewAggregationInc(field string, value ...int) *AggregationInc {
	if len(value) > 0 {
		return &AggregationInc{
			Field: field,
			Value: value[0],
		}
	}
	return &AggregationInc{
		Field: field,
		Value: 1,
	}
}

// Aggregate adds the `Field` with its `Value` to the `data`.
func (agg *AggregationInc) Aggregate(event Event, data AggregationData) error {
	data[agg.Field] = agg.Value
	return nil
}
