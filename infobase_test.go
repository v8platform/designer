package designer

import (
	"github.com/v8platform/designer/tests"
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
	"io/ioutil"
	"path"
)

func (t *designerTestSuite) TestDumpIB() {
	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")
	ib := tests.NewFileIB(t.TempIB)

	err := runner.Run(ib, LoadCfgOptions{
		File: confFile},
		runner.WithTimeout(30))

	t.R().NoError(err, "error load cf: %s", errors.GetErrorContext(err)["message"])

	dtFile, _ := ioutil.TempFile("", "temp_dt")
	dtFile.Close()

	err = runner.Run(ib, DumpIBOptions{
		File: dtFile.Name()},
		runner.WithTimeout(30))

	t.R().NoError(err, "error dump ib %s", errors.GetErrorContext(err)["message"])

	fileCreated, err2 := tests.Exists(dtFile.Name())
	t.R().NoError(err2)
	t.R().True(fileCreated, "Файл должен быть создан")

}

func (t *designerTestSuite) TestRestoreIB() {
	dtFile, _ := ioutil.TempFile("", "temp_dt")
	dtFile.Close()
	ib := tests.NewFileIB(t.TempIB)

	err := runner.Run(ib, DumpIBOptions{
		File: dtFile.Name()},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	newIB := tests.NewFileIB(t.TempIB)

	err = runner.Run(newIB, RestoreIBOptions{
		File: dtFile.Name()},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	fileCreated, err2 := tests.Exists(dtFile.Name())
	t.R().NoError(err2)
	t.R().True(fileCreated, "Файл должен быть создан")

}
