package main

import (
	"github.com/skni-kod/iot-monitor-backend/internal/common/database"
)

func main() {
	db := database.NewDatabaseClient("sqlite3", "file:sensors_simulator.db?cache=shared&_fk=1")
	defer db.Client.Close()
}
