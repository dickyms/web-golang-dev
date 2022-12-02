package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"web-golang/utils"
	"web-golang/service"
	"web-golang/model/web"
)

type MahasiswaControllerImpl struct {
	MahasiswaService service.MahasiswaService
}

func NewMahasiswaController( mahasiswaService service.MahasiswaService) MahasiswaController {
	return &MahasiswaControllerImpl{
		MahasiswaService : mahasiswaService,
	}
}

func (controller *MahasiswaControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mahasiswaResponeses := controller.MahasiswaService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
		Data: mahasiswaResponeses,
	}

	utils.WriteToResponseBody(writer, webResponse)
}