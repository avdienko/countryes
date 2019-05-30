package reloadServices

import (
	"countryes/repositories/countryesRepository"
	"countryes/repositories/phonesRepository"
)

type reloadServices struct{}

func New() (*reloadServices, error) {
	return &reloadServices{}, nil
}

func (r *reloadServices) Reload() error {
	err := r.updateCountryes()
	if err != nil {
		return err
	}

	err = r.updatePhones()
	if err != nil {
		return err
	}

	return nil
}

func (r *reloadServices) updatePhones() error {
	phonesRepo, err := phonesRepository.New()
	if err != nil {
		return err
	}

	phoneCodesMap, err := phonesRepo.GetFromRemoteAPIAsMap()
	if err != nil {
		return err
	}

	err = phonesRepo.SaveMap(phoneCodesMap)
	if err != nil {
		return err
	}

	return nil
}

func (r *reloadServices) updateCountryes() error {
	countryesRepo, err := countryesRepository.New()
	if err != nil {
		return err
	}

	countryes, err := countryesRepo.GetFromRemoteAPIAsMap()
	if err != nil {
		return err
	}

	err = countryesRepo.SaveMap(countryes)
	if err != nil {
		return err
	}

	return nil
}
