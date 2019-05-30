package phoneCodeController

import (
	"net/http"

	"countryes/common/httpResponse"
	"countryes/common/log"
	"countryes/services/phoneCodesServices"

	"github.com/gorilla/mux"
)

const (
	INTERNAL_ERROR       = "INTERNAL_ERROR "
	DATA_NOT_FOUND_ERROR = "DATA_NOT_FOUND_ERROR "
)

type responseParams struct {
	PhoneCode string `json:"phoneCode"`
}

func PhoneCodeController(w http.ResponseWriter, r *http.Request) {
	inputParams := mux.Vars(r)
	countryName := inputParams["name"] // TODO add validate input param

	log.Info(r.URL.Path)

	services, err := phoneCodesServices.New()
	if err != nil {
		httpResponse.Error(w, INTERNAL_ERROR+err.Error())
		log.Error(INTERNAL_ERROR + err.Error())
		return
	}

	phoneCode, err := services.GetPhoneCodeByCountryName(countryName)
	if err != nil {
		if err == phoneCodesServices.NotFoundError {
			httpResponse.ErrorNotFound(w, DATA_NOT_FOUND_ERROR)
			log.Error(DATA_NOT_FOUND_ERROR + err.Error())
			return
		}

		httpResponse.Error(w, INTERNAL_ERROR+err.Error())
		log.Error(DATA_NOT_FOUND_ERROR + err.Error())
		return
	}

	httpResponse.Success(w, responseParams{
		PhoneCode: phoneCode,
	})
	return
}
