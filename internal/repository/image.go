package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/restapi/domain"
)

type imageRepository struct {
	db *goqu.Database
}

func NewImage(db *sql.DB) domain.ImageRepository {
	return &imageRepository{
		db: goqu.New("postgres", db),
	}
}

func (i imageRepository) All(ctx context.Context) (images []domain.Image, err error) {
	dataset := i.db.From("images").Order(goqu.I("id").Asc())
	err = dataset.ScanStructsContext(ctx, &images)
	return
}

func (i imageRepository) Delete(ctx context.Context, id int64) error {
	executor := i.db.Delete("images").Where(goqu.Ex{
		"id": id,
	}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (i imageRepository) FindByUserId(ctx context.Context, UserId int64) (image domain.Image, err error) {
	dataset := i.db.From("images").Where(goqu.Ex{
		"users_id": UserId,
	})

	_, err = dataset.ScanStructContext(ctx, &image)
	return
}

func (i imageRepository) Insert(ctx context.Context, image *domain.Image) error {
	executor := i.db.Insert("images").Rows(goqu.Ex{
		"users_id": image.UserId,
		"images":   image.Images,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, image)
	return err
}
