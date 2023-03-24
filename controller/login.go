package controller

import (
	"net/http"
	db "testproject/_config"
	"testproject/dto"
	"testproject/model"
	"testproject/repo"
	"testproject/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		//check login
		login_data := dto.LoginRequest{}
		if err := c.ShouldBindJSON(&login_data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		res_data := model.PersonalData{}
		filter := bson.M{"username": login_data.Username, "password": login_data.Password}
		filterOp := bson.M{
			"username": 1,
		}
		err := repo.GetOneStatement(resource, "member_account", filter, filterOp, &res_data)

		if err != nil { //return ivalid
			service.Resp(c, 400, "Invalid username or password", nil)
		} else { //generate token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
				ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
			})
			ss, err := token.SignedString([]byte("MySignature"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"token": ss,
			})
		}
	}
}
