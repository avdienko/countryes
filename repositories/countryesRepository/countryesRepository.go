package countryesRepository

import (
	"countryes/common/httpRequest"
	"countryes/infrastructures/mysql"
	"database/sql"
	"encoding/json"
)

const countryesRemoteUrl = "http://country.io/names.json"

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
	countryes, err := httpRequest.SendGet(countryesRemoteUrl)
	if err != nil {
		return nil, err
	}

	countryesMap := make(map[string]string)

	err = json.Unmarshal(countryes, &countryesMap)
	if err != nil {
		return nil, err
	}

	return countryesMap, nil
}

func (r *repository) SaveMap(countryes map[string]string) error {
	sqlQuery := "INSERT IGNORE INTO countryes (country_code, country_name) VALUES "
	sqlConditions := ""
	sqlValues := make([]interface{}, 0, len(countryes))

	for countryCode, countryName := range countryes {
		sqlConditions += ",(?, ?)"
		sqlValues = append(sqlValues, countryCode, countryName)
	}

	_, err := r.mysqlClient.Exec(sqlQuery+sqlConditions[1:], sqlValues...)
	if err != nil {
		return err
	}

	return nil
}
