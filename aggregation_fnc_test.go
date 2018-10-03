package timesrs_test

import (
	"."
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aggregation", func() {
	Describe("Fnc", func() {
		It("should initialize the aggregator with default value", func() {
			agg := timesrs.NewAggregationFnc(func(e timesrs.Event, data *timesrs.AggregationData) error {
				return nil
			})
			Expect(agg).NotTo(BeNil())
		})

		It("should aggregate a field", func() {
			agg := timesrs.NewAggregationFnc(func(e timesrs.Event, data *timesrs.AggregationData) error {
				*data = append(*data, timesrs.Operation{
					Type:  timesrs.OperationTypeInc,
					Field: "field1",
					Value: 1,
				})
				return nil
			})
			data := make(timesrs.AggregationData, 0)
			Expect(agg.Aggregate(timesrs.NewEvent("name", nil, timesrs.DefaultClock.Time()), &data)).To(BeNil())
			Expect(data).To(HaveLen(1))
			Expect(data[0].Type).To(Equal(timesrs.OperationTypeInc))
			Expect(data[0].Field).To(Equal("field1"))
			Expect(data[0].Value).To(Equal(1))
		})

		It("should fail due to aggregation error", func() {
			agg := timesrs.NewAggregationFnc(func(e timesrs.Event, data *timesrs.AggregationData) error {
				return errors.New("forced error")
			})
			err := agg.Aggregate(timesrs.NewEvent("name", nil, timesrs.DefaultClock.Time()), nil)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("forced error"))
		})
	})
})
