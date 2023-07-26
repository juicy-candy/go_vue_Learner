package midware

import (
	"ginvue/pkg/database"
	"ginvue/pkg/model"
	"ginvue/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization
		tokenString := ctx.GetHeader("Authorization")

		//格式验证
		// 如果token为空或者不以Bearer开头,则token未传或错误
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的token",
			})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无权限",
			})
			ctx.Abort()
			return
		}

		userId := claims.UserId
		db := database.GetDB()

		var user model.User
		db.Where("id = ?", userId).First(&user)
		if user.ID == 0 {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "错误的token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
