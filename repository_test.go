package designer

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

type TempInfobase struct {
	File string
}

func (ib TempInfobase) ConnectionString() string {
	return "/IBConnectionString File=" + ib.File
}

type RepositoryTestSuite struct {
	suite.Suite
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) r() *require.Assertions {
	return s.Require()
}

func (t *RepositoryTestSuite) TestCreateRepository() {

	if testing.Short() {
		t.T().Skip()
	}

	tempDir, _ := ioutil.TempDir("", "v8_temp_ib")
	where := TempInfobase{tempDir}
	confFile := filepath.Join("tests", "fixtures", "0.9", "1Cv8.cf")

	err := runner.Run(nil, CreateFileInfoBaseOptions{File: tempDir})

	t.r().NoError(err, "err create infobase: %s", err)

	err = runner.Run(where, LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.r().NoError(err, "err load cf: %s", err)

	repPath, _ := ioutil.TempDir("", "1c_rep_")
	repo := Repository{
		Path: repPath,
	}
	createOptions := repo.Create(true, REPOSITORY_SUPPORT_NOT_SUPPORTED, REPOSITORY_SUPPORT_NOT_SUPPORTED)

	err = runner.Run(where, createOptions,
		runner.WithTimeout(30))

	t.r().NoError(err, "err create repository: %s", errors.GetErrorContext(err))

	err = os.RemoveAll(repPath)
	t.r().NoError(err)
	err = os.RemoveAll(tempDir)
	t.r().NoError(err)
}

func TestRepositoryCreateOptions_Values(t *testing.T) {
	type fields struct {
		Designer                  Designer
		Repository                Repository
		command                   struct{}
		AllowConfigurationChanges bool
		ChangesAllowedRule        RepositorySupportEditObjectsType
		ChangesNotRecommendedRule RepositorySupportEditObjectsType
		NoBind                    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"simple",
			fields{
				Repository: Repository{
					"./repo",
					"admin",
					"pws",
					"",
				},
			},
			[]string{
				"/ConfigurationRepositoryF ./repo",
				"/ConfigurationRepositoryN admin",
				"/ConfigurationRepositoryP pws",
				"/ConfigurationRepositoryCreate",
			},
		},
		{
			"no bind",
			fields{
				Repository: Repository{
					"./repo",
					"admin",
					"pws",
					"",
				},
				NoBind: true,
			},
			[]string{
				"/ConfigurationRepositoryF ./repo",
				"/ConfigurationRepositoryN admin",
				"/ConfigurationRepositoryP pws",
				"/ConfigurationRepositoryCreate",
				"-NoBind",
			},
		},
		{
			"all",
			fields{
				Repository: Repository{
					"./repo",
					"admin",
					"pws",
					"temp_ext",
				},
				AllowConfigurationChanges: true,
				ChangesAllowedRule:        REPOSITORY_SUPPORT_NOT_SUPPORTED,
				ChangesNotRecommendedRule: REPOSITORY_SUPPORT_NOT_SUPPORTED,
				NoBind:                    true,
			},
			[]string{
				"/ConfigurationRepositoryF ./repo",
				"/ConfigurationRepositoryN admin",
				"/ConfigurationRepositoryP pws",
				"-Extension temp_ext",
				"/ConfigurationRepositoryCreate",
				"-AllowConfigurationChanges",
				"-ChangesAllowedRule ObjectNotSupported",
				"-ChangesNotRecommendedRule ObjectNotSupported",
				"-NoBind",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ib := RepositoryCreateOptions{
				Designer:                  tt.fields.Designer,
				Repository:                tt.fields.Repository,
				command:                   tt.fields.command,
				AllowConfigurationChanges: tt.fields.AllowConfigurationChanges,
				ChangesAllowedRule:        tt.fields.ChangesAllowedRule,
				ChangesNotRecommendedRule: tt.fields.ChangesNotRecommendedRule,
				NoBind:                    tt.fields.NoBind,
			}
			if got := ib.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryCreateOptions_WithRepository(t *testing.T) {

	type args struct {
		repository Repository
	}
	tests := []struct {
		name string
		args args
		want RepositoryCreateOptions
	}{
		{
			"simple",
			args{
				repository: Repository{
					Path:     "./repo",
					User:     "admin",
					Password: "pwd",
				},
			},
			RepositoryCreateOptions{
				Repository: Repository{
					Path:     "./repo",
					User:     "admin",
					Password: "pwd",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := RepositoryCreateOptions{}
			if got := o.WithRepository(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositoryRightType_MarshalV8(t *testing.T) {
	tests := []struct {
		name    string
		t       RepositoryRightType
		want    string
		wantErr bool
	}{
		{
			"ReadOnly",
			REPOSITORY_RIGHT_READ,
			"ReadOnly",
			false,
		},
		{
			"Administration",
			REPOSITORY_RIGHT_ADMIN,
			"Administration",
			false,
		},
		{
			"LockObjects",
			REPOSITORY_RIGHT_LOCK,
			"LockObjects",
			false,
		},
		{
			"ManageConfigurationVersions",
			REPOSITORY_RIGHT_MANAGE_VERSIONS,
			"ManageConfigurationVersions",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.MarshalV8()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalV8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MarshalV8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepositorySupportEditObjectsType_MarshalV8(t *testing.T) {
	tests := []struct {
		name    string
		t       RepositorySupportEditObjectsType
		want    string
		wantErr bool
	}{
		{
			"ObjectNotSupported",
			REPOSITORY_SUPPORT_NOT_SUPPORTED,
			"ObjectNotSupported",
			false,
		},
		{
			"ObjectIsEditableSupportEnabled",
			REPOSITORY_SUPPORT_IS_EDITABLE,
			"ObjectIsEditableSupportEnabled",
			false,
		},
		{
			"ObjectNotEditable",
			REPOSITORY_SUPPORT_NOT_EDITABLE,
			"ObjectNotEditable",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.MarshalV8()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalV8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MarshalV8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Create(t *testing.T) {
	type fields struct {
		Path      string
		User      string
		Password  string
		Extension string
	}
	type args struct {
		noBind                        bool
		allowedAndNotRecommendedRules []RepositorySupportEditObjectsType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   RepositoryCreateOptions
	}{
		{
			"simple",
			fields{
				Path:      "./repo",
				User:      "admin",
				Password:  "pwd",
				Extension: "",
			},
			args{
				noBind: false,
			},
			RepositoryCreateOptions{
				Designer: NewDesigner(),
				Repository: Repository{
					Path:      "./repo",
					User:      "admin",
					Password:  "pwd",
					Extension: "",
				},
				NoBind: false,
			},
		},
		{
			"ext",
			fields{
				Path:      "./repo",
				User:      "admin",
				Password:  "pwd",
				Extension: "test",
			},
			args{
				noBind: false,
			},
			RepositoryCreateOptions{
				Designer: NewDesigner(),
				Repository: Repository{
					Path:      "./repo",
					User:      "admin",
					Password:  "pwd",
					Extension: "test",
				},
				NoBind: false,
			},
		},
		{
			"all",
			fields{
				Path:      "./repo",
				User:      "admin",
				Password:  "pwd",
				Extension: "test",
			},
			args{
				noBind: true,
				allowedAndNotRecommendedRules: []RepositorySupportEditObjectsType{
					REPOSITORY_SUPPORT_IS_EDITABLE,
					REPOSITORY_SUPPORT_NOT_EDITABLE,
				},
			},
			RepositoryCreateOptions{
				Designer: NewDesigner(),
				Repository: Repository{
					Path:      "./repo",
					User:      "admin",
					Password:  "pwd",
					Extension: "test",
				},
				AllowConfigurationChanges: true,
				ChangesAllowedRule:        REPOSITORY_SUPPORT_IS_EDITABLE,
				ChangesNotRecommendedRule: REPOSITORY_SUPPORT_NOT_EDITABLE,
				NoBind:                    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				Path:      tt.fields.Path,
				User:      tt.fields.User,
				Password:  tt.fields.Password,
				Extension: tt.fields.Extension,
			}
			if got := r.Create(tt.args.noBind, tt.args.allowedAndNotRecommendedRules...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Values(t *testing.T) {
	type fields struct {
		Path      string
		User      string
		Password  string
		Extension string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"simple",
			fields{
				Path:      "./repo",
				User:      "admin",
				Password:  "pwd",
				Extension: "",
			},
			[]string{
				"/ConfigurationRepositoryF ./repo",
				"/ConfigurationRepositoryN admin",
				"/ConfigurationRepositoryP pwd",
			},
		},
		{
			"ext",
			fields{
				Path:      "./repo",
				User:      "admin",
				Password:  "pwd",
				Extension: "test",
			},
			[]string{
				"/ConfigurationRepositoryF ./repo",
				"/ConfigurationRepositoryN admin",
				"/ConfigurationRepositoryP pwd",
				"-Extension test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				Path:      tt.fields.Path,
				User:      tt.fields.User,
				Password:  tt.fields.Password,
				Extension: tt.fields.Extension,
			}
			if got := r.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
