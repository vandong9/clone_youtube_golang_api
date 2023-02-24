package models

import "time"

type LoginToken struct {
	UserFullName string    `json:"full_name"`
	Time         time.Time `json:"time"`
	Duration     int       `json:"duration"`
}
