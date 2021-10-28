package mongo

import (
	"context"
	"errors"

	"git.cyradar.com/phinc/my-awesome-project/internal/model"
	"git.cyradar.com/phinc/my-awesome-project/internal/persistance"
	"git.cyradar.com/utilities/data/providers/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var staffCollection = "staffs"

type StaffRepository struct {
	*Repository
}

func NewMongoStaffRepository(provider *mongo.MongoProvider) persistance.StaffRepository {
	repo := &StaffRepository{&Repository{provider, staffCollection}}
	collection, close := repo.collection()
	defer close()

	collection.EnsureIndex(mgo.Index{
		Key: []string{
			"id",
		},
	})

	return repo
}

func (repo *StaffRepository) Insert(ctx context.Context, staff *model.Staff) error {
	collection, close := repo.collection()
	defer close()

	err := collection.Insert(staff)
	return repo.provider.NewError(err)
}

func (repo *StaffRepository) Update(ctx context.Context, id string, update *model.Staff) error {
	collection, close := repo.collection()
	defer close()

	err := collection.UpdateId(id, update)
	return repo.provider.NewError(err)
}

func (repo *StaffRepository) FindById(ctx context.Context, id string) (model.Staff, error) {
	collection, close := repo.collection()
	defer close()

	var t model.Staff
	return t, repo.provider.NewError(collection.FindId(bson.ObjectIdHex(id)).One(&t))
}

func (repo *StaffRepository) Remove(ctx context.Context, id string) error {
	collection, close := repo.collection()
	defer close()

	if !bson.IsObjectIdHex(id) {
		return repo.provider.NewError(errors.New("invalid staff id"))
	}
	return repo.provider.NewError(collection.RemoveId(bson.ObjectIdHex(id)))
}
