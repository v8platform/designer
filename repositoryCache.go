package designer

import (
	"github.com/v8platform/marshaler"
)

///ConfigurationRepositoryClearGlobalCache [-Extension <имя расширения>]
//- очистка глобального кэша версий конфигурации в хранилище.
type RepositoryClearGlobalCacheOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryClearGlobalCache" json:"-"`
}

func (ib RepositoryClearGlobalCacheOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (o RepositoryClearGlobalCacheOptions) WithRepository(repository Repository) RepositoryClearGlobalCacheOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

///ConfigurationRepositoryClearCache [-Extension <имя расширения>]
//— очистка локальной базы данных хранилища конфигурации.
type RepositoryClearCacheOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryClearCache" json:"-"`
}

func (ib RepositoryClearCacheOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (o RepositoryClearCacheOptions) WithRepository(repository Repository) RepositoryClearCacheOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

///ConfigurationRepositoryClearLocalCache [-Extension <имя расширения>]
//- очистка локального кэша версий конфигурации
type RepositoryClearLocalCacheOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryClearLocalCache" json:"-"`
}

func (ib RepositoryClearLocalCacheOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (o RepositoryClearLocalCacheOptions) WithRepository(repository Repository) RepositoryClearLocalCacheOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}
