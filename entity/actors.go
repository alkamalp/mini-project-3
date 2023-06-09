package entity

import "time"

type Actor struct {
	ID        uint      `gorm:"primary_key"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Role_id   uint      `gorm:"column:role_id"`
	Verified  int       `gorm:"column:verified"`
	Active    int       `gorm:"column:active"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:current_timestamp;autoUpdateTime"`
}
