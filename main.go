package main

import (
	"web-golang/utils"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/go-playground/validator/v10"
	"web-golang/app"
	"web-golang/controller"
	"web-golang/service"
	"web-golang/repository"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	mahasiswaRepository := repository.NewMahasiswaRepository()
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepository, db, validate)
	mahasiswaController := controller.NewMahasiswaController(mahasiswaService)
	router := app.NewRouter(mahasiswaController)

	server := http.Server{
        Addr: "localhost:3000",
        Handler : router,
    }

    err := server.ListenAndServe()
    utils.PanicIfError(err)
}