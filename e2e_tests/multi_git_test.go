package e2e_tests_test

import (
	. "github.com/lekan-pvp/multi-git/pkg/helpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

const baseDir = "/tmp/test-multi-git"

var repoList string

var _ = Describe("multi-git e2e tests", func() {
	var err error

	removeAll := func() {
		err = os.RemoveAll(baseDir)
		Ω(err).Should(BeNil())
	}

	BeforeEach(func() {
		removeAll()
		err = CreateDir(baseDir, "", false)
		Ω(err).Should(BeNil())
	})

	AfterSuite(removeAll)
})
