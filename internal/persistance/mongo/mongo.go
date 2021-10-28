package mongo

import (
	"log"

	"git.cyradar.com/phinc/my-awesome-project/internal/config"
	"git.cyradar.com/phinc/my-awesome-project/internal/persistance"
	"git.cyradar.com/utilities/data/providers/mongo"
	"gopkg.in/mgo.v2"
)

func (repo *Repository) collection() (collection *mgo.Collection, close func()) {
	session := repo.provider.MongoClient().GetCopySession()
	close = session.Close

	return session.DB(repo.provider.MongoClient().Database()).C(repo.collectionName), close
}

type Repository struct {
	provider       *mongo.MongoProvider
	collectionName string
}

func NewMongoDB() *persistance.Repository {
	mongoDB := new(persistance.Repository)

	env := config.Env

	mongoProvider := mongo.NewMongoProviderFromURL(env.MongoURL)
	log.Println("Connected to MongoDB")

	mongoDB.StaffRepository = NewMongoStaffRepository(mongoProvider)

	return mongoDB
}
