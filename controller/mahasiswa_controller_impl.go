package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
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

func (controller *MahasiswaControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mahasiswaId := params.ByName("mahasiswaId")
	id, err := strconv.Atoi(mahasiswaId)
	utils.PanicIfError(err)

	mahasiswaResponse := controller.MahasiswaService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: mahasiswaResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *MahasiswaControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	mahasiswaCreateRequest := web.MahasiswaCreateRequest{}
	utils.ReadFromRequestBody(request, &mahasiswaCreateRequest)

	mahasiswaResponse := controller.MahasiswaService.Create(request.Context(), mahasiswaCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: mahasiswaResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}