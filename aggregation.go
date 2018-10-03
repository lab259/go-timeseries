package timesrs

type AggregationData map[string]interface{}

// Aggregator produces an operation.
type Aggregation interface {
	Aggregate(event Event, data AggregationData) (error)
}
