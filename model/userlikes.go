package model

import (
	"fmt"
	"gorm.io/gorm"
	"musicMod/utils"
)

type Userlikes struct {
	gorm.Model
	MusicId string `gorm:"type:varchar(255);" json:"music_id"`
	UserId  string `gorm:"type:varchar(255);" json:"user_id"`
	Title   string `gorm:"type:varchar(255);" json:"title"`
	Author  string `gorm:"type:varchar(255);" json:"author"`
}

func (this *Userlikes) TableName() string {
	return "Userlikes"
}
func AddUserLikes(userId string, musicId string) int {
	musicInfo := MusicInfo(musicId)
	newUserLike := Userlikes{UserId: userId, MusicId: musicId, Title: musicInfo.MusicName, Author: musicInfo.MusicAuthor}
	err = db.Debug().Create(&newUserLike).Error
	if err != nil {
		return utils.ERROR_USERADDCOLS_WRONG
	}
	return utils.SUCCESS
}
func DesUserLikes(userId string, musicId string) int {
	var delUserLike Userlikes
	fmt.Println(delUserLike)
	err = db.Debug().Where("music_id = ? AND user_id =?", musicId, userId).Delete(&delUserLike).Error
	if err != nil {
		return utils.ERROR_USERDESLIKES_WRONG
	}
	return utils.SUCCESS
}
func GetLikesList(id string) ([]Userlikes, int) {
	var likesList []Userlikes
	err = db.Where("user_id = ?", id).Find(&likesList).Error
	if err != nil {
		return nil, utils.ERROR_USERLIKESLIST_WRONG
	}
	return likesList, utils.SUCCESS
}

func IsLikes(id string, musicId string) bool {
	var userlikes Userlikes
	db.Where("music_id = ? AND user_id =?", musicId, id).First(&userlikes)
	if userlikes.ID != 0 {
		return true
	}
	return false
}
