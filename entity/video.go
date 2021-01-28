// Package entity impls api resource entities
package entity

import "time"

// Person struct defines fields in person resource
type Person struct {
	ID        uint64 `json:"id" gorm:"primaryKey,autoIncrement"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	Age       uint8  `json:"age" binding:"gte=0,lte=130"`
	Email     string `json:"email" validate:"required,email" gorm:"type:varchar(256);UNIQUE"`
}

// Video struct defines fields in video resource
type Video struct {
	ID          uint64    `json:"id" gorm:"primaryKey,autoIncrement"`
	Title       string    `json:"title" binding:"min=2,max=100" validate:"cool" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
