/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 17-04-2018
 * |
 * | File Name:     sample_test.go
 * +===============================================
 */

package sample

import (
	"context"
	"log"
	"testing"
	"time"

	mgo "github.com/mongodb/mongo-go-driver/mongo"
)

var testDB *mgo.Database

func setupDB() {
	client, err := mgo.NewClient("mongodb://127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019/?replicaSet=shamin")
	if err != nil {
		log.Fatal(err)
	}

	testDB = client.Database("test")
}

func TestInsert(t *testing.T) {
	setupDB()

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
}
