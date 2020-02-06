package db

import (
	"context"
	"github.com/yehorshapanov/BitmediaGo/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const USER_TABLE string = "user" // the collection name

type User struct {
	ID			primitive.ObjectID  `bson:"_id,omitempty"`
	Email		string 				`bson:"email" json:"email"`
	LastName	string 				`bson:"last_name" json:"last_name"`
	Country		string 				`bson:"country" json:"country"`
	City 		string 				`bson:"city" json:"city"`
	Gender 		string 				`bson:"gender" json:"gender"`
	BirthDate 	string 				`bson:"birth_date" json:"birth_date"`
}

func (mds *MongoDBStorer) Delete(ctx context.Context, id string) (err error) {
	return
}

func (mds *MongoDBStorer) Create(ctx context.Context, s User) (err error) {
	collection := mds.DB.Collection(USER_TABLE)
	_, err = collection.InsertOne(ctx, s)
	if err != nil {
		logger.Get().Infof("Error inserting user: %v", s)
		return
	}

	logger.Get().Info("Inserted user into collection")
	return
}

func (mds *MongoDBStorer) ListAllUsers(ctx context.Context) (users []User, err error) {
	collection := mds.DB.Collection(USER_TABLE)

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	var elem User
	for cur.Next(ctx) {
		err = cur.Decode(&elem)
		users = append(users, elem)
	}
	if err = cur.Err(); err != nil {
		logger.Get().Infof("Error in listing data: ", err)
		return
	}
	return
}

func Paginate(collection *mongo.Collection, startValue primitive.ObjectID, nPerPage int64) ([]User, *User, error) {
	filter := bson.D{
		{"userid", startValue},
	}

	options := options.Find()
	options.SetLimit(nPerPage)

	cursor, _ := collection.Find(context.Background(), filter, options)

	var lastValue *User
	var results []User
	for cursor.Next(context.Background()) {
		var elem User
		err := cursor.Decode(elem)
		if err != nil {
			return results, lastValue, err
		}
		results = append(results, elem)
		lastValue = &elem
	}

	return results, lastValue, nil
}
