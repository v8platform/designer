package designer

import (
	"github.com/v8platform/marshaler"
)

///ConfigurationRepositoryAddUser [-Extension <имя расширения>] -User <Имя> -Pwd <Пароль> -Rights <Права> [-RestoreDeletedUser]
//— создание пользователя хранилища конфигурации.
// Пользователь, от имени которого выполняется подключение к хранилищу,
// должен обладать административными правами.
// Если пользователь с указанным именем существует, то пользователь добавлен не будет.
type RepositoryAddUserOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryAddUser" json:"-"`

	//-User — Имя создаваемого пользователя.
	NewUser string `v8:"-User" json:"user"`

	//-Pwd — Пароль создаваемого пользователя.
	NewPassword string `v8:"-pwd, optional" json:"pwd"`

	//-Rights — Права пользователя. Возможные значения:
	//	ReadOnly — право на просмотр,
	//	LockObjects — право на захват объектов,
	//	ManageConfigurationVersions — право на изменение состава версий,
	//	Administration — право на административные функции.
	//
	Rights RepositoryRightType `v8:"-Rights" json:"rights"`

	//-RestoreDeletedUser — Если обнаружен удаленный пользователь с таким же именем, он будет восстановлен.
	RestoreDeletedUser bool `v8:"-RestoreDeletedUser, optional" json:"restore_deleted_user"`
}

func (o RepositoryAddUserOptions) Values() []string {

	v, _ := marshaler.Marshal(o)
	fixExtensionIndex(&v)
	return v

}

func (o RepositoryAddUserOptions) WithRepository(repository Repository) RepositoryAddUserOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

func (r Repository) AddUser(user, password string, rights RepositoryRightType, restoreDeletedUser ...bool) RepositoryAddUserOptions {

	command := RepositoryAddUserOptions{
		Designer:    NewDesigner(),
		Repository:  r,
		NewUser:     user,
		NewPassword: password,
		Rights:      rights,
	}

	if len(restoreDeletedUser) > 0 {
		command.RestoreDeletedUser = restoreDeletedUser[0]
	}

	return command

}

///ConfigurationRepositoryCopyUsers  -Path <путь> -User <Имя>
//-Pwd <Пароль> [-RestoreDeletedUser][-Extension <имя расширения>]
//— копирование пользователей из хранилища конфигурации. Копирование удаленных пользователей не выполняется. Если пользователь с указанным именем существует, то пользователь не будет добавлен.
type RepositoryCopyUsersOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryCopyUsers" json:"-"`

	//-Path — Путь к хранилищу, из которого выполняется копирование пользователей.
	RemotePath string `v8:"-Path" json:"remote_path"`

	//-User — Имя создаваемого пользователя.
	RemoteUser string `v8:"-User" json:"user"`

	//-Pwd — Пароль создаваемого пользователя.
	RemotePwd string `v8:"-Pwd, optional" json:"pwd"`

	//-RestoreDeletedUser — Если обнаружен удаленный пользователь с таким же именем, он будет восстановлен.
	RestoreDeletedUser bool `v8:"-RestoreDeletedUser, optional" json:"restore_deleted_user"`
}

func (ib RepositoryCopyUsersOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	fixExtensionIndex(&v)
	return v

}

func (o RepositoryCopyUsersOptions) WithRepository(repository Repository) RepositoryCopyUsersOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

func (o RepositoryCopyUsersOptions) FromRepository(repository Repository) RepositoryCopyUsersOptions {

	newO := o
	newO.RemotePath = repository.Path
	newO.RemoteUser = repository.User
	newO.RemotePwd = repository.Password
	return newO

}

func (r Repository) CopyUsers(path, user, password string, restoreDeletedUser ...bool) RepositoryCopyUsersOptions {

	command := RepositoryCopyUsersOptions{
		Designer:   NewDesigner(),
		Repository: r,
		RemoteUser: user,
		RemotePwd:  password,
		RemotePath: path,
	}

	if len(restoreDeletedUser) > 0 {
		command.RestoreDeletedUser = restoreDeletedUser[0]
	}

	return command

}

func (r Repository) CopyUsersFromRepository(repository Repository, restoreDeletedUser ...bool) RepositoryCopyUsersOptions {

	command := RepositoryCopyUsersOptions{
		Designer:   NewDesigner(),
		Repository: r,
		RemoteUser: repository.User,
		RemotePwd:  repository.Password,
		RemotePath: repository.Path,
	}

	if len(restoreDeletedUser) > 0 {
		command.RestoreDeletedUser = restoreDeletedUser[0]
	}

	return command

}
