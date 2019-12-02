package uow

import (
	o "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository/olive"
	oo "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository/olive_oil"
)

//Service represents unit of work interface
type Service interface {
	OliveRepository() o.Repository
	OliveOilRepository() oo.Repository
	Save() error
}
