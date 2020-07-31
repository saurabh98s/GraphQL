package repository

import (
	"context"
	"errors"
	"graphql/db/models"
	"graphql/logger"
	"time"

	"github.com/globalsign/mgo"
)

//CoffeeRepo implements the functions required by the Domain
type CoffeeRepo interface {
	Save(ctx context.Context, doc *models.Coffee) (*models.Coffee,error)
}

// coffeeRepo holds all the dependencies needed
type coffeeRepo struct {
	db *mgo.Database
}

func (repo *coffeeRepo) Save(ctx context.Context, doc *models.Coffee) (*models.Coffee,error) {
	groupError := "SAVE_ONE_COFFEE_TYPE"
	doc.CreatedAt = int(time.Now().Unix())
	if *doc.Name==" " {
		logger.Log.Error(groupError)
		return nil,errors.New("[ERROR] Empty domain name")
	}
	logger.Log.Info("Adding coffee type: ",doc.Name)
	err:=repo.db.C("coffeeshop").Insert(doc)
	if err!=nil {
		logger.Log.WithError(err).Error(groupError)
		return nil,err
	}
	return doc,nil

}

//NewCoffeeRepo creates a new Instance of the DomainRepo
func NewCoffeeRepo(db *mgo.Database) CoffeeRepo {

	return &coffeeRepo{
		db: db,
	}
}
