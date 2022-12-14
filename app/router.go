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
	router.POST("/api/mahasiswa", mahasiswaController.Create)
	router.DELETE("/api/mahasiswa/:mahasiswaId", mahasiswaController.Delete)
	router.PUT("/api/mahasiswa/:mahasiswaId", mahasiswaController.Update)

	router.PanicHandler = exception.ErrorHandler

	return router
}