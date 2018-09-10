/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 17-04-2018
 * |
 * | File Name:     sample_test.go
 * +===============================================
 */

package mongo

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	mgo "github.com/mongodb/mongo-go-driver/mongo"
)

var testDB *mgo.Database

func setupDB() {
	client, err := mgo.NewClient("mongodb://127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019/?replicaSet=shamin")
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.Background()); err != nil {
		log.Fatalf("DB connection error: %s", err)
	}

	testDB = client.Database("test")
}

func TestMain(t *testing.T) {
	setupDB()

	t.Run("TestInsert", func(t *testing.T) {
		s := struct {
			Name   string
			Family string
			Time   time.Time
		}{
			Name:   "Parham",
			Family: "Alvani",
			Time:   time.Now(),
		}
		t.Log(s)

		if _, err := testDB.Collection("nice").InsertOne(context.Background(), s); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("TestFind", func(t *testing.T) {
		cur, err := testDB.Collection("nice").Find(context.Background(), bson.NewDocument(
			bson.EC.String("name", "Parham"),
		))
		if err != nil {
			t.Fatal(err)
		}

		for cur.Next(context.Background()) {
			elem := bson.NewDocument()

			if err := cur.Decode(elem); err != nil {
				t.Fatal(err)
			}

			t.Log(elem)
		}
		if err := cur.Close(context.Background()); err != nil {
			t.Fatal(err)
		}
	})
}
