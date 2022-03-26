package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Paciente struct {
	gorm.Model
	Nome  string `json:"nome" validate:"nonzero"`
	CPF   string `json:"cpf" validate:"min=11,max=11,regexp=^[0-9]*$"`
	senha string `validate:"min=6"`
}

func ValidadorPaciente(paciente *Paciente) error {
	if err := validator.Validate(paciente); err != nil {
		return err
	}
	return nil
}
