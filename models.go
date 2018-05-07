package main

import "time"

// Account struct

// User strcut
type User struct {
	CreatedAt  time.Time
	ID         int
	FirstName  string `form:"first_name" json:"first_name"`
	SecondName string `form:"second_name" json:"second_name"`
	Password   string `form:"password" json:"password"`
	Email      string `form:"email" json:"email"`
	Posts      []Post
	Groups     []Group `gorm:"many2many:user_groups;"`
	Friends    []Friend
}

// Friend struct
type Friend struct {
	ID       int
	UserID   int
	FriendID int
}

// Post struct
type Post struct {
	CreatedAt time.Time
	ID        int
	UserID    int
	GroupID   int
	Content   string
}

// Group struct
type Group struct {
	CreatedAt time.Time
	ID        int
	AdminID   int
	Name      string
	Posts     []Post
}
