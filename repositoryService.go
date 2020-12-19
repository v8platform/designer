package designer

import (
	"github.com/v8platform/marshaler"
)

type GroupByType string

func (t GroupByType) MarshalV8() (string, error) {
	return string(t), nil
}

const (
	REPOSITORY_GROUP_BY_OBJECT  GroupByType = "-GroupByObject"
	REPOSITORY_GROUP_BY_COMMENT GroupByType = "-GroupByComment"
)

//ConfigurationRepositoryReport
///ConfigurationRepositoryReport [-Extension <имя расширения>] <имя файла>
//[-NBegin <номер версии>] [-NEnd <номер версии>] [-GroupByObject] [-GroupByComment]
//— построение отчета по истории хранилища.
//Если параметры группировки не указаны и режим совместимости указан "Не используется",
//то отчет формируется с группировкой по версиям.
//В режимах совместимости "Версия 8.1" и "Версия 8.2.13" отчет формируется с группировкой по объектам.
//Если конфигурация базы данных отличается от редактируемой по свойству совместимости,
//при обработке командной строки учитывается значение режима совместимости конфигурации базы данных.

//Примеры:
//для конфигурации, не присоединенной к текущему хранилищу:
//DESIGNER /F"D:\V8\Cfgs82\ИБ82" /ConfigurationRepositoryF "D:\V8\Cfgs82" /ConfigurationRepositoryN "Администратор" /ConfigurationRepositoryReport "D:\ByObject.mxl" -NBegin 1 -NEnd 2 –GroupByObject
//для присоединенной к хранилищу конфигурации, информация для отчетов берется из текущего хранилища:
//DESIGNER /F"D:\V8\Cfgs82\ИБ82" /ConfigurationRepositoryReport "D:\ByComment.mxl" -NBegin 1 -NEnd 2 -GroupByComment
type RepositoryReportOptions struct {
	Designer   `v8:",inherit" json:"designer"`
	Repository `v8:",inherit" json:"repository"`

	File string `v8:"/ConfigurationRepositoryReport" json:"file"`

	//NBegin — номер сохраненной версии, от которой начинается строиться отчет;
	NBegin int64 `v8:"-NBegin, optional" json:"number_begin"`

	//NEnd — номер сохраненной версии, по которую строится отчет;
	NEnd int64 `v8:"-NEnd, optional" json:"number_end"`

	//GroupByObject — признак формирования отчета по версиям с группировкой по объектам;
	//GroupByComment — признак формирования отчета по версиям с группировкой по комментарию.
	GroupBy GroupByType `v8:", optional" json:"group_by"`
}

func (ib RepositoryReportOptions) Values() []string {

	v, _ := marshaler.Marshal(ib)
	fixExtensionIndex(&v)
	return v

}

func (o RepositoryReportOptions) GroupByObject() RepositoryReportOptions {

	newO := o
	newO.GroupBy = REPOSITORY_GROUP_BY_OBJECT
	return newO

}

func (o RepositoryReportOptions) GroupByComment() RepositoryReportOptions {

	newO := o
	newO.GroupBy = REPOSITORY_GROUP_BY_COMMENT
	return newO

}

func (o RepositoryReportOptions) WithRepository(repository Repository) RepositoryReportOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

func (r Repository) Report(file string, startAndEndVersions ...int64) RepositoryReportOptions {

	command := RepositoryReportOptions{
		Designer:   NewDesigner(),
		Repository: r,
		File:       file,
	}

	if len(startAndEndVersions) > 0 {
		command.NBegin = startAndEndVersions[0]
		if len(startAndEndVersions) > 2 {
			command.NEnd = startAndEndVersions[1]
		}
	}

	return command

}
