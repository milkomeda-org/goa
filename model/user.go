package model

import (
	"golang.org/x/crypto/bcrypt"
	"oa-auth/initializer/db"
)

// User 用户模型
type User struct {
	BaseModel
	UserName       string `gorm:"not null;comment:'用户名'"`
	PasswordDigest string `gorm:"not null;comment:'密码摘要'"`
	Nickname       string `gorm:"not null;comment:'昵称'"`
	Avatar         string `gorm:"size:1000;comment:'头像'"`
	PositionID     int    `gorm:"not null;comment:'身份ID'"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := db.DB.Where("id = ?", ID).First(&user)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}