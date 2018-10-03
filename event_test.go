package timesrs_test

import (
	"."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Event", func() {
	Describe("NewEvent", func() {
		It("should create a new Event", func() {
			event := timesrs.NewEvent("event1", "data1", time.Now().UTC())
			Expect(event).NotTo(BeNil())
			Expect(event.Name()).To(Equal("event1"))
			Expect(event.Data()).To(Equal("data1"))
			Expect(event.Time()).To(BeTemporally("~", time.Now().UTC(), time.Second))
		})
	})
})
