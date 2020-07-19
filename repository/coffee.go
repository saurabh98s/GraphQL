package repository

import (
	"context"
	"graphql/db/models"

	"github.com/globalsign/mgo"
)


//CoffeeRepo implements the functions required by the Domain
type CoffeeRepo interface {
	Save(ctx context.Context,doc *models.Coffee) error
}

// coffeeRepo holds all the dependencies needed
type coffeeRepo struct {

	db *mgo.Database
}


func (repo *coffeeRepo)Save(ctx context.Context,doc *models.Coffee) error{
	groupError:="SAVE_ONE_COFFEE_TYPE"
	doc.CreatedAt = time.Now().Unix()

}

//NewDomainRepo creates a new Instance of the DomainRepo
func NewCoffeeRepo(db *mgo.Database, validate) CoffeeRepo {

	return &coffeeRepo{
		db:       db,
	}
}