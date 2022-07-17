package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Robot struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name       string    `json:"name" gorm:"type:varchar(100);uniqueIndex"`
	Generation uint      `json:"generation" gorm:"type:smallint"`
}

type Flight struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	RobotID   uuid.UUID `json:"robotId" gorm:"type:uuid REFERENCES robots(ID);uniqueIndex:idx_flight"`
	StartTime time.Time `json:"startTime" gorm:"uniqueIndex:idx_flight"`
	EndTime   time.Time `json:"endTime" gorm:"uniqueIndex:idx_flight"`
	Lat       float64   `json:"lat" gorm:"type:decimal(10,8);uniqueIndex:idx_flight"`
	Lng       float64   `json:"lng" gorm:"type:decimal(11,8);uniqueIndex:idx_flight"`
}
