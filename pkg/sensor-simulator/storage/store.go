package storage

import (
	"context"

	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/ent"
	"github.com/skni-kod/iot-monitor-backend/pkg/sensor-simulator/ent/sensor"
)

type Sensor struct {
	ID         int
	Name       string
	Location   string
	Active     bool
	SensorType int
}

type ISensorStorage interface {
	GetAll(ctx context.Context) ([]*ent.Sensor, error)
	GetByID(ctx context.Context, id int) (*ent.Sensor, error)
}

type SensorStorage struct {
	client *ent.Client
}

func NewSensorStorage(client *ent.Client) ISensorStorage {
	return &SensorStorage{client: client}
}

func (s *SensorStorage) GetAll(ctx context.Context) ([]*ent.Sensor, error) {
	return s.client.Sensor.Query().All(ctx)
}

func (s *SensorStorage) GetByID(ctx context.Context, id int) (*ent.Sensor, error) {
	return s.client.Sensor.Query().Where(sensor.ID(id)).Only(ctx)
}

// UpdateSensor

// DeleteSensor
