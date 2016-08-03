package validation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestValidationGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ValidationGinkgo Suite")
}
