package req

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

)
func Handler(c *gin.Context) {
	x, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("%s", string(x))
	c.JSON(http.StatusOK, c)
}
func Req() {
	router := gin.Default()

	router.GET("/", Handler)

	router.POST("/new-message", Handler)

	router.Run()
}