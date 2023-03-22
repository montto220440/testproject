package testproject

import (
	db "testproject/_config"
	"testproject/controller"

	"github.com/gin-gonic/gin"
)

// MemberRoute create route
func Route(r *gin.Engine, resource *db.Resource) {
	r.GET("/test/:username", controller.TestPayToWalletV2Ufabet(resource))
	r.GET("/test_create", controller.CreateUserAccount(resource))
	r.GET("/test_update", controller.UpdateUserAccount(resource))
	r.GET("/test_delete", controller.DeleteUserAccount(resource))
}
