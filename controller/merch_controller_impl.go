package controller

import (
	"e-commerce-api/model/web"
	"e-commerce-api/service"
	"e-commerce-api/utils"
	"net/http"
)

type MerchControllerImpl struct {
	MerchService service.MerchService
}

func NewMerchController(merchService service.MerchService) MerchController {
	return &MerchControllerImpl{
		MerchService: merchService,
	}
}

func (controller *MerchControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	merchResponses := controller.MerchService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code: 204,
		Status: "No Content",
		Data: merchResponses,
	}

	utils.WriteToResponseBody(w, webResponse)
}

func (controller *MerchControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	//merchUpdateRequest.Id = params.ByName("merchId")
	//id, err := strconv.Atoi(merchId)
	//utils.PanicIfError(err)

	id := 1 // placeholder
	merchResponse := controller.MerchService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code: 204,
		Status: "No Content",
		Data: merchResponse,
	}

	utils.WriteToResponseBody(w, webResponse)
}

func (controller *MerchControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
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

func (controller *MerchControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	merchUpdateRequest := web.MerchUpdateRequest{}
	utils.ReadFromRequestBody(r, &merchUpdateRequest)

	//merchUpdateRequest.Id = params.ByName("merchId")
	//id, err := strconv.Atoi(merchId)
	//utils.PanicIfError(err)

	merchResponse := controller.MerchService.Update(r.Context(), merchUpdateRequest)
	webResponse := web.WebResponse{
		Code: 201,
		Status: "Created",
		Data: merchResponse,
	}

	utils.WriteToResponseBody(w, webResponse)
}

func (controller *MerchControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	//merchUpdateRequest.Id = params.ByName("merchId")
	//id, err := strconv.Atoi(merchId)
	//utils.PanicIfError(err)

	id := 1 // placeholder

	controller.MerchService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code: 204,
		Status: "No Content",
	}

	utils.WriteToResponseBody(w, webResponse)
}

// func (controller *MerchControllerImpl) Order(w http.ResponseWriter, r *http.Request) {
// }