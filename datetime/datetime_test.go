package datetime_test

import (
	. "github.com/1ambda/github-archive-downloader/datetime"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Datetime", func() {
	Describe("IsValidTimeString", func() {
		It("should return true given time RFC3339 style string", func() {
			_, err := ParseRawTime("2016-11-07T10")
			Expect(err).Should(BeNil())
		})

		It("should return false given invalid datetime string", func() {
			_, err := ParseRawTime("2016-11-07-10")
			Expect(err).ShouldNot(BeNil())
		})
	})
})
