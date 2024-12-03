package controller

import (
	"net/http"
	"simple-crud-api/data/request"
	"simple-crud-api/data/response"
	"simple-crud-api/helper"
	"simple-crud-api/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type SongController struct {
	SongService service.SongService
}

func NewSongController(songService service.SongService) *SongController {
	return &SongController{SongService: songService}
}

func (controller *SongController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	songCreateRequest := request.SongCreateRequest{}
	helper.ReadRequestBody(requests, &songCreateRequest)

	controller.SongService.Create(requests.Context(), songCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SongController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	songUpdateRequest := request.SongUpdateRequest{}
	helper.ReadRequestBody(requests, &songUpdateRequest)

	songId := params.ByName("songId")
	id, err := strconv.Atoi(songId)
	helper.PanicIfError(err)
	songUpdateRequest.Id = id

	controller.SongService.Update(requests.Context(), songUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SongController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	songId := params.ByName("songId")
	id, err := strconv.Atoi(songId)
	helper.PanicIfError(err)

	controller.SongService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SongController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.SongService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *SongController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	songId := params.ByName("songId")
	id, err := strconv.Atoi(songId)
	helper.PanicIfError(err)

	result := controller.SongService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: result,
	}
	helper.WriteResponseBody(writer, webResponse)
}
