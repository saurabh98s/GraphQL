package instance

import (
	"graphql/config"
	"graphql/logger"
	"sync"

	"github.com/globalsign/mgo"
)

type instance struct {
	db *mgo.Database
}

var singleton = &instance{}

var once sync.Once

// Init initializes the instance
func Init() {
	mongoConfig := config.Mongo()

	once.Do(func() {
		logger.Log.Info("Connecting to mongodb...")
		session, err := mgo.Dial("mongodb://127.0.0.1:27017")
		if err != nil {
			logger.Log.Fatal(err)
		}
		singleton.db = session.DB(mongoConfig.Database())

		logger.Log.Info("Connected to mongodb successfully...")
	})

}

// DB returns the database object
func DB() *mgo.Database {
	return singleton.db
}