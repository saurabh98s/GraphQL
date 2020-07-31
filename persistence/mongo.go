package persistence

import (
	"graphql/logger"
	"time"

	"github.com/globalsign/mgo"
)


const (
	DB="graphql"
	USERS="users"
	COLLECTION="demo"
)


type MongoDBLayer struct {
	session *mgo.Session
}


func NewMongoDBLayer(connection string)  {
	_,err:=mgo.Dial("mongodb://127.0.0.1:27017")
	if err!=nil {
		logger.Log.WithTime(time.Now().Local()).Info(err)
		return
	}
}