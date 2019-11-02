package mic_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mic Suite")
}
