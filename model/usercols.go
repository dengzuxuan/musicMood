package model

import (
	"gorm.io/gorm"
	"musicMod/utils"
)

type Usercols struct {
	gorm.Model
	MusicId string `gorm:"type:varchar(255);" json:"music_id"`
	UserId  string `gorm:"type:varchar(255);" json:"user_id"`
	Title   string `gorm:"type:varchar(255);" json:"title"`
	Author  string `gorm:"type:varchar(255);" json:"author"`
}

func (this *Usercols) TableName() string {
	return "Usercols"
}

//func AddUserCols(userId string,musicId string) int {
//	newUserCol:=Usercols{UserId: userId,MusicId: musicId}
//	err=db.Create(&newUserCol).Error
//	if err!=nil{
//		return utils.ERROR_USERADDCOLS_WRONG
//	}
//	return utils.SUCCESS
//}
//func DesUserCols(userId string,musicId string) int {
//	delUserCol:=Usercols{UserId: userId,MusicId: musicId}
//	err=db.Debug().Delete(&delUserCol).Error
//	if err!=nil{
//		return utils.ERROR_USERADDCOLS_WRONG
//	}
//	return utils.SUCCESS
//}
func AddUserCols(userId string, musicId string) int {
	musicInfo := MusicInfo(musicId)
	newUserCol := Usercols{UserId: userId, MusicId: musicId, Title: musicInfo.MusicName, Author: musicInfo.MusicAuthor}
	err = db.Debug().Create(&newUserCol).Error
	if err != nil {
		return utils.ERROR_USERADDCOLS_WRONG
	}
	return utils.SUCCESS
}
func DesUserCols(userId string, musicId string) int {
	var delUserCol Usercols
	err = db.Debug().Where("music_id = ? AND user_id =?", musicId, userId).Delete(&delUserCol).Error
	if err != nil {
		return utils.ERROR_USERDESCOLS_WRONG
	}
	return utils.SUCCESS
}
func GetColsList(id string) ([]Usercols, int) {
	var colsList []Usercols
	err = db.Where("user_id = ?", id).Find(&colsList).Error
	if err != nil {
		return nil, utils.ERROR_USERLIKESLIST_WRONG
	}
	return colsList, utils.SUCCESS
}
func IsCols(id string, musicId string) bool {
	var usercols Usercols
	db.Where("music_id = ? AND user_id =?", musicId, id).First(&usercols)
	if usercols.ID != 0 {
		return true
	}
	return false
}
