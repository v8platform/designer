package repository

import (
	"github.com/stretchr/testify/suite"
	"github.com/v8platform/designer"
	"github.com/v8platform/designer/tests"
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
	"io/ioutil"
	"path"
	"testing"
)

type RepositoryTestSuite struct {
	tests.TestSuite
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (t *RepositoryTestSuite) TestCreateRepository() {

	confFile := path.Join(t.Pwd, "..", "..", "tests", "fixtures", "0.9", "1Cv8.cf")

	err := runner.Run(tests.NewFileIB(t.TempIB), designer.LoadCfgOptions{
		File: confFile},
		runner.WithTimeout(30))

	t.R().NoError(err, "err load cf: %s", errors.GetErrorContext(err))

	repPath, _ := ioutil.TempDir("", "1c_rep_")

	createOptions := RepositoryCreateOptions{
		NoBind:                    true,
		AllowConfigurationChanges: true,
		ChangesAllowedRule:        REPOSITORY_SUPPORT_NOT_SUPPORTED,
		ChangesNotRecommendedRule: REPOSITORY_SUPPORT_NOT_SUPPORTED,
	}.WithPath(repPath)

	err = runner.Run(tests.NewFileIB(t.TempIB), createOptions,
		runner.WithTimeout(30))

	t.R().NoError(err, "err create repository: %s", errors.GetErrorContext(err))

}
