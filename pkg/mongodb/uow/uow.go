package uow

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

const oliveCollectionName = "olives"

//Uow represents unit of work implementation
type Uow struct {
	client          *mongo.Client
	dbName          string
	oliveRepository repository.OliveRepository
	//other repo interfaces goes here
}

//OliveRepository function creates a new Olive object
func (u *Uow) OliveRepository() repository.OliveRepository {

	repo := repository.Olive{}
	repo.DBName = u.dbName
	repo.CollectionName = oliveCollectionName

	return &repo
}

//Save completes a transaction
func (u *Uow) Save() error {
	return nil
}

// New creates new uow objecte
func New(client *mongo.Client, dbName string) *Uow {
	return &Uow{
		client: client,
		dbName: dbName,
	}
}
