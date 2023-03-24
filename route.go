package testproject

import (
	db "testproject/_config"
	"testproject/controller"
	"testproject/middleware"

	"github.com/gin-gonic/gin"
)

// MemberRoute create route
func Route(r *gin.Engine, resource *db.Resource) {
	r.POST("/login", controller.Login(resource))
	r.GET("/test/:username", middleware.AuthorizationMiddleware(), controller.TestPayToWalletV2Ufabet(resource))
	r.POST("/test_create", middleware.AuthorizationMiddleware(), controller.CreateUserAccount(resource))
	r.PATCH("/test_update", middleware.AuthorizationMiddleware(), controller.UpdateUserAccount(resource))
	r.DELETE("/test_delete", middleware.AuthorizationMiddleware(), controller.DeleteUserAccount(resource))
}
