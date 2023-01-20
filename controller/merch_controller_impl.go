package controller

import (
	"e-commerce-api/model/web"
	"e-commerce-api/service"
	"e-commerce-api/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MerchControllerImpl struct {
	MerchService service.MerchService
}

func NewMerchController(merchService service.MerchService) MerchController {
	return &MerchControllerImpl{
		MerchService: merchService,
	}
}

func (controller *MerchControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	merchResponses := controller.MerchService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code: 204,
		Status: "No Content",
		Data: merchResponses,
	}

	utils.WriteToResponseBody(w, webResponse)
}

func (controller *MerchControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	merchId := params.ByName("merchId")
	id, err := strconv.Atoi(merchId)
	utils.PanicIfError(err)

	merchResponse := controller.MerchService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code: 204,
		Status: "No Content",
		Data: merchResponse,
	}

	utils.WriteToResponseBody(w, webResponse)
}

func (controller *MerchControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	merchCreateRequest := web.MerchCreateRequest{}
	utils.ReadFromRequestBody(r, &merchCreateRequest)

	merchResponse := controller.MerchService.Create(r.Context(), merchCreateRequest)
	webResponse := web.WebResponse{
		Code: 201,
		Status: "Created",
		Data: merchResponse,
	}

	utils.WriteToResponseBody(w, webResponse)
}

func (controller *MerchControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	merchUpdateRequest := web.MerchUpdateRequest{}
	utils.ReadFromRequestBody(r, &merchUpdateRequest)

	merchId := params.ByName("merchId")
	id, err := strconv.Atoi(merchId)
	utils.PanicIfError(err)

	merchUpdateRequest.Id = id
	
	merchResponse := controller.MerchService.Update(r.Context(), merchUpdateRequest)
	webResponse := web.WebResponse{
		Code: 201,
		Status: "Created",
		Data: merchResponse,
	}

	utils.WriteToResponseBody(w, webResponse)
}

func (controller *MerchControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	merchId := params.ByName("merchId")
	id, err := strconv.Atoi(merchId)
	utils.PanicIfError(err)

	controller.MerchService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code: 204,
		Status: "No Content",
	}

	utils.WriteToResponseBody(w, webResponse)
}

// func (controller *MerchControllerImpl) Order(w http.ResponseWriter, r *http.Request) {
// }