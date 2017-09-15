package req

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Handler ( w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}


func req() {
	router := httprouter.New()

	router.GET("/", Handler)

	router.POST("/new-message", Handler)
}