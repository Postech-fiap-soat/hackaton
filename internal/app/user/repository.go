package user

import (
	"database/sql"
	"hackaton/internal/app/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (u *UserRepository) GetUserById(userID int) (*domain.User, error) {
	stmt, err := u.db.Prepare("select id, name, registration, email from users where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}
	var user domain.User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Registration, &user.Email)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (u *UserRepository) GetUserByRegistration(registration string) (*domain.User, error) {
	stmt, err := u.db.Prepare("select id, name, registration, email, password from users where registration = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(registration)
	if err != nil {
		return nil, err
	}
	var user domain.User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Registration, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
