package sub_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSub(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sub Suite")
}