package timesrs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"

	"."
)

var _ = Describe("Clock", func() {
	Describe("UTCClock", func() {
		It("should get the current UTC time", func() {
			Expect(timesrs.DefaultClock.Time()).To(BeTemporally("~", time.Now().UTC(), time.Millisecond))
		})
	})

	Describe("LocalClock", func() {
		It("should get the current local time", func() {
			var clock timesrs.LocalClock
			Expect(clock.Time()).To(BeTemporally("~", time.Now(), time.Millisecond))
		})
	})
})