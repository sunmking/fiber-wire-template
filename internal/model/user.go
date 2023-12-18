package model

import "fiber-wire-template/pkg/util/table"

type User struct {
	Id        uint   `gorm:"primarykey"`
	UserId    string `gorm:"unique;not null"`
	Username  string `gorm:"unique;not null"`
	Nickname  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Skill     string `gorm:"not null"`
	Sex       int    `gorm:"type:tinyint(1);not null"`
	Status    int    `gorm:"type:tinyint(1);not null"`
	Address   string `gorm:"not null"`
	Mobile    string `gorm:"not null"`
	QQ        string `gorm:"not null"`
	Email     string `gorm:"not null"`
	CreatedAt int
	UpdatedAt int
	DeletedAt int
}

func (u *User) TableName() string {
	return table.TbaUser
}
