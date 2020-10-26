package repository

import (
	"github.com/stretchr/testify/suite"
	"github.com/v8platform/designer"
	"github.com/v8platform/designer/tests"
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

type RepositoryCfgTestSuite struct {
	tests.TestSuite
	Repository Repository
}

func TestRepositoryCfg(t *testing.T) {
	suite.Run(t, new(RepositoryCfgTestSuite))
}

func (t *RepositoryCfgTestSuite) AfterTest(suite, testName string) {
	t.ClearTempInfoBase()
}

func (t *RepositoryCfgTestSuite) BeforeTest(suite, testName string) {
	t.CreateTempInfoBase()
	t.createTestRepository()

}

func (t *RepositoryCfgTestSuite) createTestRepository() {
	confFile := path.Join(t.Pwd, "..", "..", "tests", "fixtures", "0.9", "1Cv8.cf")

	err := runner.Run(tests.NewFileIB(t.TempIB), designer.LoadCfgOptions{
		Designer: designer.NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	repPath, _ := ioutil.TempDir("", "1c_rep_")

	t.Repository = Repository{
		Path: repPath,
		User: "admin",
	}

	createOptions := RepositoryCreateOptions{
		NoBind:                    true,
		AllowConfigurationChanges: true,
		ChangesAllowedRule:        REPOSITORY_SUPPORT_NOT_SUPPORTED,
		ChangesNotRecommendedRule: REPOSITORY_SUPPORT_NOT_SUPPORTED,
	}.WithRepository(t.Repository)

	err = runner.Run(tests.NewFileIB(t.TempIB), createOptions,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))
}

func (t *RepositoryCfgTestSuite) TestRepositoryBindCfg() {

	command := RepositoryBindCfgOptions{
		ForceBindAlreadyBindedUser: true,
		ForceReplaceCfg:            true,
	}.WithRepository(t.Repository)

	err := runner.Run(tests.NewFileIB(t.TempIB), command,
		runner.WithTimeout(30),
		runner.WithUC("code"))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *RepositoryCfgTestSuite) TestRepositoryUnbindCfg() {

	command := RepositoryBindCfgOptions{
		ForceBindAlreadyBindedUser: true,
		ForceReplaceCfg:            true,
	}.WithRepository(t.Repository)

	err := runner.Run(tests.NewFileIB(t.TempIB), command,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	command2 := RepositoryUnbindCfgOptions{
		Force: true,
	}.WithRepository(t.Repository)

	err = runner.Run(tests.NewFileIB(t.TempIB), command2,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *RepositoryCfgTestSuite) TestRepositoryDumpCfg() {

	cfFile, _ := ioutil.TempFile("", "v8_DumpResult_*.cf")

	command := RepositoryDumpCfgOptions{
		File: cfFile.Name(),
	}.WithRepository(t.Repository)

	cfFile.Close()

	err := runner.Run(tests.NewFileIB(t.TempIB), command,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	fileCreated, err2 := Exists(command.File)
	t.R().NoError(err2)
	t.R().True(fileCreated, "Файл базы должен быть создан")

}

func (t *RepositoryCfgTestSuite) TestRepositoryUpdateCfg() {

	command := RepositoryUpdateCfgOptions{
		Force:   true,
		Version: -1,
		Revised: true,
	}.WithRepository(t.Repository)

	err := runner.Run(tests.NewFileIB(t.TempIB), command,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
