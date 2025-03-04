package database

import (
	ctx "context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/ent"
)

type DatabaseClient struct {
	Client *ent.Client
}

func NewDatabaseClient(driverName string, connectionString string) *DatabaseClient {
	client, err := ent.Open(driverName, connectionString)
	if err != nil {
		log.Fatalf("Failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(ctx.Background()); err != nil {
		log.Fatalf("Failed creating schema resource: %v", err)
	}

	dbClient := &DatabaseClient{Client: client}

	return dbClient
}
