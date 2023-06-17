package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(scheme string, user string, passwd string, clusterName string, serverEndpoint string, dbName string) *Connection {
	c := &Connection{
		Scheme:         scheme,
		Username:       user,
		Password:       passwd,
		ClusterName:    clusterName,
		ServerEndpoint: serverEndpoint,
		DBName:         dbName,
	}
	return c
}

func (c *Connection) Connect() *Connection {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoDbServerURI := fmt.Sprintf("%s://%s:%s@%s.%s/?retryWrites=true&w=majority", c.Scheme, c.Username, c.Password, c.ClusterName, c.ServerEndpoint)
	opts := options.Client().ApplyURI(mongoDbServerURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println("ERROR: Failed to create a DB connection. Error: ", err)
		return nil
		//panic(err)
	}
	fmt.Println("DB Connection Successful")
	c.Client = client

	return c
}

func (c *Connection) Disconnect() bool {
	err := c.Client.Disconnect(context.TODO())
	if err != nil {
		fmt.Println("ERROR: Failed to Disconnect DB connection. Error: ", err)
		return false
	}
	fmt.Println("Disconnected from the DB")
	return true
}

func (c *Connection) PingDB() bool {

	if err := c.Client.Database(c.DBName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		fmt.Println("ERROR: Failed to Ping Work Log DB connection. Error: ", err)
		return false
	}
	fmt.Println("Ping to the Work Log DB was Successful!")
	return true

}

func (c *Connection) InsertNewTask(myTask Task) error {
	collection := c.Client.Database(c.DBName).Collection("tasks")
	result, err := collection.InsertOne(context.TODO(), myTask)
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	return err
}
