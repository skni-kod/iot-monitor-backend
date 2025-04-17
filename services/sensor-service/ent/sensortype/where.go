// Code generated by ent, DO NOT EDIT.

package sensortype

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/skni-kod/iot-monitor-backend/services/sensor-service/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SensorType {
	return predicate.SensorType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SensorType {
	return predicate.SensorType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SensorType {
	return predicate.SensorType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SensorType {
	return predicate.SensorType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SensorType {
	return predicate.SensorType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SensorType {
	return predicate.SensorType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SensorType {
	return predicate.SensorType(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldName, v))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldModel, v))
}

// Manufacturer applies equality check predicate on the "manufacturer" field. It's identical to ManufacturerEQ.
func Manufacturer(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldManufacturer, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldDescription, v))
}

// Unit applies equality check predicate on the "unit" field. It's identical to UnitEQ.
func Unit(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldUnit, v))
}

// MinValue applies equality check predicate on the "min_value" field. It's identical to MinValueEQ.
func MinValue(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldMinValue, v))
}

// MaxValue applies equality check predicate on the "max_value" field. It's identical to MaxValueEQ.
func MaxValue(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldMaxValue, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContainsFold(FieldName, v))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLTE(FieldModel, v))
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContains(FieldModel, v))
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasPrefix(FieldModel, v))
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasSuffix(FieldModel, v))
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEqualFold(FieldModel, v))
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContainsFold(FieldModel, v))
}

// ManufacturerEQ applies the EQ predicate on the "manufacturer" field.
func ManufacturerEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldManufacturer, v))
}

// ManufacturerNEQ applies the NEQ predicate on the "manufacturer" field.
func ManufacturerNEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNEQ(FieldManufacturer, v))
}

// ManufacturerIn applies the In predicate on the "manufacturer" field.
func ManufacturerIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldIn(FieldManufacturer, vs...))
}

// ManufacturerNotIn applies the NotIn predicate on the "manufacturer" field.
func ManufacturerNotIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNotIn(FieldManufacturer, vs...))
}

// ManufacturerGT applies the GT predicate on the "manufacturer" field.
func ManufacturerGT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGT(FieldManufacturer, v))
}

// ManufacturerGTE applies the GTE predicate on the "manufacturer" field.
func ManufacturerGTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGTE(FieldManufacturer, v))
}

// ManufacturerLT applies the LT predicate on the "manufacturer" field.
func ManufacturerLT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLT(FieldManufacturer, v))
}

// ManufacturerLTE applies the LTE predicate on the "manufacturer" field.
func ManufacturerLTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLTE(FieldManufacturer, v))
}

// ManufacturerContains applies the Contains predicate on the "manufacturer" field.
func ManufacturerContains(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContains(FieldManufacturer, v))
}

// ManufacturerHasPrefix applies the HasPrefix predicate on the "manufacturer" field.
func ManufacturerHasPrefix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasPrefix(FieldManufacturer, v))
}

// ManufacturerHasSuffix applies the HasSuffix predicate on the "manufacturer" field.
func ManufacturerHasSuffix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasSuffix(FieldManufacturer, v))
}

// ManufacturerIsNil applies the IsNil predicate on the "manufacturer" field.
func ManufacturerIsNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldIsNull(FieldManufacturer))
}

// ManufacturerNotNil applies the NotNil predicate on the "manufacturer" field.
func ManufacturerNotNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldNotNull(FieldManufacturer))
}

// ManufacturerEqualFold applies the EqualFold predicate on the "manufacturer" field.
func ManufacturerEqualFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEqualFold(FieldManufacturer, v))
}

// ManufacturerContainsFold applies the ContainsFold predicate on the "manufacturer" field.
func ManufacturerContainsFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContainsFold(FieldManufacturer, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContainsFold(FieldDescription, v))
}

// UnitEQ applies the EQ predicate on the "unit" field.
func UnitEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldUnit, v))
}

// UnitNEQ applies the NEQ predicate on the "unit" field.
func UnitNEQ(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNEQ(FieldUnit, v))
}

// UnitIn applies the In predicate on the "unit" field.
func UnitIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldIn(FieldUnit, vs...))
}

// UnitNotIn applies the NotIn predicate on the "unit" field.
func UnitNotIn(vs ...string) predicate.SensorType {
	return predicate.SensorType(sql.FieldNotIn(FieldUnit, vs...))
}

// UnitGT applies the GT predicate on the "unit" field.
func UnitGT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGT(FieldUnit, v))
}

// UnitGTE applies the GTE predicate on the "unit" field.
func UnitGTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldGTE(FieldUnit, v))
}

// UnitLT applies the LT predicate on the "unit" field.
func UnitLT(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLT(FieldUnit, v))
}

// UnitLTE applies the LTE predicate on the "unit" field.
func UnitLTE(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldLTE(FieldUnit, v))
}

// UnitContains applies the Contains predicate on the "unit" field.
func UnitContains(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContains(FieldUnit, v))
}

// UnitHasPrefix applies the HasPrefix predicate on the "unit" field.
func UnitHasPrefix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasPrefix(FieldUnit, v))
}

// UnitHasSuffix applies the HasSuffix predicate on the "unit" field.
func UnitHasSuffix(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldHasSuffix(FieldUnit, v))
}

// UnitIsNil applies the IsNil predicate on the "unit" field.
func UnitIsNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldIsNull(FieldUnit))
}

// UnitNotNil applies the NotNil predicate on the "unit" field.
func UnitNotNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldNotNull(FieldUnit))
}

// UnitEqualFold applies the EqualFold predicate on the "unit" field.
func UnitEqualFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldEqualFold(FieldUnit, v))
}

// UnitContainsFold applies the ContainsFold predicate on the "unit" field.
func UnitContainsFold(v string) predicate.SensorType {
	return predicate.SensorType(sql.FieldContainsFold(FieldUnit, v))
}

// MinValueEQ applies the EQ predicate on the "min_value" field.
func MinValueEQ(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldMinValue, v))
}

// MinValueNEQ applies the NEQ predicate on the "min_value" field.
func MinValueNEQ(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldNEQ(FieldMinValue, v))
}

// MinValueIn applies the In predicate on the "min_value" field.
func MinValueIn(vs ...float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldIn(FieldMinValue, vs...))
}

// MinValueNotIn applies the NotIn predicate on the "min_value" field.
func MinValueNotIn(vs ...float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldNotIn(FieldMinValue, vs...))
}

// MinValueGT applies the GT predicate on the "min_value" field.
func MinValueGT(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldGT(FieldMinValue, v))
}

// MinValueGTE applies the GTE predicate on the "min_value" field.
func MinValueGTE(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldGTE(FieldMinValue, v))
}

// MinValueLT applies the LT predicate on the "min_value" field.
func MinValueLT(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldLT(FieldMinValue, v))
}

// MinValueLTE applies the LTE predicate on the "min_value" field.
func MinValueLTE(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldLTE(FieldMinValue, v))
}

// MinValueIsNil applies the IsNil predicate on the "min_value" field.
func MinValueIsNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldIsNull(FieldMinValue))
}

// MinValueNotNil applies the NotNil predicate on the "min_value" field.
func MinValueNotNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldNotNull(FieldMinValue))
}

// MaxValueEQ applies the EQ predicate on the "max_value" field.
func MaxValueEQ(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldMaxValue, v))
}

// MaxValueNEQ applies the NEQ predicate on the "max_value" field.
func MaxValueNEQ(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldNEQ(FieldMaxValue, v))
}

// MaxValueIn applies the In predicate on the "max_value" field.
func MaxValueIn(vs ...float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldIn(FieldMaxValue, vs...))
}

// MaxValueNotIn applies the NotIn predicate on the "max_value" field.
func MaxValueNotIn(vs ...float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldNotIn(FieldMaxValue, vs...))
}

// MaxValueGT applies the GT predicate on the "max_value" field.
func MaxValueGT(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldGT(FieldMaxValue, v))
}

// MaxValueGTE applies the GTE predicate on the "max_value" field.
func MaxValueGTE(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldGTE(FieldMaxValue, v))
}

// MaxValueLT applies the LT predicate on the "max_value" field.
func MaxValueLT(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldLT(FieldMaxValue, v))
}

// MaxValueLTE applies the LTE predicate on the "max_value" field.
func MaxValueLTE(v float64) predicate.SensorType {
	return predicate.SensorType(sql.FieldLTE(FieldMaxValue, v))
}

// MaxValueIsNil applies the IsNil predicate on the "max_value" field.
func MaxValueIsNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldIsNull(FieldMaxValue))
}

// MaxValueNotNil applies the NotNil predicate on the "max_value" field.
func MaxValueNotNil() predicate.SensorType {
	return predicate.SensorType(sql.FieldNotNull(FieldMaxValue))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SensorType {
	return predicate.SensorType(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SensorType {
	return predicate.SensorType(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SensorType {
	return predicate.SensorType(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SensorType {
	return predicate.SensorType(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SensorType {
	return predicate.SensorType(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SensorType {
	return predicate.SensorType(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SensorType {
	return predicate.SensorType(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SensorType {
	return predicate.SensorType(sql.FieldLTE(FieldCreatedAt, v))
}

// HasSensors applies the HasEdge predicate on the "sensors" edge.
func HasSensors() predicate.SensorType {
	return predicate.SensorType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SensorsTable, SensorsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSensorsWith applies the HasEdge predicate on the "sensors" edge with a given conditions (other predicates).
func HasSensorsWith(preds ...predicate.Sensor) predicate.SensorType {
	return predicate.SensorType(func(s *sql.Selector) {
		step := newSensorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SensorType) predicate.SensorType {
	return predicate.SensorType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SensorType) predicate.SensorType {
	return predicate.SensorType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SensorType) predicate.SensorType {
	return predicate.SensorType(sql.NotPredicates(p))
}
