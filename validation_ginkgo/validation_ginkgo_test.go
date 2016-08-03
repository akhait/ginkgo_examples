package validation_ginkgo_test

import (
	"ginkgo_examples/validation_ginkgo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

type test_data struct {
	username string
	greeting string
}

var tests = []test_data{
			{"12345678", "Hello, 12345678"},
			{"SuperUserName","Hello, SuperUserName"},
			{"Test This One", "Hello, Test This One"},
		}

//Top-level container
var _ = Describe("ValidationGinkgo", func() {
	Describe("Positive tests - no errors should occure", func(){
		for _, test_case := range tests {
			greet := test_case.greeting
			greet, err := validation_ginkgo.RegisterNewUsername(test_case.username)
			Describe("Username", func(){
				Context("longer than 8 chars and no restricted chars", func(){
					It("should be legal", func(){
						Expect(err).NotTo(HaveOccurred())
						Expect(greet).To(Equal(greet))
						})
					})
				})
		}
	})

})
