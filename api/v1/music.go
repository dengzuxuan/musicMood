package v1

import (
	"github.com/gin-gonic/gin"
	"musicMod/model"
	"musicMod/utils"
	"net/http"
)

func GetMusic(c *gin.Context) {
	musictype := c.Query("type")
	mode := c.Query("mod")
	musics, code := model.GetMusic(musictype, mode)
	//fmt.Println(musics[0].Url)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
		"data":    musics,
	})
}
func AddMusic(c *gin.Context) {
	musicName := c.Query("name")
	musicAuthor := c.Query("author")
	musictype := c.Query("type")
	mood := c.Query("mod")
	url := "http://8.140.38.47/music/" + musicAuthor + " - " + musicName + ".mp3"
	code := model.AddMusic(musicName, musicAuthor, musictype, mood, url)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func AddLikes(c *gin.Context) {
	id := c.Query("id")
	userId := c.Query("userId")
	//音乐点赞+1
	code := model.AddLikes(id)
	if code != utils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": utils.GetErrMsg(code),
		})
		return
	}
	//用户->点赞列表 +1
	code = model.AddUserLikes(userId, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func DesLikes(c *gin.Context) {
	id := c.Query("id")
	//音乐点赞-1
	code := model.DesLikes(id)
	userId := c.Query("userId")
	if code != utils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": utils.GetErrMsg(code),
		})
		return
	}
	//用户->音乐列表-1
	code = model.DesUserLikes(userId, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func AddCols(c *gin.Context) {
	id := c.Query("id")
	code := model.AddCols(id)
	userId := c.Query("userId")
	if code != utils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": utils.GetErrMsg(code),
		})
		return
	}
	code = model.AddUserCols(userId, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func DesCols(c *gin.Context) {
	id := c.Query("id")
	code := model.DesCols(id)
	userId := c.Query("userId")
	if code != utils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": utils.GetErrMsg(code),
		})
		return
	}
	code = model.DesUserCols(userId, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func IsLikes(c *gin.Context) {
	id := c.Query("id")
	userId := c.Query("userId")
	data := model.IsLikes(userId, id)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func IsCols(c *gin.Context) {
	id := c.Query("id")
	userId := c.Query("userId")
	data := model.IsCols(userId, id)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
