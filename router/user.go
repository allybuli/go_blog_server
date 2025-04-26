package router

import (
	"server/api"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, AdminRouter *gin.RouterGroup) {
	userLoginRouter := PublicRouter.Group("user").Use(middleware.LoginRecord())
	userApi := api.ApiGroupApp.UserApi
	{
		userLoginRouter.POST("register", userApi.Register)
		userLoginRouter.POST("login", userApi.Login)
	}
}
