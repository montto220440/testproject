package controller

import (
	"fmt"
	db "testproject/_config"
	"testproject/model"
	"testproject/repo"
	"testproject/service"

	"testproject/dto"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TestPayToWalletV2Ufabet(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {

		username := c.Param("username")
		fmt.Println("username", username)
		data := model.PersonalData{}
		filter := bson.M{"username": username}
		filterOp := bson.M{
			"username": 1,
		}
		repo.GetOneStatement(resource, "member_account", filter, filterOp, &data)
		// c.JSON(200, data)
		service.Resp(c, 200, "Success", data)
	}
}

func CreateUserAccount(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {

		form := dto.ListRequest{}
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		repo.CreateUser(resource, "member_account", form)
		// c.JSON(200, gin.H{"Success to create member": form})
		service.Resp(c, 200, "Success to create member", form)
	}
}

func UpdateUserAccount(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {

		form := dto.ListRequest{}
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		filter := bson.M{"id": form.Id}
		filterOp := bson.M{
			"$set": bson.M{
				"username": form.Username,
				"password": form.Password,
				"phone":    form.Phone,
				"age":      form.Age,
				"gender":   form.Gender,
			}}
		repo.UpdateUser(resource, "member_account", filter, filterOp)
		// c.JSON(200, gin.H{"Success to update member": form})
		service.Resp(c, 200, "Success to update member", form)
	}
}

func DeleteUserAccount(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {

		form := dto.ListRequest{}
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		filter := bson.M{"id": form.Id}
		repo.DeleteUser(resource, "member_account", filter)
		// c.JSON(200, gin.H{"Success to delete member": form.Id})
		service.Resp(c, 200, "Success to delete member", form.Id)
	}
}
