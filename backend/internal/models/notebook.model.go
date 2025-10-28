package models

import "time"

type Notebook struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	UserID    uint      `json:"userId" gorm:"index"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Chapters  []Chapter `json:"chapters" gorm:"foreignKey:NotebookID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
