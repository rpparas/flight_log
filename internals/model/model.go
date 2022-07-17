package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Robot struct {
	gorm.Model
	ID         uuid.UUID `json:"robotId" gorm:"primaryKey;type:uuid"`
	Name       string    `json:"name" gorm:"type:varchar(100);unique_index"`
	Generation uint      `json:"generation" gorm:"type:smallint"`
}

type Flight struct {
	gorm.Model
	ID uuid.UUID `json:"flightId" gorm:"primaryKey; type:uuid"`
	// RobotID   uuid.UUID
	// Robot     Robot     `json:"robot" gorm:"referencesRobotID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Lat       float64   `json:"lat" gorm:"type:decimal(10,8)"`
	Lng       float64   `json:"lng" gorm:"type:decimal(11,8)"`
}
