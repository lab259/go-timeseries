package timesrs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"."
)

var _ = Describe("Operation", func() {
	It("should return a new operation inc", func() {
		operation := timesrs.NewOperationInc("field", 1)
		Expect(operation.Type).To(Equal(timesrs.OperationTypeInc))
		Expect(operation.Field).To(Equal("field"))
		Expect(operation.Value).To(Equal(1))
	})

	It("should return a new operation set", func() {
		operation := timesrs.NewOperationSet("field", 1)
		Expect(operation.Type).To(Equal(timesrs.OperationTypeSet))
		Expect(operation.Field).To(Equal("field"))
		Expect(operation.Value).To(Equal(1))
	})
})
