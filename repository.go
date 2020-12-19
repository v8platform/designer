package designer

import (
	"github.com/hashicorp/go-multierror"
	"github.com/v8platform/errors"
	"github.com/v8platform/marshaler"
	"strings"
)

type RepositoryRightType string
type RepositorySupportEditObjectsType string

const (
	REPOSITORY_RIGHT_READ            RepositoryRightType = "ReadOnly"
	REPOSITORY_RIGHT_LOCK                                = "LockObjects"
	REPOSITORY_RIGHT_MANAGE_VERSIONS                     = "ManageConfigurationVersions"
	REPOSITORY_RIGHT_ADMIN                               = "Administration"
)

const (
	REPOSITORY_SUPPORT_NOT_EDITABLE  RepositorySupportEditObjectsType = "ObjectNotEditable"
	REPOSITORY_SUPPORT_IS_EDITABLE                                    = "ObjectIsEditableSupportEnabled"
	REPOSITORY_SUPPORT_NOT_SUPPORTED                                  = "ObjectNotSupported"
)

func (t RepositorySupportEditObjectsType) MarshalV8() (string, error) {
	return string(t), nil
}

func (t RepositoryRightType) MarshalV8() (string, error) {
	return string(t), nil
}

type Repository struct {
	///ConfigurationRepositoryF <каталог хранилища>
	//— указание имени каталога хранилища.
	Path string `v8:"/ConfigurationRepositoryF" json:"path"`

	///ConfigurationRepositoryN <имя>
	//— указание имени пользователя хранилища.
	User string `v8:"/ConfigurationRepositoryN, default=Администратор, optional" json:"user"`

	///ConfigurationRepositoryP <пароль>
	//— указание пароля пользователя хранилища.
	Password string `v8:"/ConfigurationRepositoryP, optional" json:"password"`

	//-Extension <имя расширения> — Имя расширения.
	// Если параметр указан, выполняется попытка соединения с
	// хранилищем указанного расширения, и команда выполняется для этого хранилища.
	Extension string `v8:"-Extension, optional" json:"extension"`
}

func (r Repository) Values() []string {

	v, _ := marshaler.Marshal(r)
	return v

}

//ConfigurationRepositoryCreate
///ConfigurationRepositoryCreate [-Extension <имя расширения>] [-AllowConfigurationChanges
//-ChangesAllowedRule <Правило поддержки> -ChangesNotRecommendedRule <Правило поддержки>] [-NoBind]
//— предназначен для создания хранилища конфигурации. Доступны следующие параметры:
//Пример:
//
//DESIGNER /F "D:\V8\Cfgs83\ИБ83" /ConfigurationRepositoryF "D:\V8\Cfgs83" /ConfigurationRepositoryN "Администратор"
// /ConfigurationRepositoryP "123456" /ConfigurationRepositoryCreate - AllowConfigurationChanges
// -ChangesAllowedRule ObjectNotEditable -ChangesNotRecommendedRule ObjectNotEditable
type RepositoryCreateOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryCreate" json:"-"`

	//-AllowConfigurationChanges — если конфигурация находится на поддержке без возможности изменения, будет включена возможность изменения.
	AllowConfigurationChanges bool `v8:"-AllowConfigurationChanges, optional" json:"allow_configuration_changes"`

	//-ChangesAllowedRule <Правило поддержки> — устанавливает правило поддержки для объектов,
	// для которых изменения разрешены поставщиком. Может быть установлено одно из следующих правил:
	//	ObjectNotEditable - Объект поставщика не редактируется,
	//	ObjectIsEditableSupportEnabled - Объект поставщика редактируется с сохранением поддержки,
	//	ObjectNotSupported - Объект поставщика снят с поддержки.
	ChangesAllowedRule RepositorySupportEditObjectsType `v8:"-ChangesAllowedRule, optional" json:"changes_allowed_rule"`

	//-ChangesNotRecommendedRule — устанавливает правило поддержки для объектов,
	// для которых изменения не рекомендуются поставщиком. Может быть установлено одно из следующих правил:
	//	ObjectNotEditable - Объект поставщика не редактируется,
	//	ObjectIsEditableSupportEnabled - Объект поставщика редактируется с сохранением поддержки,
	//	ObjectNotSupported - Объект поставщика снят с поддержки.
	ChangesNotRecommendedRule RepositorySupportEditObjectsType `v8:"-ChangesNotRecommendedRule, optional" json:"changes_not_recommended_rule"`

	//-NoBind — К созданному хранилищу подключение выполнено не будет.
	NoBind bool `v8:"-NoBind, optional" json:"no_bind"`
}

func (r Repository) Create(noBind bool, allowedAndNotRecommendedRules ...RepositorySupportEditObjectsType) RepositoryCreateOptions {

	command := RepositoryCreateOptions{
		Designer:   NewDesigner(),
		Repository: r,
		NoBind:     noBind,
	}

	if len(allowedAndNotRecommendedRules) > 0 {
		command.AllowConfigurationChanges = true
		command.ChangesAllowedRule = allowedAndNotRecommendedRules[0]
		if len(allowedAndNotRecommendedRules) == 2 {
			command.ChangesNotRecommendedRule = allowedAndNotRecommendedRules[1]
		}
	}

	return command

}

func (o RepositoryCreateOptions) WithRepository(repository Repository) RepositoryCreateOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

func (ib RepositoryCreateOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	fixExtensionIndex(&v)
	return v

}

func (ib RepositoryCreateOptions) Check() error {

	var err multierror.Error

	if ib.AllowConfigurationChanges && (len(ib.ChangesNotRecommendedRule) == 0 || len(ib.ChangesAllowedRule) == 0) {

		multierror.Append(&err, errors.Check.New("configuration changes must be set").
			WithContext("msg", "field ChangesNotRecommendedRule or ChangesAllowedRule not set"))

	}

	return err.ErrorOrNil()

}

func fixExtensionIndex(values *[]string) {

	val := *values
	extIdx := -1
	for i := 0; i < len(val); i++ {
		if strings.Contains(val[i], "-Extension") {
			extIdx = i
			break
		}
	}

	if extIdx != -1 {
		ext := val[extIdx]

		val = append(val[:extIdx], val[extIdx+1:]...)
		val = append(val, ext)

		values = &val
	}

}
