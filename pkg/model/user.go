package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string ` gorm:"type:varchar(20); not null,unique" `
	Mail   string ` gorm:"type:varchar(255); unique; not null"`
	Passwd string `gorm:"type: varchar(255); not null"`
}

type UserDto struct {
	Name string `json:"name"`
	Mail string `json:"Mail"`
	Id   uint   `json:"id"`
}

func ToUserDto(user User) UserDto {
	return UserDto{
		Name: user.Name,
		Mail: user.Mail,
		Id:   user.Model.ID,
	}
}
