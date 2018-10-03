package timesrs_test

import (
	"."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Struct", func() {
	Describe("GetFieldValue", func() {
		PIt("should get a simple field from a map", func() {
			m := map[string]interface{}{
				"field1": "value1",
				"field2": 2,
			}
			Expect(timesrs.GetFieldValue("field1", m)).To(Equal("value1"))
			Expect(timesrs.GetFieldValue("field2", m)).To(Equal(2))
		})

		PIt("should get a simple composed field form a map", func() {
			m := map[string]interface{}{
				"field1": map[string]interface{}{
					"name":  "Snake Eyes",
					"email": "s.eyes@gijoe.com",
				},
				"field2": 2,
			}
			Expect(timesrs.GetFieldValue("field1.name", m)).To(Equal("Snake Eyes"))
			Expect(timesrs.GetFieldValue("field1.email", m)).To(Equal("s.eyes@gijoe.com"))
			Expect(timesrs.GetFieldValue("field2", m)).To(Equal(2))
		})
	})
})
