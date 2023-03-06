package model

import (
	"gorm.io/gorm"
	"musicMod/utils"
	"strconv"
)

type User struct {
	gorm.Model
	Username   string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password   string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email      string `gorm:"type:varchar(255);not null" json:"email"`
	Likemusics string `gorm:"type:varchar(255);" json:"likes"`
	Liketypes  string `gorm:"type:varchar(255);" json:"liketypes"`
	Photo      string `gorm:"type:varchar(255);" json:"photo"`
	Motto      string `gorm:"type:varchar(255);" json:"motto"`
}

func (this *User) TableName() string {
	return "User"
}
func AddUser(username string, password string, email string) int {
	newUser := User{
		Username: username,
		Password: password,
		Email:    email,
	}
	err = db.Create(&newUser).Error
	if err != nil {
		return utils.ERROR_CREAT_WRONG
	}
	return utils.SUCCESS
}
func CheckAdd(username string, email string) (code int) {
	var user User
	db.Debug().Select("ID").Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return utils.ERROR_TEL_USED
	}
	db.Debug().Select("ID").Where("username = ?", username).First(&user)
	if user.ID != 0 {
		return utils.ERROR_USERNAME_USED
	}
	return utils.SUCCESS
}
func CheckLogin(username string, password string) (code int, ID string) {
	var user User
	db.Select("ID").Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return utils.ERROR_USERNAME_NOT_EXIST, ""
	}
	db.Select("password").Where("username = ?", username).First(&user)
	if user.Password == password {
		return utils.SUCCESS, strconv.Itoa(int(user.ID))
	} else {
		return utils.ERROR_PASSWORD_WRONG, ""
	}
}
func CheckLoginEmail(email string) int {
	var user User
	db.Select("ID").Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return utils.ERROR_EMAIL_NOT_EXIST
	}
	return utils.SUCCESS
}
func FindLoginId(email string) string {
	var user User
	db.Select("ID").Where("email = ?", email).First(&user)
	return strconv.Itoa(int(user.ID))
}
func AddMusicKind(id string, musicKind string) int {
	err = db.Model(&User{}).Where("id = ?", id).Update("liketypes", musicKind).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}
func FindUser(id string) (int, User) {
	var user User
	err = db.Debug().Where("ID = ?", id).First(&user).Error
	if err != nil {
		return utils.ERROR, user
	}
	return utils.SUCCESS, user
}
func UpdateUser(id string, user User) int {
	err = db.Where("id = ?", id).Updates(user).Error
	if err != nil {
		return utils.ERROR_CHANGE_WRONG
	}
	return utils.SUCCESS
}
