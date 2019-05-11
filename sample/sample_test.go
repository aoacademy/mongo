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
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Record struct {
	Name   string
	Family string
	Time   time.Time
}

type TestSuite struct {
	suite.Suite
	db *mongo.Database
}

func (suite *TestSuite) SetupSuite() {
	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb://127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019/?replicaSet=shamin",
	))
	suite.NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	suite.NoError(client.Connect(ctx))

	suite.db = client.Database("test")
}

func TestMain(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestInsert() {
	s := Record{
		Name:   "Parham",
		Family: "Alvani",
		Time:   time.Now(),
	}

	_, err := suite.db.Collection("nice").InsertOne(context.Background(), s)
	suite.NoError(err)
}

func (suite *TestSuite) TestFind() {
	s := Record{
		Name:   "Parham",
		Family: "Alvani",
		Time:   time.Now(),
	}

	_, err := suite.db.Collection("nice").InsertOne(context.Background(), s)
	suite.NoError(err)

	// Fetches the inserted data

	cur, err := suite.db.Collection("nice").Find(context.Background(), bson.M{
		"name": "Parham",
	})
	suite.NoError(err)

	for cur.Next(context.Background()) {
		var r Record

		suite.NoError(cur.Decode(&r))

		suite.T().Log(r)
	}
	suite.NoError(cur.Close(context.Background()))
}
