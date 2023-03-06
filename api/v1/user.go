package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"musicMod/middleware"
	"musicMod/model"
	"musicMod/utils"
	"net/http"
)

var emailcodesave = make(map[string]string)        //username -> emailcode
var emailcodeloginsave = make(map[string]string)   //email -> emailcode
var emailcodechangesave = make(map[string]string)  //id -> emailcode
var emailcodechangeemail = make(map[string]string) //id -> emailcode
func LoginUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if len(username) < 4 || len(username) > 12 {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_USERNAME_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_USERNAME_WRONG),
		})
		return
	}
	if len(password) < 6 || len(password) > 20 {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_PASSWORD_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_PASSWORD_WRONG),
		})
		return
	}
	code, id := model.CheckLogin(username, password)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
		"id":      id,
	})
}
func LoginEmailPre(c *gin.Context) {
	email := c.Query("email")
	code := model.CheckLoginEmail(email)
	if code != utils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": utils.GetErrMsg(code),
		})
		return
	}
	emailcode, status := middleware.CheckMail(email)
	if status == utils.SUCCESS {
		emailcodeloginsave[email] = emailcode
	}
	//code = model.AddUser(username,password,email,musictype)
	c.JSON(http.StatusOK, gin.H{
		"status":    status,
		"message":   utils.GetErrMsg(status),
		"emailcode": emailcode,
	})
}
func LoginEmail(c *gin.Context) {
	email := c.Query("email")
	emailcode := c.Query("emailcode")
	if emailcodeloginsave[email] != emailcode {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.ERROR_EMAIL_CHECK,
			"message": utils.GetErrMsg(utils.ERROR_EMAIL_CHECK),
		})
		return
	}
	id := model.FindLoginId(email)
	c.JSON(http.StatusOK, gin.H{
		"code":    utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
		"id":      id,
	})
}
func AddUserPre(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")
	//musictype:=c.Query("musictype")
	//fmt.Println(username,password,tel,musictype)
	if len(username) < 4 || len(username) > 12 {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_USERNAME_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_USERNAME_WRONG),
		})
		return
	}
	if len(password) < 6 || len(password) > 20 {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_PASSWORD_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_PASSWORD_WRONG),
		})
		return
	}
	code := model.CheckAdd(username, email)
	if code != utils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": utils.GetErrMsg(code),
		})
		return
	} else {
		emailcode, status := middleware.CheckMail(email)
		if status == utils.SUCCESS {
			emailcodesave[username] = emailcode
		}
		//code = model.AddUser(username,password,email,musictype)
		c.JSON(http.StatusOK, gin.H{
			"status":    status,
			"message":   utils.GetErrMsg(status),
			"emailcode": emailcode,
		})
	}
}

func AddUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")
	emailcode := c.Query("emailcode")
	v := emailcodesave[username]
	if emailcode != v {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.ERROR_EMAIL_CHECK,
			"message": utils.GetErrMsg(utils.ERROR_EMAIL_CHECK),
		})
		return
	}
	code := model.AddUser(username, password, email)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": utils.GetErrMsg(code),
	})

}

func GetUser(c *gin.Context) {
	id := c.Query("id")
	code, user := model.FindUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
		"data":    user,
	})
}

func UpdateUserPre(c *gin.Context) {
	//需要更改什么就传入什么
	id := c.Query("id")
	_, user := model.FindUser(id)
	emailcode, status := middleware.CheckMail(user.Email)
	if status == utils.SUCCESS {
		emailcodechangesave[id] = emailcode
	}
	//code = model.AddUser(username,password,email,musictype)
	c.JSON(http.StatusOK, gin.H{
		"status":    status,
		"message":   utils.GetErrMsg(status),
		"emailcode": emailcode,
	})
}
func UpdateEmailPre(c *gin.Context) {
	email := c.Query("email")
	id := c.Query("id")
	emailcode, status := middleware.CheckMail(email)
	if status == utils.SUCCESS {
		emailcodechangeemail[id] = emailcode
	}
	//code = model.AddUser(username,password,email,musictype)
	c.JSON(http.StatusOK, gin.H{
		"status":    status,
		"message":   utils.GetErrMsg(status),
		"emailcode": emailcode,
	})
}
func UpdateUser(c *gin.Context) {
	//需要更改什么就传入什么
	id := c.Query("id")
	newpwd := c.Query("password")
	newname := c.Query("name")
	newemail := c.Query("email")
	newmotto := c.Query("motto")
	emailcode := c.Query("emailcode")
	emailnewcode := c.Query("emailnewcode")
	if (len(newname) < 4 || len(newname) > 12) && len(newname) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_USERNAME_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_USERNAME_WRONG),
		})
		return
	}
	if (len(newpwd) < 6 || len(newpwd) > 20) && len(newpwd) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  utils.ERROR_PASSWORD_WRONG,
			"message": utils.GetErrMsg(utils.ERROR_PASSWORD_WRONG),
		})
		return
	}
	code := model.CheckAdd(newname, newemail)
	if code != utils.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": utils.GetErrMsg(code),
		})
		return
	}
	//更改密码 绑定邮件需要发送邮件认证
	if newpwd != "" || newemail != "" {
		fmt.Println(emailcode)
		fmt.Println(emailcodechangesave[id])
		if emailcode != emailcodechangesave[id] || emailcode == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    utils.ERROR_EMAIL_CHECK,
				"message": utils.GetErrMsg(utils.ERROR_EMAIL_CHECK),
			})
			return
		}
		if newemail != "" {
			if emailnewcode != emailcodechangeemail[id] || emailnewcode == "" {
				c.JSON(http.StatusOK, gin.H{
					"code":    utils.ERROR_EMAIL_CHECK,
					"message": utils.GetErrMsg(utils.ERROR_EMAIL_CHECK),
				})
				return
			}
		}
	}
	newuser := model.User{
		Username: newname,
		Password: newpwd,
		Email:    newemail,
		Motto:    newmotto,
	}
	code = model.UpdateUser(id, newuser)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
	//更改姓名以及个性签名则不需要
}
func MusicLike(c *gin.Context) {
	id := c.Query("id")
	musicLikes := c.Query("musiclikes")
	code := model.AddMusicKind(id, musicLikes)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

//todo
func GetLikesList(c *gin.Context) {
	id := c.Query("id")
	likesList, code := model.GetLikesList(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
		"data":    likesList,
	})
}

func GetColsLike(c *gin.Context) {
	id := c.Query("id")
	colsList, code := model.GetColsList(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": utils.GetErrMsg(code),
		"data":    colsList,
	})
}
