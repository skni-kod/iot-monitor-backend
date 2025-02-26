package main

import (
	"math"
	"math/rand/v2"
	"time"
)

type Sensor struct {
    ID          string
    Type        string
    MinValue    float64
    MaxValue    float64
    NoiseLevel  float64
}

func (s *Sensor) generateReading() float64 {
    baseValue := (s.MaxValue + s.MinValue) / 2
    noise := (rand.Float64() - 0.5) * s.NoiseLevel
    trend := math.Sin(float64(time.Now().Hour())/24.0*2*math.Pi) * (s.MaxValue-s.MinValue)/4
    
    value := baseValue + noise + trend
    // Zapewnienie, że wartość mieści się w zakresie
    return math.Max(s.MinValue, math.Min(s.MaxValue, value))
}