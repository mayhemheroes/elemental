package e2e_test

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/rancher-sandbox/ele-testhelpers/vm"
)

var _ = Describe("E2E - Getting logs node", Label("logs"), func() {
	var (
		sut *vm.SUT
	)
	BeforeEach(func() {
		sut = &vm.SUT{
			Host:     "192.168.122.2:22",
			Username: "root",
			Password: "root",
		}
	})
	It("gets the downstream logs", func() {
		sut.Command("curl -L https://github.com/Itxaka/elemental-operator/releases/download/v100.0.0/elemental-support_100.0.0_linux_amd64 -o /tmp/elemental-support")
		sut.Command("/tmp/elemental-support")
		out, _ := sut.Command("find /root -name `hostname`*.tar.gz -print")
		sut.GatherLog(out)
	})
})
