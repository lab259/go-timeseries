package timesrs_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"."
)

var _ = Describe("Pipeline", func() {
	It("should aggregate one field", func() {
		evt := timesrs.NewEvent("event1", nil, timesrs.DefaultClock.Time())

		pipeline := timesrs.Pipeline{
			Collection: "aggregation1",
			Aggregations: []timesrs.Aggregation{
				timesrs.NewAggregationInc("field1"),
			},
			Granularities: []timesrs.GranularityComposite{
				{
					Collection: timesrs.GranularityHour,
					Record:     timesrs.GranularityDay,
				},
			},
			Storage: NewStorageMock(func(entry *timesrs.StorageEntry) error {
				Expect(entry.Event).To(Equal(evt))
				Expect(entry.Keys).To(BeEmpty())
				Expect(entry.Result.Data).To(HaveLen(1))
				Expect(entry.Result.Data).To(HaveKeyWithValue("field1", 1))
				return nil
			}),
			TagFnc: timesrs.NoTags,
		}
		_, err := pipeline.Run(evt)
		Expect(err).To(BeNil())
	})

	It("should aggregate multiple field", func() {
		evt := timesrs.NewEvent("event1", nil, timesrs.DefaultClock.Time())

		pipeline := timesrs.Pipeline{
			Collection: "aggregation1",
			Aggregations: []timesrs.Aggregation{
				timesrs.NewAggregationInc("field1"),
				timesrs.NewAggregationInc("field2", 2),
			},
			Granularities: []timesrs.GranularityComposite{
				{
					Collection: timesrs.GranularityHour,
					Record:     timesrs.GranularityDay,
				},
			},
			Storage: NewStorageMock(func(entry *timesrs.StorageEntry) error {
				Expect(entry.Event).To(Equal(evt))
				Expect(entry.Keys).To(BeEmpty())
				Expect(entry.Result.Data).To(HaveLen(2))
				Expect(entry.Result.Data).To(HaveKeyWithValue("field1", 1))
				Expect(entry.Result.Data).To(HaveKeyWithValue("field2", 2))

				return nil
			}),
			TagFnc: timesrs.NoTags,
		}
		_, err := pipeline.Run(evt)
		Expect(err).To(BeNil())
	})

	It("should aggregate multiple field", func() {
		evt := timesrs.NewEvent("event1", nil, timesrs.DefaultClock.Time())

		pipeline := timesrs.Pipeline{
			Collection: "aggregation1",
			Aggregations: []timesrs.Aggregation{
				timesrs.NewAggregationInc("field1"),
				timesrs.NewAggregationInc("field2", 2),
			},
			Granularities: []timesrs.GranularityComposite{
				{
					Collection: timesrs.GranularityHour,
					Record:     timesrs.GranularityDay,
				},
			},
			Storage: NewStorageMock(func(entry *timesrs.StorageEntry) error {
				Expect(entry.Event).To(Equal(evt))
				Expect(entry.Keys).To(BeEmpty())
				Expect(entry.Result.Data).To(HaveLen(2))
				Expect(entry.Result.Data).To(HaveKeyWithValue("field1", 1))
				Expect(entry.Result.Data).To(HaveKeyWithValue("field2", 2))

				return nil
			}),
			TagFnc: timesrs.NoTags,
		}
		_, err := pipeline.Run(evt)
		Expect(err).To(BeNil())
	})

	It("should fail due to aggregation error with no `ErrorHandler`", func() {
		evt := timesrs.NewEvent("event1", nil, timesrs.DefaultClock.Time())

		pipeline := timesrs.Pipeline{
			Collection: "aggregation1",
			Aggregations: []timesrs.Aggregation{
				timesrs.NewAggregationFnc(func(e timesrs.Event, data timesrs.AggregationData) error {
					return errors.New("aggregation error")
				}),
			},
			Granularities: []timesrs.GranularityComposite{
				{
					Collection: timesrs.GranularityHour,
					Record:     timesrs.GranularityDay,
				},
			},
			Storage: NewStorageMock(func(entry *timesrs.StorageEntry) error {
				Fail("storage should not be reached")
				return nil
			}),
			TagFnc: timesrs.NoTags,
		}
		_, err := pipeline.Run(evt)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("aggregation error"))
	})

	It("should fail due to aggregation error with `ErrorHandler` enhancing the error", func() {
		evt := timesrs.NewEvent("event1", nil, timesrs.DefaultClock.Time())

		pipeline := timesrs.Pipeline{
			Collection: "aggregation1",
			Aggregations: []timesrs.Aggregation{
				timesrs.NewAggregationFnc(func(e timesrs.Event, data timesrs.AggregationData) error {
					return errors.New("aggregation error")
				}),
			},
			Granularities: []timesrs.GranularityComposite{
				{
					Collection: timesrs.GranularityHour,
					Record:     timesrs.GranularityDay,
				},
			},
			Storage: NewStorageMock(func(entry *timesrs.StorageEntry) error {
				Fail("storage should not be reached")
				return nil
			}),
			TagFnc: timesrs.NoTags,
			ErrorHandler: func(aggregation timesrs.Aggregation, err error) error {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("aggregation error"))
				return errors.New("enhanced aggregation error")
			},
		}
		_, err := pipeline.Run(evt)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("enhanced aggregation error"))
	})

	It("should fail due to aggregation error with `ErrorHandler` discarding the error", func() {
		evt := timesrs.NewEvent("event1", nil, timesrs.DefaultClock.Time())

		pipeline := timesrs.Pipeline{
			Collection: "aggregation1",
			Aggregations: []timesrs.Aggregation{
				timesrs.NewAggregationInc("field1"),
				timesrs.NewAggregationFnc(func(e timesrs.Event, data timesrs.AggregationData) error {
					return errors.New("aggregation error")
				}),
				timesrs.NewAggregationInc("field3", 3),
			},
			Granularities: []timesrs.GranularityComposite{
				{
					Collection: timesrs.GranularityHour,
					Record:     timesrs.GranularityDay,
				},
			},
			Storage: NewStorageMock(func(entry *timesrs.StorageEntry) error {
				Expect(entry.Result.Data).To(HaveLen(2))
				Expect(entry.Result.Data).To(HaveKeyWithValue("field1", 1))
				Expect(entry.Result.Data).To(HaveKeyWithValue("field3", 3))
				return nil
			}),
			ErrorHandler: func(aggregation timesrs.Aggregation, err error) error {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("aggregation error"))
				return nil
			},
			TagFnc: timesrs.NoTags,
		}
		_, err := pipeline.Run(evt)
		Expect(err).To(BeNil())
	})

	It("should fail due to TagFnc error", func() {
		evt := timesrs.NewEvent("event1", nil, timesrs.DefaultClock.Time())

		pipeline := timesrs.Pipeline{
			Collection: "aggregation1",
			Aggregations: []timesrs.Aggregation{
				timesrs.NewAggregationInc("field1"),
			},
			Granularities: []timesrs.GranularityComposite{
				{
					Collection: timesrs.GranularityHour,
					Record:     timesrs.GranularityDay,
				},
			},
			Storage: NewStorageMock(func(entry *timesrs.StorageEntry) error {
				Expect(entry.Result.Data).To(HaveLen(1))
				Expect(entry.Result.Data).To(HaveKeyWithValue("field1", 1))
				return nil
			}),
			TagFnc: func(e timesrs.Event) (map[string]interface{}, error) {
				return nil, errors.New("TagFnc error")
			},
		}
		_, err := pipeline.Run(evt)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("TagFnc error"))
	})
})
