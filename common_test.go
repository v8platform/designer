package designer

import (
	"github.com/khorevaa/go-v8platform/errors"
	"github.com/stretchr/testify/suite"
	"github.com/v8platform/designer/tests"
	"github.com/v8platform/runner"
	"path"
	"testing"
)

type designerTestSuite struct {
	tests.TestSuite
}

func TestDesigner(t *testing.T) {
	suite.Run(t, new(designerTestSuite))
}

func (t *designerTestSuite) TestLoadCfg() {
	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")

	err := runner.Run(tests.NewFileIB(t.TempIB), LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *designerTestSuite) TestLoadCfgWithUpdateCfgDB() {

	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")
	loadCfg := LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile,
		UpdateDBCfg: &UpdateDBCfgOptions{
			Dynamic: false,
		},
	}

	err := runner.Run(tests.NewFileIB(t.TempIB), loadCfg,
		runner.WithTimeout(30),
	)

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *designerTestSuite) TestUpdateCfg() {

	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")
	loadCfg := LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile,
	}.WithUpdateDBCfg(UpdateDBCfgOptions{})

	err := runner.Run(tests.NewFileIB(t.TempIB), loadCfg,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	confFile2 := path.Join(t.Pwd, "..", "tests", "fixtures", "1.0", "1Cv8.cf")
	task := UpdateCfgOptions{
		Designer: NewDesigner(),
		File:     confFile2,
	}

	err = runner.Run(tests.NewFileIB(t.TempIB), task,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	//t.R().Equal(len(codes), 1, "Промокод должен быть START")
	//t.R().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
