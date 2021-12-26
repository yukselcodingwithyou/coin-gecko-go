package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Client struct {
	Uri string
	DB  *mongo.Client
}

func New(uri string) *Client {
	return &Client{Uri: uri}
}

func (c *Client) Connect() *Client {
	client := c.connect()
	c.ping(client)
	c.DB = client
	return c
}

func (c *Client) Collection(dbName string, collectionName string) *mongo.Collection {
	return c.DB.Database(dbName).Collection(collectionName)
}

func (c *Client) connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(c.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (c *Client) ping(client *mongo.Client) {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Connected to MongoDB on %s\n", c.Uri)
	}
}
