// Handler for localhost testing
/*package handlers

import (
	"fmt"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/service"
	"net/http"
)
func SensorsHttpHandler(sensorService *service.SensorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sensors, err := sensorService.GetAllSensors(r.Context())
		if err != nil {
			http.Error(w, "Problem while loading sensors: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// HTML header
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>IoT Sensors</title>
				<style>
					table { border-collapse: collapse; width: 100%; }
					th, td { padding: 12px; text-align: left; border-bottom: 1px solid #ddd; }
					tr:hover { background-color: #f5f5f5; }
					.active-true { color: green; }
					.active-false { color: red; }
				</style>
			</head>
			<body>
				<h1>Lista Sensorów</h1>
				<table>
					<thead>
						<tr>
							<th>ID</th>
							<th>Nazwa</th>
							<th>Lokalizacja</th>
							<th>Status</th>
						</tr>
					</thead>
					<tbody>
		`)

		// Table data
		for _, sensor := range sensors {
			activeClass := "active-false"
			activeStatus := "❌ Nieaktywny"
			if sensor.Active {
				activeClass = "active-true"
				activeStatus = "✅ Aktywny"
			}

			fmt.Fprintf(w, `
				<tr>
					<td>%d</td>
					<td>%s</td>
					<td>%s</td>
					<td class="%s">%s</td>
				</tr>
			`, sensor.ID, sensor.Name, sensor.Location, activeClass, activeStatus)
		}

		fmt.Fprint(w, `
					</tbody>
				</table>
			</body>
			</html>
		`)
	}
}
*/
package handlers

import (
	"context"
	"gioui.org/app"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/service"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/storage"
	"log"
	"sync"
)

// AppState holds the application state with proper mutex protection
type AppState struct {
	Mu       sync.Mutex
	Sensors  []*storage.Sensor
	Loading  bool
	ErrorMsg string
}

// GetSensors safely retrieves sensors from the state
func (s *AppState) GetSensors() []*storage.Sensor {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	return s.Sensors
}

// IsLoading safely checks loading state
func (s *AppState) IsLoading() bool {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	return s.Loading
}

// GetErrorMsg safely retrieves error message
func (s *AppState) GetErrorMsg() string {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	return s.ErrorMsg
}

// SetLoading safely updates loading state
func (s *AppState) SetLoading(loading bool) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Loading = loading
}

// SetErrorMsg safely updates error message
func (s *AppState) SetErrorMsg(msg string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.ErrorMsg = msg
}

// setSensors safely updates sensors list
func (s *AppState) SetSensors(sensors []*storage.Sensor) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Sensors = sensors
}

// Handler
// LoadSensorData loads sensor data in a background goroutine
func LoadSensorData(service *service.SensorService, state *AppState, window *app.Window) {
	state.SetLoading(true)
	state.SetErrorMsg("")
	window.Invalidate()

	ctx := context.Background()
	sensors, err := service.GetAllSensors(ctx)

	if err != nil {
		state.SetErrorMsg("Error: " + err.Error())
	} else {
		// Convert from ent.Sensor to storage.Sensor
		storageSensors := make([]*storage.Sensor, len(sensors))
		for i, s := range sensors {
			storageSensors[i] = &storage.Sensor{
				ID:       s.ID,
				Name:     s.Name,
				Location: s.Location,
				Active:   s.Active,
			}
		}
		state.SetSensors(storageSensors)
	}

	state.SetLoading(false)
	log.Println("Finished loading sensors")
	window.Invalidate()
}
