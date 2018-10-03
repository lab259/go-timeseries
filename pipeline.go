package timesrs

type PipelineResult struct {
	Data AggregationData
}

// PipelineTagFunction receives an `Event` and set all keys beside the timeline
// which will, automatically, be set as _id.
type PipelineTagFnc func(e Event) (map[string]interface{}, error)

// NoTags define that the aggregation will have no tags but the time.
func NoTags(e Event) (map[string]interface{}, error) {
	return nil, nil
}

// TagByField implements a `PipelineKeyFunction` that will only use field names
// (through reflection) to get the data value. It is a helper function for
// sugar coding.
func TagByField(fields ...string) PipelineTagFnc {
	return func(e Event) (map[string]interface{}, error) {
		result := make(map[string]interface{}, len(fields))
		for _, field := range fields {
			v, err := GetFieldValue(field, e.Data())
			if err != nil {
				return nil, err
			}
			result[field] = v
		}
		return result, nil
	}
}

// PipelineErrorHandler is called when an error is triggered in the pipeline
// process.
type PipelineErrorHandler func(aggregation Aggregation, err error) error

// Pipeline is responsible to receive an event and trigger each `Bucket` in
// sequence, aggregating each result and storing it using the `Storage`.
//
// A Pipeline keeps an array of buckets that will be triggered when `Run` method
// is called.
type Pipeline struct {
	// Granularities is the set of granularity configurations that will be
	// stored in the `Storage` set in this `Pipeline`.
	Granularities []GranularityComposite

	// Aggregations is the set of aggregations that will be applied in this
	// `Pipeline`.
	Aggregations []Aggregation

	// Collection is the base name of the
	Collection string

	// TagFnc will evaluate the `Event` and recreate
	TagFnc PipelineTagFnc

	// Storage is the storaging system used in this `Pipeline`.
	Storage Storage

	// ErrorHandler is the handling error function that will be called if any
	// error is triggered in the pipeline process.
	ErrorHandler PipelineErrorHandler
}

// Run triggers the processing of the buckets in order of registering.
//
// Each `BucketResult` is synchronously stored in the `Storage` specified at the
// `Pipeline` "constructor".
//
// The param `e` informed is the source of the data that will go through the
// pipeline.
func (p *Pipeline) Run(e Event) (*PipelineResult, error) {
	var pResult PipelineResult
	aggrResult := make(AggregationData, 0)
	for _, aggr := range p.Aggregations {
		err := aggr.Aggregate(e, &aggrResult)
		if err != nil {
			if p.ErrorHandler == nil {
				return nil, err
			}
			err = p.ErrorHandler(aggr, err)
			if err != nil {
				return nil, err
			}
		}
	}
	keys, err := p.TagFnc(e)
	pResult.Data = aggrResult
	if err != nil {
		return nil, err
	}
	err = p.Storage.Store(&StorageEntry{
		Event:         e,
		Keys:          keys,
		Collection:    p.Collection,
		Granularities: p.Granularities,
		Result:        &pResult,
	})
	if err != nil {
		return nil, err
	}
	return &pResult, nil
}
