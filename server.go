package testproject

import (
	"net/http"
	"time"

	db "testproject/_config"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	_ = godotenv.Load()
	color.Green("Server starting...")

	r := gin.Default()
	r.Use(CORS)

	// m := ginmetrics.GetMonitor()
	// m.SetMetricPath("/metrics")
	// m.SetSlowTime(2)
	// m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	// m.Use(r)

	resource, err := db.CreateResource()
	if err != nil {
		color.Red("Connection database failure, Please check connection")
		color.Cyan(err.Error())
		logrus.Error(err)
		time.Sleep(5 * time.Second)
		return
	}
	defer resource.Close()

	// Route(r)
	Route(r, resource)

	r.Run()
}

func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
		return
	} else {
		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		// c.AbortWithStatusJSON(401, gin.H{
		// 	"code":    "ERROR",
		// 	"message": "Unauthorized",
		// })
		c.AbortWithStatus(http.StatusOK)
		return
	}
}
