package mongo

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongo interface {
	Connect() error
	Disconnect() error
	GetClient() *mongo.Client
}

type Mongo struct {
	uri     string
	timeout int64
	client  *mongo.Client
}

func (m *Mongo) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(m.uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	m.client = client

	return nil
}

func (m *Mongo) Disconnect() error {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (m *Mongo) GetClient() *mongo.Client {
	return m.client
}

func NewMongo(uri string, timeout int64) IMongo {
	return &Mongo{uri: uri, timeout: timeout}
}
