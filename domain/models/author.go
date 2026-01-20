package models

import "github.com/google/uuid"

type Author struct {
	Id   uuid.UUID `json:"id" binding:"required" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name" binding:"required" gorm:"type:string"`
	Bio  string    `json:"bio" binding:"required" gorm:"type:string"`
}
