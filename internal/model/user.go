package model

import "fiber-wire-template/pkg/util/table"

type User struct {
	Id        uint   `gorm:"primarykey" json:"id,omitempty"`
	UserId    string `gorm:"unique;not null" json:"user_id,omitempty"`
	Username  string `gorm:"unique;not null" json:"username,omitempty"`
	Nickname  string `gorm:"not null" json:"nickname,omitempty"`
	Password  string `gorm:"not null" json:"password,omitempty"`
	Skill     string `gorm:"not null" json:"skill,omitempty"`
	Sex       int    `gorm:"type:tinyint(1);not null" json:"sex,omitempty"`
	Status    int    `gorm:"type:tinyint(1);not null" json:"status,omitempty"`
	Address   string `gorm:"not null" json:"address,omitempty"`
	Mobile    string `gorm:"not null" json:"mobile,omitempty"`
	QQ        string `gorm:"not null" json:"qq,omitempty"`
	Email     string `gorm:"not null" json:"email,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	UpdatedAt int    `json:"updated_at,omitempty"`
	DeletedAt int    `json:"deleted_at,omitempty"`
}

func (u *User) TableName() string {
	return table.TbaUser
}
