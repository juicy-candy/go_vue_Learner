package controller

import (
	"ginvue/pkg/database"
	"ginvue/pkg/model"
	"ginvue/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	db := database.GetDB()

	name := ctx.PostForm("name")
	mail := ctx.PostForm("mail")
	passwd := ctx.PostForm("passwd")

	if !utils.IsEmail(mail) || utils.IsEmailExist(db, mail) {
		utils.Failed(ctx, nil, "邮箱错误！")
		return
	}

	if len(name) < 5 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名长度过短!",
			"len":  len(name),
		})
		return
	}

	if len(passwd) < 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码过短!",
		})
		return
	}

	key, err := utils.PasswdEncode(passwd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "服务器错误!",
		})
		return
	}

	newUser := model.User{
		Name:   name,
		Mail:   mail,
		Passwd: key,
	}

	db.Create(&newUser)
	ctx.JSON(200, gin.H{"code": 200, "msg": "success"})
	log.Println(name, mail, passwd)
}

func Login(ctx *gin.Context) {
	db := database.GetDB()
	name := ctx.PostForm("name")
	passwd := ctx.PostForm("passwd")

	user, err := utils.GetUser(db, name, passwd)

	if err != nil {
		ctx.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	//发放token

	token, token_err := utils.GetToken(user)

	if token_err != nil {
		ctx.JSON(400, gin.H{
			"code": 500,
			"msg":  "系统错误,发放token失败",
		})
		log.Print("token发放失败")
		return
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功！",
		"data": gin.H{
			"token": token,
		},
	})

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"user": model.ToUserDto(user.(model.User))}})
}
