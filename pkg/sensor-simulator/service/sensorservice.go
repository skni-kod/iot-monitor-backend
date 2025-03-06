package service

import (
	"context"

	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/ent"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/storage"
)

type SensorService struct {
	sensorRepository storage.ISensorStorage
}

func NewSensorService(storage storage.ISensorStorage) *SensorService {
	return &SensorService{sensorRepository: storage}
}

func (s *SensorService) GetSensorByID(ctx context.Context, id int) (*ent.Sensor, error) {
	return s.sensorRepository.GetByID(ctx, id)
}

func (s *SensorService) GetAllSensors(ctx context.Context) ([]*ent.Sensor, error) {
	return s.sensorRepository.GetAll(ctx)
}

// UpdateSensor

// DeleteSensor
