package repository

import (
	"database/sql"
	"fmt"
	"go-project/model"
)

type UserRepo struct {
	db *sql.DB
}

// Cria novo repositório de usuários para operações de CRUD
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// Create cria um novo usuário no Banco
func (repo UserRepo) Create(user model.User) (uint64, error) {
	//Preparar query - Camada de segurança que evita ataque de SQL Iject
	statement, err := repo.db.Prepare("insert into users (name, username, mail,password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close() //Fechar Statement ao finalizar

	result, err := statement.Exec(user.Name, user.Username, user.Mail, user.Password)
	if err != nil {
		return 0, err
	}

	//Get lastId
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil

}

// Find Al Users
func (repo UserRepo) FindByUsernameOrName(UsernameOrName string) ([]model.User, error) {
	//Formatar UsernameOuName para conter os % no incio e fim
	UsernameOrName = fmt.Sprintf("%%%s%%", UsernameOrName)

	row, err := repo.db.Query("select id, name, username, mail, register from users where name LIKE ? or Username LIKE ?", UsernameOrName)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var users []model.User
	for row.Next() {
		var user model.User
		if erro := row.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Mail,
			&user.Register,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// Find By Mail
func (repo UserRepo) FindByMail(mail string) (model.User, error) {
	row, err := repo.db.Query("select id,username, password from users where mail = ?", mail)
	if err != nil {
		return model.User{}, err
	}
	defer row.Close()

	var user model.User
	if row.Next() {
		if err := row.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
		); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

// Find By ID
func (repo UserRepo) FindByID(ID uint64) (model.User, error) {
	row, err := repo.db.Query("select id, name, username, mail, register from users where id = ?", ID)
	if err != nil {
		return model.User{}, err
	}
	defer row.Close()

	var user model.User
	if row.Next() {
		if err := row.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Mail,
			&user.Register,
		); err != nil {
			return model.User{}, err
		}

		return user, nil
	}

	return model.User{}, fmt.Errorf("Não encontrado")

}

// Find All Users
func (repo UserRepo) FindAll() ([]model.User, error) {
	row, err := repo.db.Query("select id, name, username, register from users")
	if err != nil {
		return []model.User{}, err
	}
	defer row.Close()

	var users []model.User
	for row.Next() {
		var user model.User
		if err := row.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Register,
		); err != nil {
			return []model.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Update an User
func (repo UserRepo) Update(ID uint64, user model.User) error {
	statement, err := repo.db.Prepare("update users set name = ?, username = ?, mail = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Username, user.Mail, ID)
	if err != nil {
		return err
	}

	return nil

}

// Delete an User
func (repo UserRepo) Delete(ID uint64) error {
	statement, err := repo.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
