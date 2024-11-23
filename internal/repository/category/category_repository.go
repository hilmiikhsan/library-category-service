package category

import (
	"context"
	"database/sql"

	"github.com/hilmiikhsan/library-category-service/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CategoryRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
}

func (r *CategoryRepository) InsertNewCategory(ctx context.Context, category *models.Category) error {
	_, err := r.DB.ExecContext(ctx, r.DB.Rebind(queryInsertNewCategory), category.Name, category.Description)
	if err != nil {
		r.Logger.Error("category::InsertNewCategory - failed to insert new category: ", err)
		return err
	}

	return nil
}

func (r *CategoryRepository) FindCategoryByName(ctx context.Context, name string) (*models.Category, error) {
	var res = new(models.Category)

	err := r.DB.GetContext(ctx, res, r.DB.Rebind(queryFindCategoryByName), name)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Error("category::FindCategoryByName - category doesnt exist")
			return res, nil
		}

		r.Logger.Error("category::FindCategoryByName - failed to find category by name: ", err)
		return nil, err
	}

	return res, nil
}
