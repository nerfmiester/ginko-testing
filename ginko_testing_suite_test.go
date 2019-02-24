package ginko_testing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGinkoTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkoTesting Suite")
}
