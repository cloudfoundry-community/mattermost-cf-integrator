package mci_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mci Suite")
}
