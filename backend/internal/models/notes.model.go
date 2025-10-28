package models

import "time"

type Notes struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Content   string    `json:"content" gorm:"type:text"`
	ChapterID uint      `json:"chapterId" gorm:"index"`
	Chapter   Chapter   `json:"chapter" gorm:"foreignKey:ChapterID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
