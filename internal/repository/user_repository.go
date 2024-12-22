package repository

import (
	"awesomeProject/internal/model"
	"database/sql"
	"errors"
)

type IUserRepository interface {
	GetUserByUsername(username string) (*model.User, error)
	SaveUser(u *model.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	row := r.db.QueryRow("select ut.id,ut.username,ut.password  from user_tbl ut where username = $1", username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}
	roleRow, err := r.db.Query("select role_tbl.role_name "+
		"from user_role_tbl inner join role_tbl on user_role_tbl.role_id = role_tbl.id where user_role_tbl.user_id = $1 ", user.ID)
	if err != nil {
		return nil, err
	}

	for roleRow.Next() {
		var role string
		err := roleRow.Scan(&role)
		if err != nil {
			return nil, err
		}
		user.Roles = append(user.Roles, role)
	}
	return &user, nil
}

// SaveUser func to save user under database @return err if you can't save to database, if it not return nil error
func (r *UserRepository) SaveUser(u *model.User) error {
	trans, err := r.db.Begin()
	if err != nil {
		return err
	}
	var userID int

	err = trans.QueryRow("insert into user_tbl(username, password) values ($1, $2) RETURNING id", u.Username, u.Password).Scan(&userID)

	if err != nil {
		trans.Rollback()
		return err
	}

	_, err = trans.Exec("insert into user_role_tbl(user_id, role_id) values ($1, $2)", userID, 3)
	if err != nil {
		err := trans.Rollback()
		if err != nil {
			return err
		}
	}
	if err = trans.Commit(); err != nil {
		return err
	}
	return nil
}
