package httpResponse

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, data interface{}) error {
	if response, err := json.Marshal(data); err != nil {
		return err
	} else {
		w.Header().Add("Content-type", "application/json")
		w.Write(response)
	}

	return nil
}

func Error(w http.ResponseWriter, err string) error {
	errResponse := map[string]string{
		"msg": err,
	}

	if response, err := json.Marshal(errResponse); err != nil {
		return err
	} else {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	}

	return nil
}

func ErrorNotFound(w http.ResponseWriter, err string) error {
	errResponse := map[string]string{
		"msg": err,
	}

	if response, err := json.Marshal(errResponse); err != nil {
		return err
	} else {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
	}

	return nil
}
