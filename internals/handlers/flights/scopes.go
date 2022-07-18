package flightsHandler

import (
	"time"

	"gorm.io/gorm"
)

func GenerationEquals(generation int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("left join robots on flights.robot_id = robots.id").Where("generation = ?", generation)
	}
}

func StartingFrom(dateFrom time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("start_time >= ?", dateFrom)
	}
}

func EndingIn(dateTo time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("end_time <= ?", dateTo)
	}
}
