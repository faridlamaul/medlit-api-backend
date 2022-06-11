package models

import (
	"errors"

	"github.com/faridlamaul/medlit-api-backend/api/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID		  uint	 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	Password  string `gorm:"size:255;not null" json:"password"`
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func GetNameByEmail(email string) string {
	var user User
	DB.Where("email = ?", email).Take(&user)
	return user.Name
}

func GetUserByID(uid uint) (User, error) {
	var user User

	if err := DB.First(&user, uid).Error; err != nil {
		return user, errors.New("User not found")
	}

	user.PrepareGive()

	return user, nil
}

func (user *User) PrepareGive() {
	user.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
} 

func (user *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}