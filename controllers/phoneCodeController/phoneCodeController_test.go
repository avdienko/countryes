package phoneCodeController

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type ApiResponse struct {
	PhoneCode string `json:"phoneCode"`
}

func TestPhoneCodeController(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8090/code/Jordan")
	if err != nil {
		t.Fatal("Error http request: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status code = " + string(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Errorread response body")
	}

	var response ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal("Err unmarshal response json")
	}

	if len(response.PhoneCode) == 0 {
		t.Fatal("Reponse empty")
	}

	// todo ... add some rules ....

}
