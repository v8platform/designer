package tests

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
	"io/ioutil"
	"os"
	"path"
)

type TestSuite struct {
	suite.Suite
	TempIB string
	Runner *runner.Runner
	Pwd    string
}

func NewFileIB(path string) TempInfobase {

	ib := TempInfobase{
		File: path,
	}

	return ib
}

type TempInfobase struct {
	File string
}

func (ib TempInfobase) Path() string {
	return ib.File
}

func (ib TempInfobase) ConnectionString() string {
	return "/F" + ib.File
}

func (ib TempInfobase) Values() []string {

	return []string{"file=" + ib.File}
}

type TempCreateInfobase struct {
}

type TestCommon struct {
	DisableStartupDialogs  bool `v8:"/DisableStartupDialogs" json:"disable_startup_dialogs"`
	DisableStartupMessages bool `v8:"/DisableStartupDialogs" json:"disable_startup_messages"`
	Visible                bool `v8:"/Visible" json:"visible"`
	ClearCache             bool `v8:"/ClearCache" json:"clear_cache"`
}

func (cv TestCommon) Values() []string {

	var v []string

	if cv.Visible {
		v = append(v, "/Visible")
	}
	if cv.DisableStartupDialogs {
		v = append(v, "/DisableStartupDialogs")
	}
	if cv.DisableStartupMessages {
		v = append(v, "/DisableStartupMessages")
	}
	if cv.ClearCache {
		v = append(v, "/ClearCache")
	}

	return v
}

func (ib TempCreateInfobase) Command() string {
	return "CREATEINFOBASE"
}

func (ib TempCreateInfobase) Check() error {
	return nil
}
func (ib TempCreateInfobase) Values() []string {
	var v []string
	return v
}

func (s *TestSuite) Run(where runner.Infobase, command runner.Command, opts ...interface{}) error {
	return runner.Run(where, command, opts...)
}

func (s *TestSuite) R() *require.Assertions {
	return s.Require()
}

func (t *TestSuite) SetupSuite() {

	//common := TestCommon{
	//	DisableStartupDialogs:  true,
	//	DisableStartupMessages: true,
	//	Visible:                false,
	//	//ClearCache:             false,
	//}

	//t.Runner = runner.NewRunner(runner.WithCommonValues(common))
	ibPath, _ := ioutil.TempDir("", "1c_DB_")
	t.TempIB = ibPath
	pwd, _ := os.Getwd()

	t.Pwd = path.Join(pwd)

}

func (t *TestSuite) AfterTest(suite, testName string) {
	t.ClearTempInfoBase()
}

func (t *TestSuite) BeforeTest(suite, testName string) {
	t.CreateTempInfoBase()
}

func (t *TestSuite) CreateTempInfoBase() {

	ib := TempInfobase{File: t.TempIB}

	err := t.Run(ib, TempCreateInfobase{},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *TestSuite) ClearTempInfoBase() {

	err := os.RemoveAll(t.TempIB)
	t.R().NoError(err, errors.GetErrorContext(err))
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
