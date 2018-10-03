package timesrs_test

import (
	"."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aggregation", func() {
	Describe("Inc", func() {
		It("should initialize the aggregator with default value", func() {
			agg := timesrs.NewAggregationInc("field1")
			Expect(agg).NotTo(BeNil())
			Expect(agg.Field).To(Equal("field1"))
			Expect(agg.Value).To(Equal(1))
		})

		It("should aggregate one field with specified value", func() {
			agg := timesrs.NewAggregationInc("field1", 5)
			Expect(agg).NotTo(BeNil())
			Expect(agg.Field).To(Equal("field1"))
			Expect(agg.Value).To(Equal(5))
		})

		It("should aggregate a field with default value", func() {
			agg := timesrs.NewAggregationInc("field1")
			data := make(timesrs.AggregationData, 0)
			Expect(agg.Aggregate(timesrs.NewEvent("name", nil, timesrs.DefaultClock.Time()), &data)).To(BeNil())
			Expect(data).To(HaveLen(1))
			Expect(data[0].Field).To(Equal("field1"))
			Expect(data[0].Value).To(Equal(1))
		})
	})
})
