package model

import (
	"gorm.io/gorm"
	"musicMod/utils"
	"strconv"
	"strings"
)

type Music struct {
	gorm.Model
	MusicName   string `gorm:"type:varchar(255);not null" json:"music_name" `
	MusicAuthor string `gorm:"type:varchar(255);not null" json:"music_author"`
	MusicType   string `gorm:"type:varchar(20);" json:"music_type"`
	MusicMod    string `gorm:"type:varchar(20);" json:"music_mod"`
	Url         string `gorm:"type:varchar(255);" json:"url"`
	Likes       int    `gorm:"type:int; default:0;" json:"likes"`
	Cols        int    `gorm:"type:int;default:0;" json:"col"`
}

func (this *Music) TableName() string {
	return "Music"
}

func GetMusic(musicType string, mode string) ([]Music, int) {
	musicTypes := strings.Split(musicType, " ")
	var musics []Music
	err = db.Where(map[string]interface{}{"music_type": musicTypes, "music_mod": mode}).Find(&musics).Error
	if err != nil {
		return nil, utils.ERROR
	}
	return musics, utils.SUCCESS
}
func AddMusic(musicName string, musicAuthor string, musictype string, mood string, url string) int {
	music := Music{
		MusicName:   musicName,
		MusicAuthor: musicAuthor,
		MusicType:   musictype,
		MusicMod:    mood,
		Url:         url,
	}
	err = db.Create(&music).Error
	if err != nil {
		return utils.ERROR
	}
	return utils.SUCCESS
}
func AddLikes(id string) int {
	music := Music{}
	v, _ := strconv.Atoi(id)
	music.ID = uint(v)
	err = db.Model(&music).UpdateColumn("Likes", gorm.Expr("Likes + 1")).Error
	if err != nil {
		return utils.ERROR_ADDLIKES_WRONG
	}
	return utils.SUCCESS
}
func DesLikes(id string) int {
	music := Music{}
	v, _ := strconv.Atoi(id)
	music.ID = uint(v)
	err = db.Model(&music).UpdateColumn("Likes", gorm.Expr("Likes - 1")).Error
	if err != nil {
		return utils.ERROR_DESLIKES_WRONG
	}
	return utils.SUCCESS
}
func AddCols(id string) int {
	music := Music{}
	v, _ := strconv.Atoi(id)
	music.ID = uint(v)
	err = db.Model(&music).UpdateColumn("Cols", gorm.Expr("Cols + 1")).Error
	if err != nil {
		return utils.ERROR_ADDCOLS_WRONG
	}
	return utils.SUCCESS
}
func DesCols(id string) int {
	music := Music{}
	v, _ := strconv.Atoi(id)
	music.ID = uint(v)
	err = db.Model(&music).UpdateColumn("Cols", gorm.Expr("Cols - 1")).Error
	if err != nil {
		return utils.ERROR_DESCOLS_WRONG
	}
	return utils.SUCCESS
}
func MusicInfo(id string) Music {
	var music Music
	db.Where("id = ?", id).Find(&music)
	return music
}
