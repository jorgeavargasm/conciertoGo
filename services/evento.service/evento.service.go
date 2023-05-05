package evento_service

import (
	m "../../models"
	EventoRepository "../../repositories/evento.repository"
)

func Create(evento m.Evento) error {
	err := EventoRepository.Create(evento)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Eventos, error) {
	Lista, err := EventoRepository.Read()

	if err != nil {
		return nil, err
	}
	return Lista, nil
}

func Update(evento m.Evento, idEvento string) error {
	err := EventoRepository.Update(evento, idEvento)
	if err != nil {
		return err
	}
	return nil
}

func Delete(idEvento string) error {
	err := EventoRepository.Delete(idEvento)
	if err != nil {
		return err
	}
	return nil
}
