package mic_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

func TestMic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mic Suite")
}

var ctl *gomock.Controller
var cleaner func()
var ctx context.Context
var _ = BeforeEach(func() {
	ctl = gomock.NewController(GinkgoT())
	cleaner = ctl.Finish
	ctx = context.Background()
})
var _ = AfterEach(func() { cleaner() })
