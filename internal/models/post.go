package model

import (
    "time"
    "gorm.io/gorm"
)

type Post struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Title       string         `gorm:"not null" json:"title"`
    Content     string         `gorm:"type:text;not null" json:"content"`
    Slug        string         `gorm:"unique" json:"slug"`
    Status      string         `gorm:"default:draft" json:"status"` // draft | published
    UserID      uint           `json:"user_id"`
    User        User           `gorm:"foreignKey:UserID" json:"author,omitempty"`
    Comments    []Comment      `gorm:"foreignKey:PostID" json:"comments,omitempty"`
    Tags        []Tag          `gorm:"many2many:post_tags" json:"tags,omitempty"`
    ViewCount   int            `gorm:"default:0" json:"view_count"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Comment struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Content   string    `gorm:"not null" json:"content"`
    PostID    uint      `json:"post_id"`
    UserID    uint      `json:"user_id"`
    User      User      `gorm:"foreignKey:UserID" json:"author,omitempty"`
    CreatedAt time.Time `json:"created_at"`
}

type Tag struct {
    ID    uint   `gorm:"primaryKey" json:"id"`
    Name  string `gorm:"unique" json:"name"`
    Posts []Post `gorm:"many2many:post_tags" json:"-"`
}
