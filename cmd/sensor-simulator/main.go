// Localhost testing
/*
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skni-kod/iot-monitor-backend/internal/common/database"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/handlers"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/service"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/storage"
	"log"
	"net/http"
)

func main() {
	// Database initialization
	db := database.NewDatabaseClient("sqlite3", "file:sensors_simulator.db?cache=shared&_fk=1")

	// Database info for debugging
	//db.Client = db.Client.Debug()
	defer db.Client.Close()

	// Sensors' storage and service initialization
	sensorStorage := storage.NewSensorStorage(db.Client)
	sensorService := service.NewSensorService(sensorStorage)

	// Activating handler for listening
	http.HandleFunc("/sensors", handlers.SensorsHttpHandler(sensorService))

	// Launching local server
	log.Println("Check the server at: http://localhost:8080/sensors")
	log.Fatal(http.ListenAndServe(":8080", nil))
}*/

package main

import (
	// Gio packages
	"fmt"
	"image/color"
	"log"
	"net/http"
	"strconv"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	// Generated packages
	"github.com/skni-kod/iot-monitor-backend/internal/common/database"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/handlers"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/service"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/storage"
	// Default packages
)

func main() {
	// Window
	w := new(app.Window)
	w.Option(
		app.Title("IoT Monitor"),
		app.Size(unit.Dp(800), unit.Dp(600)),
	)

	// Displaying function
	go func() {
		if err := run(w); err != nil {
			log.Fatal(err)
		}
	}()

	app.Main()
}

func run(window *app.Window) error {
	// App state (Gio)
	state := &handlers.AppState{}
	// Theme (Gio)
	theme := material.NewTheme()

	// Sqlite database
	db := database.NewDatabaseClient("sqlite3", "file:sensors_simulator.db?cache=shared&_fk=1")
	defer db.Client.Close()

	// sensorRepository := storage.NewSensorStorage(db.Client)

	// sensorService := service.NewSensorService(sensorRepository)

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "haloooo")
	})

	fmt.Printf("Starting HTTP server at port: %s", "8000")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
	sensorStorage := storage.NewSensorStorage(db.Client)
	sensorService := service.NewSensorService(sensorStorage)

	// Widgets (Gio)
	var (
		list = &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		}
		launchButton widget.Clickable
	)

	// Mock
	/*go func() {
		mockSensors := []*storage.Sensor{
			{ID: 1, Name: "Sensor ...", Active: true},
			{ID: 2, Name: "Sensor ...", Active: false},
			{ID: 3, Name: "Sensor ...", Active: true},
		}

		state.SetLoading(true)
		window.Invalidate()

		state.SetSensors(mockSensors)
		state.SetLoading(false)
		window.Invalidate()
	}()*/

	// List of operations + default window events handling (Gio)
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if launchButton.Clicked(gtx) {
				go handlers.LoadSensorData(sensorService, state, window)
			}

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceBetween,
			}.Layout(gtx,
				// Subtitle
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					title := material.H6(theme, "IoT Sensors Monitor")
					title.Alignment = text.Middle
					title.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
					return title.Layout(gtx)
				}),
				// Launch button
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					btn := material.Button(theme, &launchButton, "Launch")
					btn.Background = color.NRGBA{R: 15, G: 15, B: 15, A: 255}
					btn.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
					return btn.Layout(gtx)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					if state.IsLoading() {
						return material.Loader(theme).Layout(gtx)
					}

					errorMsg := state.GetErrorMsg()
					if errorMsg != "" {
						return material.Body1(theme, errorMsg).Layout(gtx)
					}

					sensors := state.GetSensors()
					return material.List(theme, list).Layout(gtx, len(sensors), func(gtx layout.Context, index int) layout.Dimensions {
						sensor := sensors[index]
						return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{
								Axis: layout.Vertical,
							}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return material.Body1(theme, "ID: "+strconv.Itoa(int(sensor.ID))).Layout(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return material.Body2(theme, "Name: "+sensor.Name).Layout(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									if sensor.Location != "" {
										return material.Body2(theme, "Location: "+sensor.Location).Layout(gtx)
									}
									return layout.Dimensions{}
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									status := "❌ Inactive"
									if sensor.Active {
										status = "✅ Active"
									}

									return material.Body2(theme, "Status: "+status).Layout(gtx)
								}),
							)
						})
					})
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
