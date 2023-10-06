package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/restapi/domain"
	"github.com/khairulharu/restapi/dto"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		goqu.New("postgres", con),
	}
}

func (u userRepository) Insert(ctx context.Context, user *domain.User) error {
	passGen, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	executor := u.db.Insert("users").Rows(goqu.Record{
		"full_name": user.FullName,
		"phone":     user.Phone,
		"username":  user.Username,
		"password":  string(passGen),
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, user)
	return err
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"username": username,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) Update(ctx context.Context, user *domain.User) (userResult dto.UserData, err error) {
	passGen, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	dataset := u.db.Update("users").Where(goqu.Ex{
		"id": user.ID,
	}).Set(goqu.Record{
		"full_name": user.FullName,
		"phone":     user.Phone,
		"username":  user.Username,
		"password":  string(passGen),
	}).Executor()

	_, err = dataset.ScanStructContext(ctx, &userResult)
	return
}

func (u userRepository) FindAll(ctx context.Context) (user []domain.User, err error) {
	dataset := u.db.From("users").Order(goqu.I("id").Asc())
	err = dataset.ScanStructsContext(ctx, &user)
	return
}

func (u userRepository) Delete(ctx context.Context, id int64) error {
	executor := u.db.Delete("users").Where(goqu.Ex{
		"id": id,
	}).Executor()

	_, err := executor.ExecContext(ctx)
	return err
}
