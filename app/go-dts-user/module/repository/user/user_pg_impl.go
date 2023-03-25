package user

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Calmantara/go-dts-user/module/model"
)

type UserPgRepoImpl struct {
	db *sql.DB
}

func NewUserPgRepo(db *sql.DB) UserRepo {
	return &UserPgRepoImpl{
		db: db,
	}
}

func (u *UserPgRepoImpl) FindUserById(ctx context.Context, userId uint64) (user model.User, err error) {
	query := `
		SELECT 
			id, 
			email, 
			name,
			dob
		FROM users u
		WHERE u.id = $1
			AND deleted_at is null;
	`
	// $1 -> sql argument / variable inputan
	// prepare untuk checking error di query
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			err = errors.New("error duplication email")
		}
		return
	}
	defer rows.Close()

	for rows.Next() {
		// SCAN -> binding data ke golang struct
		if err = rows.Scan(&user.Id, &user.Email, &user.Name, &user.Dob); err != nil {
			return
		}
	}

	if user.Id <= 0 {
		err = errors.New("user is not found")
	}

	return
}

func (u *UserPgRepoImpl) FindAllUsers(ctx context.Context) (users []model.User, err error) {
	query := `
		SELECT 
			id, 
			email, 
			name,
			dob
		FROM users u
		WHERE deleted_at is null
		ORDER BY created_at ASC;
	`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.Id, &user.Email, &user.Name, &user.Dob); err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func (u *UserPgRepoImpl) InsertUser(ctx context.Context, userIn model.User) (user model.User, err error) {
	query := `
		INSERT INTO users
				(name, email, dob)
		VALUES
			($1, $2, $3)
		RETURNING 
			id,
			email,
			name,
			dob;
	`
	// RETURNING akan mengeluarkan
	// affected rows dari hasil query kita
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx,
		userIn.Name,
		userIn.Email,
		userIn.Dob)
	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Name, &user.Dob); err != nil {
			return
		}
	}
	return
}

func (u *UserPgRepoImpl) UpdateUser(ctx context.Context, userIn model.User) (err error) {
	query := `
		UPDATE users
		SET
			name  = $1,
			email = $4,
			dob   = $2
		WHERE id = $3
			AND deleted_at is null;
	`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		userIn.Name,
		userIn.Dob,
		userIn.Id,
		userIn.Email)
	if err != nil {
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affected <= 0 {
		err = errors.New("user is not found")
		return
	}
	return
}

func (u *UserPgRepoImpl) DeleteUserById(ctx context.Context, userId uint64) (user model.User, err error) {
	query := `
		UPDATE users
		SET
			deleted_at = now()
		WHERE id = $1 
			AND deleted_at is null
		RETURNING 
			id,
			email,
			name,
			dob;
	`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx,
		userId)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Name, &user.Dob); err != nil {
			return
		}
	}
	if user.Id <= 0 {
		err = errors.New("user is not found")
	}
	return
}
