package ginko_testing_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkoTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkoTesting Suite")
}

// test me
 