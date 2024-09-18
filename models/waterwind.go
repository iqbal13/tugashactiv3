package models

import "time"

type WaterWind struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Water       int       `json:"water"`
	Wind        int       `json:"wind"`
	WaterStatus string    `json:"water_status"`
	WindStatus  string    `json:"wind_status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
