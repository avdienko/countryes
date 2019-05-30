package reloadController

import (
	"countryes/common/httpResponse"
	"countryes/common/log"
	"countryes/services/reloadServices"
	"net/http"
)

const (
	INTERNAL_ERROR    = "INTERNAL_ERROR "
	DATA_UPDATE_ERROR = "DATA_UPDATE_ERROR "
)

type responseParams struct {
	Success bool `json:"success"`
}

func ReloadController(w http.ResponseWriter, r *http.Request) {
	reloadServices, err := reloadServices.New()
	if err != nil {
		httpResponse.Error(w, INTERNAL_ERROR+err.Error())
		log.Error(INTERNAL_ERROR + err.Error())
		return
	}

	err = reloadServices.Reload()
	if err != nil {
		httpResponse.Error(w, DATA_UPDATE_ERROR+err.Error())
		log.Error(DATA_UPDATE_ERROR + err.Error())
		return
	}

	httpResponse.Success(w, &responseParams{
		Success: true,
	})
	return
}
