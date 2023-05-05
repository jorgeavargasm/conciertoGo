package asientoservice

import (
	m "../../models"
	asientoRepository "../../repositories/asiento.repository"
)

func Create(asiento m.Asiento) error {
	err := asientoRepository.Create(asiento)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Asientos, error) {
	Lista, err := asientoRepository.Read()

	if err != nil {
		return nil, err
	}
	return Lista, nil
}

func Update(asiento m.Asiento, idAsiento string) error {
	err := asientoRepository.Update(asiento, idAsiento)
	if err != nil {
		return err
	}
	return nil
}

func Delete(idAsiento string) error {
	err := asientoRepository.Delete(idAsiento)
	if err != nil {
		return err
	}
	return nil
}
