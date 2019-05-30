package phonesRepository

import (
	"countryes/common/httpRequest"
	"countryes/infrastructures/mysql"
	"countryes/repositories"
	"database/sql"
	"encoding/json"
)

const phonesRemoteUrl = "http://country.io/phone.json"

type repository struct {
	mysqlClient *sql.DB
}

func New() (*repository, error) {
	mysqlClient, err := mysql.GetInstance()
	if err != nil {
		return nil, err
	}

	return &repository{
		mysqlClient: mysqlClient,
	}, nil
}

func (r *repository) GetFromRemoteAPIAsMap() (map[string]string, error) {
	phoneCodes, err := httpRequest.SendGet(phonesRemoteUrl)
	if err != nil {
		return nil, err
	}

	phoneCodesMap := make(map[string]string)

	err = json.Unmarshal(phoneCodes, &phoneCodesMap)
	if err != nil {
		return nil, err
	}

	return phoneCodesMap, nil
}

func (r *repository) SaveMap(phoneCodes map[string]string) error {
	sqlQuery := "INSERT IGNORE INTO phones (country_code, phone_code) VALUES "
	sqlConditions := ""
	sqlValues := make([]interface{}, 0, len(phoneCodes))

	for countryCode, phoneCode := range phoneCodes {
		sqlConditions += ",(?, ?)"
		sqlValues = append(sqlValues, countryCode, phoneCode)
	}

	_, err := r.mysqlClient.Exec(sqlQuery+sqlConditions[1:], sqlValues...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetPhoneCodeByCountryName(name string) (string, error) {
	var phoneCode string

	err := r.mysqlClient.QueryRow("SELECT phone_code "+
		"FROM phones "+
		"WHERE country_code IN (SELECT country_code "+
		"FROM countryes "+
		"WHERE country_name= ?)", name).Scan(&phoneCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", &repositories.NotFoundError{}
		}
		return "", err
	}

	return phoneCode, nil
}
