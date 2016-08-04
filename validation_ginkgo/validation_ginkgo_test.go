package validation_ginkgo_test

import (
	"ginkgo_examples/validation_ginkgo"
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type test_data struct {
	username string
	greeting string
}

var postests = []test_data{
			{"12345678", "Hello, 12345678"},
			{"SuperUserName","Hello, SuperUserName"},
			{"Test This One", "Hello, Test This One"},
		}

//Top-level container
var _ = Describe("ValidationGinkgo", func() {
	//Positive test cases
	Describe("Positive tests - no errors should occure", func(){
		for _, test_case := range postests {
			greeting := test_case.greeting
			greet, err := validation_ginkgo.RegisterNewUsername(test_case.username)
			Describe("Username", func(){
				Context("longer than 8 chars and no restricted chars", func(){
					It("should be legal", func(){
						Expect(err).NotTo(HaveOccurred())
						Expect(greet).To(Equal(greeting))
						})
					})
				})
		}
	})
	//Negative test cases
	Describe("Negative tests - should return empty string and an error", func(){
		//Empty username
		Describe("Empty username", func(){
			Context("is shorter than 8 characters", func(){
				It("should return error", func(){
					greet, err := validation_ginkgo.RegisterNewUsername("")
					Expect(err).To(Equal(errors.New("Your username should be longer than 8 characters")))
					Expect(greet).To(BeEmpty())
				})
			})
		})
		//Short username
		Describe("Short username", func(){
			Context("is shorter than 8 characters", func(){
				It("should return corresponding error", func(){
					greet, err := validation_ginkgo.RegisterNewUsername("Short")
					Expect(err).To(Equal(errors.New("Your username should be longer than 8 characters")))
					Expect(greet).To(BeEmpty())
				})
			})
		})
		//Username with special characters
		Describe("Username with ^, @, #, $, &", func(){
			Context("contains special symbols", func(){
				It("should return corresponding error", func(){
					greet, err := validation_ginkgo.RegisterNewUsername("John#Doe")
					Expect(err).To(Equal(errors.New("Do not use special characters in your username")))
					Expect(greet).To(BeEmpty())
				})
			})
		})
	})

})
