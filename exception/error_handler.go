package exception

import (
	"e-commerce-api/model/web"
	"e-commerce-api/utils"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func notFoundError(writer http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
	
		webResponse := web.WebResponse{
			Code: http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data: exception.Error,
		}

		utils.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, req *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}

	utils.WriteToResponseBody(writer, webResponse)
}
