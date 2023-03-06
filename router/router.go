package router

import (
	"github.com/gin-gonic/gin"
	v1 "musicMod/api/v1"
	"musicMod/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	user := r.Group("api/v1/user")
	{
		user.GET("/getcode", v1.AddUserPre)
		user.GET("/add", v1.AddUser)

		user.GET("/login", v1.LoginUser)
		user.GET("/loginemailpre", v1.LoginEmailPre)
		user.GET("/loginemail", v1.LoginEmail)

		user.GET("/getuser", v1.GetUser)
		user.GET("/updatepre", v1.UpdateUserPre)
		user.GET("/updateemailpre", v1.UpdateEmailPre)
		user.GET("/update", v1.UpdateUser)
		user.GET("/musiclikes", v1.MusicLike)

		user.GET("/getlikeslist", v1.GetLikesList)
		user.GET("/getcolslist", v1.GetColsLike)
	}
	music := r.Group("api/v1/music")
	{
		music.GET("/find", v1.GetMusic)
		music.GET("/add", v1.AddMusic)

		music.GET("/addlikes", v1.AddLikes)
		music.GET("/deslikes", v1.DesLikes)

		music.GET("/addcols", v1.AddCols)
		music.GET("/descols", v1.DesCols)

		music.GET("/islikes", v1.IsLikes)
		music.GET("/iscols", v1.IsCols)
	}
	return r
}
