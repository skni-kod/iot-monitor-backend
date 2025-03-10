package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skni-kod/iot-monitor-backend/internal/common/database"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/service"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/storage"
)

func main() {
	db := database.NewDatabaseClient("sqlite3", "file:sensors_simulator.db?cache=shared&_fk=1")
	defer db.Client.Close()

	sensorRepository := storage.NewSensorStorage(db.Client)

	sensorService := service.NewSensorService(sensorRepository)

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "haloooo")
	})

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
