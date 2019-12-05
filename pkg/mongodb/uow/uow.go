package uow

import (
	o "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository/olive"
	oo "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository/olive_oil"
	"go.mongodb.org/mongo-driver/mongo"
)

//Uow represents unit of work implementation
type Uow struct {
	client *mongo.Client
	dbName string
}

//Save completes a transaction
func (u *Uow) Save() error {
	//todo save tx
	return nil
}

// New creates new uow object
func New(client *mongo.Client, dbName string) *Uow {
	return &Uow{
		client: client,
		dbName: dbName,
	}
}

//repository implementations goes here..

//GetOliveRepository function creates a new Olive object
func (u *Uow) GetOliveRepository() *o.Olive {
	repo := o.New(u.client, u.dbName)
	return repo
}

//GetOliveOilRepository function creates a new OliveOil object
func (u *Uow) GetOliveOilRepository() *oo.OliveOil {
	repo := oo.New(u.client, u.dbName)
	return repo
}
