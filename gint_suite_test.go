package gint_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGint(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gint Suite")
}
