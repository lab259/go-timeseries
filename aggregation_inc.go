package timesrs

// AggregationFnc increases a value of a field
type AggregationFnc struct {
	fnc func(e Event, data *AggregationData) error
}

// NewAggregationFnc creates a new instance of `AggregationFnc`.
func NewAggregationFnc(fnc func(e Event, data *AggregationData) error) *AggregationFnc {
	return &AggregationFnc{
		fnc: fnc,
	}
}

// Aggregate calls fnc and returns its error
func (agg *AggregationFnc) Aggregate(event Event, data *AggregationData) error {
	return agg.fnc(event, data)
}
