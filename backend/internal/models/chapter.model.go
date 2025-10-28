package models

import "time"

type Chapter struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	NotebookID uint      `json:"notebookId" gorm:"index"`
	Notebook   Notebook  `json:"notebook" gorm:"foreignKey:NotebookID"`
	Files      []Notes   `json:"notes" gorm:"foreignKey:ChapterID"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
