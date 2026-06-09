package model

import (
    "time"
    "gorm.io/gorm"
)


type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `gorm:"not null" json:"name"`
    Email     string         `gorm:"unique;not null" json:"email"`
    Password  string         `gorm:"not null" json:"-"`  // json:"-" یعنی پسورد هیچوقت برنمیگرده
    Role      string         `gorm:"default:user" json:"role"` // user | admin
    Posts     []Post         `gorm:"foreignKey:UserID" json:"posts,omitempty"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // soft delete
}