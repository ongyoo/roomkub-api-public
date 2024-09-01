package meterModel

import (
	"time"

	"github.com/ongyoo/roomkub-api/pkg/document"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meter struct {
	StarWaterUnit        float32             `json:"star_water_unit" bson:"starWaterUnit"`
	EndWaterUnit         float32             `json:"end_water_unit" bson:"endWaterUnit"`
	StarElectricityUnit  float32             `json:"star_electricity_unit" bson:"starElectricityUnit"`
	EndElectricityUnit   float32             `json:"end_electricity_unit" bson:"endElectricityUnit"`
	TotalWaterUnit       float32             `json:"total_water_unit" bson:"totalWaterUnit"`
	TotalElectricityUnit float32             `json:"total_electricity_unit" bson:"totalElectricityUnit"`
	Documents            []document.Document `json:"documents" bson:"documents"`
	StartAt              time.Time           `json:"star_at" bson:"starAt"`
	EndAt                time.Time           `json:"end_at" bson:"endAt"`
	UpdateAt             time.Time           `json:"update_at" bson:"updateAt"`
	RecorderBy           primitive.ObjectID  `json:"recorder_by" bson:"recorderBy"` // userId
}
