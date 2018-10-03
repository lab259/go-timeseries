package timesrs

type AggregationData []Operation

// Aggregator produces an operation.
type Aggregation interface {
	Aggregate(event Event, data *AggregationData) (error)
}
