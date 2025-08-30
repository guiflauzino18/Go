package model

import (
	"errors"
	"go-project/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Mail     string    `json:"mail,omitempty"`
	Password string    `json:"password,omitempty"`
	Register time.Time `json:"register,omitempty"`
}

// Validar dados para cadastro do Usuario
func (user *User) Prepare(action string) error {
	if err := user.validate(action); err != nil {
		return err
	}

	if err := user.format(action); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(action string) error {

	if user.Name == "" {
		return errors.New("Nome não pode ser em branco")
	}
	if user.Nick == "" {
		return errors.New("Nick não pode ser em branco")
	}
	if user.Mail == "" {
		return errors.New("Nome não pode ser em branco")
	}

	if err := checkmail.ValidateFormat(user.Mail); err != nil {
		return errors.New("E-mail informado é inválido")
	}

	if action == "register" && user.Password == "" {
		return errors.New("Senha não pode ser em branco")
	}

	return nil
}

// Formata os campos para ser salvo
func (user *User) format(action string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Mail = strings.TrimSpace(user.Mail)

	if action == "register" {
		passWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passWithHash)
	}

	return nil
}
