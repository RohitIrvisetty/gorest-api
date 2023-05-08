package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	about "github.com/RohitIrvisetty/gorestapi/api/aboutapi"
)

var aboutrestapi = []about.Aboutapi{
	{App: "gorestapi", Author: "RohitIrvisetty", Github: "https://github.com/RohitIrvisetty/gorestapi", Version: 2},
}

/* perform post request */
func maintest() {
	router := gin.Default()

	testapi := router.Group("/testapi")
	{
		router.GET("/test,", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
			})
		})

		router.GET("/aboutest", getAboutApi)

		/* testapi router only for post functions */
		testapi.POST("/aboutest", postAboutApi)

		router.Run()
	}
}

func getAboutApi(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, aboutrestapi)
}

func postAboutApi(c *gin.Context) {
	var newAboutApi about.Aboutapi

	if err := c.BindJSON(&newAboutApi); err != nil {
		return
	}

	aboutrestapi = append(aboutrestapi, newAboutApi)
	c.IndentedJSON(http.StatusCreated, newAboutApi)
}
