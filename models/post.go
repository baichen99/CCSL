package models

import uuid "github.com/iris-contrib/go.uuid"

// Post model 讨论区帖子
type Post struct {
	Base
	Creator   User      `json:"creator"`
	CreatorID uuid.UUID `gorm:"not null" json:"creatorID"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Replies   []Reply   `json:"replies"`
}

// Reply model 讨论区回复
type Reply struct {
	Base
	Creator   User      `json:"creator"`
	CreatorID uuid.UUID `gorm:"not null" json:"creatorID"`
	Content   string    `json:"content"`
}
