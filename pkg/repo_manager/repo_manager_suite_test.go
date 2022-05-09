package repo_manager

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepoManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RepoManager Suite")
}
