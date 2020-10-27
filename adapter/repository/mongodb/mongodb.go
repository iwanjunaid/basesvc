package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type AuthorRepositoryReaderImpl struct {
	mdb *mongo.Database
}
