package utils

import (
	"errors"
	"ginvue/pkg/model"

	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func IsEmail(mail string) bool {
	if len(mail) < 3 {
		return false
	}
	result, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, mail)
	return result
}

func IsEmailExist(db *gorm.DB, mail string) bool {
	var user model.User
	db.Where("mail = ?", mail).First(&user)
	return user.ID != 0
}

func GetUser(db *gorm.DB, name string, passwd string) (model.User, error) {
	var user model.User
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		err := errors.New("用户不存在")
		return user, err
	}
	if err := IsPasswdTrue(passwd, user); err != nil {
		return user, err
	}

	return user, nil
}

func IsPasswdTrue(passwd string, user model.User) error {
	err := errors.New("错误的密码")
	if p_err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd)); p_err != nil {
		return err
	}

	return nil
}

func PasswdEncode(raw string) (string, error) {
	rtn, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	return string(rtn), err
}
