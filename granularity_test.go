package timesrs_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"

	"."
)

var _ = Describe("Granularity", func() {
	Describe("String", func() {
		It("should get second", func() {
			Expect(timesrs.GranularitySecond.String()).To(Equal("second"))
		})

		It("should get minute", func() {
			Expect(timesrs.GranularityMinute.String()).To(Equal("minute"))
		})

		It("should get hour", func() {
			Expect(timesrs.GranularityHour.String()).To(Equal("hour"))
		})

		It("should get day", func() {
			Expect(timesrs.GranularityDay.String()).To(Equal("day"))
		})

		It("should get week", func() {
			Expect(timesrs.GranularityWeek.String()).To(Equal("week"))
		})

		It("should get month", func() {
			Expect(timesrs.GranularityMonth.String()).To(Equal("month"))
		})

		It("should get empty text when dealing with unknown granularity definition", func() {
			Expect(timesrs.Granularity(1234).String()).To(Equal(""))
		})
	})

	Describe("RelativeTo", func() {
		It("should get relative seconds", func() {
			Expect(timesrs.GranularitySecond.RelativeTo(timesrs.DefaultClock.Time(), timesrs.DefaultClock.Time().Add(5*time.Second))).To(Equal(5))
		})

		It("should get relative minutes", func() {
			Expect(timesrs.GranularityMinute.RelativeTo(timesrs.DefaultClock.Time(), timesrs.DefaultClock.Time().Add(5*time.Minute))).To(Equal(5))
		})

		It("should get relative hours", func() {
			Expect(timesrs.GranularityHour.RelativeTo(timesrs.DefaultClock.Time(), timesrs.DefaultClock.Time().Add(5*time.Hour))).To(Equal(5))
		})

		It("should get relative day", func() {
			dt := time.Date(2018, time.October, 1, 0, 0, 0, 0, time.UTC)
			fmt.Println(dt.String())
			dtRelative := time.Date(2018, time.October, 3, 13, 18, 2, 12312, time.UTC)
			Expect(timesrs.GranularityDay.RelativeTo(dt, dtRelative)).To(Equal(2))
		})

		It("should get relative weeks", func() {
			Expect(timesrs.GranularityWeek.RelativeTo(timesrs.DefaultClock.Time(), timesrs.DefaultClock.Time().Add(7*5*24*time.Hour))).To(Equal(5))
		})

		It("should get relative months", func() {
			Expect(timesrs.GranularityMonth.RelativeTo(timesrs.DefaultClock.Time(), timesrs.DefaultClock.Time())).To(Equal(int(timesrs.DefaultClock.Time().Month())))
		})

		It("should fail getting a relative when invalid granularity", func() {
			Expect(timesrs.Granularity(1234).RelativeTo(timesrs.DefaultClock.Time(), timesrs.DefaultClock.Time())).To(Equal(-1))
		})
	})

	Describe("Truncate", func() {
		var dt = time.Date(2018, 10, 1, 2, 3, 4, 5, time.UTC)

		It("should truncate seconds", func() {
			Expect(timesrs.GranularitySecond.Truncate(dt)).To(Equal(time.Date(2018, 10, 1, 2, 3, 4, 0, time.UTC)))
		})

		It("should truncate minutes", func() {
			Expect(timesrs.GranularityMinute.Truncate(dt)).To(Equal(time.Date(2018, 10, 1, 2, 3, 0, 0, time.UTC)))
		})

		It("should truncate hours", func() {
			Expect(timesrs.GranularityHour.Truncate(dt)).To(Equal(time.Date(2018, 10, 1, 2, 0, 0, 0, time.UTC)))
		})

		It("should truncate day", func() {
			Expect(timesrs.GranularityDay.Truncate(dt)).To(Equal(time.Date(2018, 10, 1, 0, 0, 0, 0, time.UTC)))
		})

		It("should truncate weeks", func() {
			Expect(timesrs.GranularityWeek.Truncate(dt)).To(Equal(time.Date(2018, 9, 30, 0, 0, 0, 0, time.UTC)))
		})

		It("should truncate months", func() {
			Expect(timesrs.GranularityMonth.Truncate(dt)).To(Equal(time.Date(2018, 10, 1, 0, 0, 0, 0, time.UTC)))
		})

		It("should fail truncating a unknown granularity", func() {
			Expect(timesrs.Granularity(1234).Truncate(dt).IsZero()).To(BeTrue())
		})
	})
})
