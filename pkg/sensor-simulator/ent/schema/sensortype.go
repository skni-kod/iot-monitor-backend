package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SensorType holds the schema definition for the SensorType entity.
type SensorType struct {
	ent.Schema
}

// Fields of the SensorType.
func (SensorType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique().
			Comment("Name of the sensor type"),
		field.String("model").
			NotEmpty().
			Comment("Model of the sensor type"),
		field.String("manufacturer").
			Optional().
			Comment("Manufacturer of the sensor type"),
		field.String("description").
			Optional().
			Comment("Description of the sensor type"),
	}
}

// Edges of the SensorType.
func (SensorType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sensors", Sensor.Type).
			Ref("type").
			Comment("Sensors of this type"),
		
		edge.To("settings", SensorSettings.Type).
			Unique().
			Comment("Settings for this sensor type"),
	}
}
