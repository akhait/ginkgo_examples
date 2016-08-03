package validation

import (
	"testing"
)

type test_data struct {
	username string
	greeting string
}

//Preparing test data
var tests = []test_data{
	{"John Doe", "Hello, John Doe"},
	{"John", "Your username should be longer than 8 characters"},
	{"Katy", "Your username should be longer than 8 characters"},
	{"", "Your username should be longer than 8 characters"},
	{"John#Doe", "Do not use special characters in your username"},
}

func TestRegisterNewUsername(t *testing.T) {
//Running tests
	for i, test_case := range tests {
		expected := RegisterNewUsername(test_case.username)
		if expected != test_case.greeting {
			t.Error(
				i+1, ")",
				"For", test_case.username,
				"expected", test_case.greeting,
				"got", expected,
			)
		}
	}
}