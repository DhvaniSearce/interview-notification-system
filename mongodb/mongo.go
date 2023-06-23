package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Fields struct {
	Summary       string `bson:"summary"`
	Description   string `bson:"description"`
	CandidateName string `bson:"candidate_name"`
	StartTime     string `bson:"start_time"`
	EndTime       string `bson:"end_time"`
	EmailID       string `bson:"email_id"`
}

func Run() []Fields {
	clientOption := options.Client().ApplyURI("mongodb+srv://dhvani:9X9PkA9OcopfHlX6@cluster0.ttcc0sg.mongodb.net/?retryWrites=true&w=majority")

	const dbName = "Hackathon"
	const colName = "Interview_Information"

	var collection *mongo.Collection

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb connected successfully")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")

	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var interviewer []Fields
	if err = cur.All(context.TODO(), &interviewer); err != nil {
		fmt.Println(err)
	}
	defer cur.Close(context.Background())
	for _, i := range interviewer {
		fmt.Println(i)
	}
	// fmt.Println(interviewer[0].CandidateName)
	return interviewer
}
