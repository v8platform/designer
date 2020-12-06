package designer

import (
	"github.com/v8platform/marshaler"
)

///ConfigurationRepositoryBindCfg [-Extension <имя расширения>] [-forceBindAlreadyBindedUser][-forceReplaceCfg]
//— подключение неподключенной конфигурации к хранилищу конфигурации. Доступны параметры:
//
//-Extension <имя расширения> — Имя расширения. Если параметр не указан,
// выполняется попытка соединения с хранилищем основной конфигурации, и команда выполняется для основной конфигурации.
// Если параметр указан, выполняется попытка соединения с хранилищем указанного расширения, и команда выполняется для этого хранилища.
//
//-forceBindAlreadyBindedUser — Подключение будет выполнено даже в случае,
// если для данного пользователя уже есть конфигурация, связанная с данным хранилищем;
//
//-forceReplaceCfg — Если конфигурация не пустая, текущая конфигурация будет заменена конфигурацией из хранилища.
type RepositoryBindCfgOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryBindCfg" json:"-"`

	//-Extension <имя расширения> — Имя расширения. Если параметр не указан,
	// выполняется попытка соединения с хранилищем основной конфигурации, и команда выполняется для основной конфигурации.
	// Если параметр указан, выполняется попытка соединения с хранилищем указанного расширения, и команда выполняется для этого хранилища.
	ForceBindAlreadyBindedUser bool `v8:"-forceBindAlreadyBindedUser, optional" json:"force_bind"`

	//-forceReplaceCfg — Если конфигурация не пустая, текущая конфигурация будет заменена конфигурацией из хранилища.
	ForceReplaceCfg bool `v8:"-ForceReplaceCfg, optional" json:"force_replace"`
}

func (ib RepositoryBindCfgOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (r Repository) Bind(force ...bool) RepositoryBindCfgOptions {

	command := RepositoryBindCfgOptions{
		Designer:   NewDesigner(),
		Repository: r,
	}

	if len(force) > 0 {
		command.ForceReplaceCfg = force[0]
		command.ForceBindAlreadyBindedUser = force[0]
		if len(force) > 2 {
			command.ForceBindAlreadyBindedUser = force[1]
		}
	}

	return command

}

///ConfigurationRepositoryUnbindCfg [-Extension <имя расширения>] [-force]
//— отключение конфигурации от хранилища конфигурации (у пользователя должны быть административные права в данной информационной базе).
//Если пользователь аутентифицируется в хранилище (интерактивно или через параметры командной строки),
//то отключение конфигурации от хранилища также отражается в самом хранилище конфигурации (информация о подключении удаляется),
//если же пользователь не аутентифицировался в хранилище, то производится только локальное отключение конфигурации от хранилища.
//
//В случае, если в конфигурации имеются захваченные объекты, которые были изменены относительно хранилища,
//то будет выдано соответствующее сообщение и отключения не выполнится.
//
//-Extension <имя расширения> — имя расширения. Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
//и команда выполняется для основной конфигурации. Если параметр указан, выполняется попытка соединения с хранилищем указанного расширения, и команда выполняется для этого хранилища.
//
//-force — параметр для форсирования отключения от хранилища
//(пропуск диалога аутентификации, если не указаны параметры пользователя хранилища, игнорирование наличия захваченных и измененных объектов).
//
type RepositoryUnbindCfgOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryUnbindCfg" json:"-"`

	//-force — параметр для форсирования отключения от хранилища
	//(пропуск диалога аутентификации, если не указаны параметры пользователя хранилища, игнорирование наличия захваченных и измененных объектов).
	Force bool `v8:"-force, optional" json:"force"`
}

func (ib RepositoryUnbindCfgOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (r Repository) Unbind(force ...bool) RepositoryUnbindCfgOptions {

	command := RepositoryUnbindCfgOptions{
		Designer:   NewDesigner(),
		Repository: r,
	}

	if len(force) > 0 {
		command.Force = force[0]
	}

	return command

}

///ConfigurationRepositoryDumpCfg [-Extension <имя расширения>] <имя cf файла> [-v <номер версии хранилища>]
//— сохранить конфигурацию из хранилища в файл (пакетный режим запуска). Доступны параметры:
//
//-Extension <имя расширения> — Имя расширения. Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
// и команда выполняется для основной конфигурации.
// Если параметр указан, выполняется попытка соединения с хранилищем указанного расширения, и команда выполняется для этого хранилища.
//
//-v <номер версии хранилища> v — Номер версии, если номер версии не указан,
// или равен -1, будет сохранена последняя версия.
type RepositoryDumpCfgOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	File string `v8:"/ConfigurationRepositoryDumpCfg" json:"file"`

	//-v <номер версии хранилища> v — Номер версии, если номер версии не указан,
	// или равен -1, будет сохранена последняя версия.
	Version int64 `v8:"-v, optional" json:"version"`
}

func (ib RepositoryDumpCfgOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (r Repository) DumpCfg(file string, version ...int64) RepositoryDumpCfgOptions {

	command := RepositoryDumpCfgOptions{
		Designer:   NewDesigner(),
		Repository: r,
		File:       file,
	}

	if len(version) > 0 {
		command.Version = version[0]
	}

	return command

}

///ConfigurationRepositoryUpdateCfg [-Extension <имя расширения>]
//[-v <номер версии хранилища>] [-revised] [-force]
//[-objects <имя файла со списком объектов>]
//— обновить конфигурацию хранилища из хранилища (пакетный режим запуска).
//
//-Extension <имя расширения> — имя расширения.
//Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
//и команда выполняется для основной конфигурации.
//Если параметр указан, выполняется попытка соединения с хранилищем указанного расширения,
//и команда выполняется для этого хранилища.
//
//-v <номер версии хранилища> — номер версии в хранилище конфигурации.
//Если конфигурация подключена к хранилищу, то номер версии (если он указан) игнорируется
//и будет получена актуальная версия конфигурации хранилища.
//Если конфигурация не подключена к хранилищу, то выполняется получение указанной версии,
//а если версия не указана (или значение равно -1) – будет получена актуальная версия конфигурации;
//
//-revised — получать захваченные объекты, если потребуется.
//Если конфигурация не подключена к хранилищу, то параметр игнорируется;
//
//-force — если при пакетном обновлении конфигурации из хранилища должны быть получены новые объекты конфигурации или удалиться существующие,
//указание этого параметра свидетельствует о подтверждении пользователем описанных выше операций. Если параметр не указан — действия выполнены не будут.
//
//-objects <имя файла со списком объектов> — путь к файлу формата XML со списком объектов.
//Если параметр используется, будет выполнена попытка обновления только объектов, указанных в файле.
//Если параметр не используется, обновляется вся конфигурация целиком.
type RepositoryUpdateCfgOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryUpdateCfg" json:"-"`

	//-v <номер версии хранилища> — номер версии в хранилище конфигурации.
	//Если конфигурация подключена к хранилищу, то номер версии (если он указан) игнорируется
	//и будет получена актуальная версия конфигурации хранилища.
	//Если конфигурация не подключена к хранилищу, то выполняется получение указанной версии,
	//а если версия не указана (или значение равно -1) – будет получена актуальная версия конфигурации;
	Version int64 `v8:"-v, optional" json:"version"`

	//-revised — получать захваченные объекты, если потребуется.
	//Если конфигурация не подключена к хранилищу, то параметр игнорируется;
	//
	Revised bool `v8:"-revised, optional" json:"revised"`

	//-force — если при пакетном обновлении конфигурации из хранилища должны быть получены новые объекты конфигурации или удалиться существующие,
	//указание этого параметра свидетельствует о подтверждении пользователем описанных выше операций. Если параметр не указан — действия выполнены не будут.
	//
	Force bool `v8:"-force, optional" json:"force"`

	//-objects <имя файла со списком объектов> — путь к файлу формата XML со списком объектов.
	//Если параметр используется, будет выполнена попытка обновления только объектов, указанных в файле.
	//Если параметр не используется, обновляется вся конфигурация целиком.
	Objects string `v8:"-objects, optional" json:"objects"`
}

func (ib RepositoryUpdateCfgOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (o RepositoryUpdateCfgOptions) WithObjects(objectsFile string) RepositoryUpdateCfgOptions {

	newO := o
	newO.Objects = objectsFile
	return newO

}

func (o RepositoryUpdateCfgOptions) WithRepository(repository Repository) RepositoryUpdateCfgOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

func (r Repository) UpdateCfg(version int64, force ...bool) RepositoryUpdateCfgOptions {

	command := RepositoryUpdateCfgOptions{
		Designer:   NewDesigner(),
		Repository: r,
		Version:    version,
	}

	if len(force) > 0 {
		command.Force = force[0]
	}

	return command

}

func (o RepositoryDumpCfgOptions) WithRepository(repository Repository) RepositoryDumpCfgOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

func (o RepositoryBindCfgOptions) WithRepository(repository Repository) RepositoryBindCfgOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

func (o RepositoryUnbindCfgOptions) WithRepository(repository Repository) RepositoryUnbindCfgOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}
