package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(32);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:varchar(128);not null" json:"-"`
	Role     string `gorm:"type:varchar(16);not null;default:'user'" json:"role"`
}

// HashPassword 将明文密码加密
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword 检查密码是否正确
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// HashPassword 为用户设置加密后的密码
func (u *User) HashPassword() error {
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
