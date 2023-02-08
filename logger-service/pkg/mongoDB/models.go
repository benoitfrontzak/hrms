package mongoDB

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(mongo *mongo.Client) Models {
	Client = mongo

	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string `bson:"name" json:"name"`
	Data      string `bson:"data" json:"data"`
	CreatedAt string `bson:"created_at" json:"created_at"`
}

// Insert function inserts a log entry to a given collection c
// returns error if any
func (l *LogEntry) Insert(c string) (*mongo.InsertOneResult, error) {
	collection := Client.Database(DB).Collection(c)

	rowID, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      l.Name,
		Data:      l.Data,
		CreatedAt: l.CreatedAt,
	})
	if err != nil {
		log.Println("Error inserting into logs:", err)
		return rowID, err
	}

	return rowID, nil
}

func (l *LogEntry) All() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := Client.Database("logs").Collection("logs")

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Finding all docs error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []*LogEntry

	for cursor.Next(ctx) {
		var item LogEntry

		err := cursor.Decode(&item)
		if err != nil {
			log.Print("Error decoding log into slice:", err)
			return nil, err
		} else {
			logs = append(logs, &item)
		}
	}

	return logs, nil
}

func (l *LogEntry) GetOne(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := Client.Database(DB).Collection("logs")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var entry LogEntry
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (l *LogEntry) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := Client.Database(DB).Collection("logs")

	if err := collection.Drop(ctx); err != nil {
		return err
	}

	return nil
}

func (l *LogEntry) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := Client.Database(DB).Collection("logs")

	docID, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{Key: "$set", Value: bson.D{
				{"name", l.Name},
				{"data", l.Data},
			}},
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
