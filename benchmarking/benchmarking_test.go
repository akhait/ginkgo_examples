package benchmarking_test

import (
	. "ginkgo_examples/benchmarking"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("Benchmarking", func() {
	Describe("Reverse small data", func() {
		Describe("Short strings", func() {
			Context("run 1 time", func(){
				Measure("should be reversed fast", func(b Benchmarker) {
					runtime := b.Time("runtime", func(){
						data := "shortstring"
						ReverseOrder(string(data))
					})
					Ω(runtime.Seconds()).Should(BeNumerically("<", 0.01), "should not take more than 10 ms")
				}, 1)
			})
		})
	})
	Describe("Reverse long data", func(){
		Describe("Long text", func(){
			Context("run 1 time", func(){
				Measure("should be reversed fast", func(b Benchmarker) {
					runtime := b.Time("runtime", func(){
						data, err := ioutil.ReadFile("/usr/share/dict/words")
						Expect(err).NotTo(HaveOccurred())
						ReverseOrder(string(data))
					})
					Ω(runtime.Seconds()).Should(BeNumerically("<", 0.03), "should not take more than 20 ms")
				}, 1)
			})
		})
	})
})
