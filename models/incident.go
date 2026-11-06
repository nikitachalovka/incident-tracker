package models

import "time"

type Incident struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:100;not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Priority    string    `gorm:"size:20;not null" json:"priority"`
	Status      string    `gorm:"size:20;not null;default:open" json:"status"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"-"`
}
