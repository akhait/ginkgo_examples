package benchmarking_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBenchmarking(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Benchmarking Suite")
}
