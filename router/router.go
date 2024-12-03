package router

import (
	"fmt"
	"net/http"
	"simple-crud-api/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(songController *controller.SongController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome")
	})

	router.GET("/api/song", songController.FindAll)
	router.GET("/api/song/:songId", songController.FindById)
	router.POST("/api/song", songController.Create)
	router.PATCH("/api/song/:songId", songController.Update)
	router.DELETE("/api/song/:songId", songController.Delete)

	return router
}
