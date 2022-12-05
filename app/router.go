package app

import (
	"github.com/julienschmidt/httprouter"
	"web-golang/controller"
	"web-golang/exception"
)

func NewRouter(mahasiswaController controller.MahasiswaController) *httprouter.Router{
	router := httprouter.New()

	router.GET("/api/mahasiswa", mahasiswaController.FindAll)
	router.GET("/api/mahasiswa/:mahasiswaId", mahasiswaController.FindById)

	router.PanicHandler = exception.ErrorHandler

	return router
}