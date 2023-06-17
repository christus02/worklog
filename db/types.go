package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Connection struct {
	Scheme         string
	Username       string
	Password       string
	ClusterName    string
	ServerEndpoint string
	Client         *mongo.Client
	DBName         string
}

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Tag         string             `bson:"tag"`
	Summary     string             `bson:"summary"`
	Comments    []string           `bson:"comments"`
	CreatedAt   time.Time          `bson:"createdAt"`
	LastUpdated time.Time          `bson:"lastUpdated"`
}
