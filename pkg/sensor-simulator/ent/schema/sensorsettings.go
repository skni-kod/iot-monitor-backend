package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SensorSettings holds the schema definition for the SensorSettings entity.
type SensorSettings struct {
	ent.Schema
}

// Fields of the SensorSettings.
func (SensorSettings) Fields() []ent.Field {
	return []ent.Field{
		field.Int("polling_interval_seconds").
			Default(60).
			Positive().
			Comment("How often the simulated sensor should generate new data"),
		field.Float("min_value").
			Optional().
			Comment("Minimum possible value for this sensor type (for simulation)"),
		field.Float("max_value").
			Optional().
			Comment("Maximum possible value for this sensor type (for simulation)"),
		field.Float("mean_value").
			Optional().
			Comment("Mean value for normal distribution in simulation"),
		field.Float("std_deviation").
			Optional().
			Comment("Standard deviation for normal distribution in simulation"),
		field.Float("drift_factor").
			Default(0.0).
			Comment("Gradual drift to apply to values over time"),
		field.Bool("seasonal_variation").
			Default(false).
			Comment("Whether to apply seasonal/cyclic variations"),
		field.Int("cycle_period_seconds").
			Optional().
			Comment("Period of cyclic variation in seconds"),
		field.Float("noise_factor").
			Default(0.05).
			Comment("Amount of random noise to add to readings (0-1)"),
		field.JSON("possible_values", []string{}).
			Optional().
			Comment("Possible values for categorical sensors"),
		field.JSON("anomaly_config", map[string]interface{}{}).
			Optional().
			Comment("Configuration for simulating anomalies"),
	}
}

// Edges of the SensorSettings.
func (SensorSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sensor_type", SensorType.Type).
			Ref("settings").
			Unique().
			Required().
			Comment("The sensor type these settings apply to"),
	}
}
