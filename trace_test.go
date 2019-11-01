package mic

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trace", func() {
	DescribeTable("normalize",
		func(src string, expected string) {
			Expect(normalize(src)).To(Equal(expected))
		},
		Entry("id at tail", "/foo/1", "/foo/:num"),
		Entry("id at middle", "/foo/1/bar", "/foo/:num/bar"),
		Entry("id with param", "/foo/1?tick=xxx", "/foo/:num"),
		Entry("id 3", "/foo/123455abc", "/foo/123455abc"),
		Entry("hash", "/foo/57e48711e4d0f314382b45c9/bar", "/foo/:hash/bar"),
		Entry("crop at tail", "/foo/100x200", "/foo/:crop"),
		Entry("all", "/f/57e48711e4d0f314382b45c9/20/10x20", "/f/:hash/:num/:crop"),
	)
})
