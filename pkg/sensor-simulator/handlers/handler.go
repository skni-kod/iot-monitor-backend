package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/service"
)

type SensorHandler struct {
	sensorService service.ISensorService
}

func NewSensorHandler(sensorService service.ISensorService) *SensorHandler {
	return &SensorHandler{sensorService: sensorService}
}

func (h *SensorHandler) GetSensor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}

	sensor, err := h.sensorService.GetSensorByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sensor not found"})
		return
	}

	c.JSON(http.StatusOK, sensor)
}

func (h *SensorHandler) GetAllSensors(c *gin.Context) {
	sensors, err := h.sensorService.GetAllSensors(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sensors"})
		return
	}

	c.JSON(http.StatusOK, sensors)
}
