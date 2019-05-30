package phoneCodesServices

import (
	"countryes/repositories"
	"countryes/repositories/phonesRepository"
)

type phoneCodesServices struct{}

func New() (*phoneCodesServices, error) {
	return &phoneCodesServices{}, nil
}

func (r *phoneCodesServices) GetPhoneCodeByCountryName(country string) (string, error) {
	phonesRepository, err := phonesRepository.New()
	if err != nil {
		return "", err
	}

	phoneCode, err := phonesRepository.GetPhoneCodeByCountryName(country)
	if err != nil {
		if _, ok := err.(*repositories.NotFoundError); ok {
			return "", NotFoundError
		}
		return "", err
	}
	return phoneCode, nil
}
