package datetime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDatetime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Datetime Suite")
}
