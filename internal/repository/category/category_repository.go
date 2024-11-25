package category

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hilmiikhsan/library-category-service/constants"
	"github.com/hilmiikhsan/library-category-service/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CategoryRepository struct {
	DB     *sqlx.DB
	Logger *logrus.Logger
	Redis  *redis.Client
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
	var (
		res      = new(models.Category)
		cacheKey = fmt.Sprintf("category_by_name:%s", name)
	)

	cachedData, err := r.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), res)
		if err == nil {
			r.Logger.Info("category::FindCategoryByName - Data retrieved from cache")
			return res, nil
		}
		r.Logger.Warn("category::FindCategoryByName - Failed to unmarshal cache data: ", err)
	}

	err = r.DB.GetContext(ctx, res, r.DB.Rebind(queryFindCategoryByName), name)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Error("category::FindCategoryByName - category doesnt exist")
			return res, nil
		}

		r.Logger.Error("category::FindCategoryByName - failed to find category by name: ", err)
		return nil, err
	}

	dataToCache, err := json.Marshal(res)
	if err != nil {
		r.Logger.Warn("category::FindCategoryByName - Failed to marshal data for caching: ", err)
	} else {
		err = r.Redis.Set(ctx, cacheKey, dataToCache, 5*time.Minute).Err()
		if err != nil {
			r.Logger.Warn("category::FindCategoryByName - Failed to cache data: ", err)
		}
	}

	return res, nil
}

func (r *CategoryRepository) FindCategoryByID(ctx context.Context, id string) (*models.Category, error) {
	var (
		res = new(models.Category)
	)

	err := r.DB.GetContext(ctx, res, r.DB.Rebind(queryFindCategoryByID), id)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Logger.Error("category::FindCategoryByID - category doesnt exist")
			return res, errors.New(constants.ErrCategoryNotFound)
		}

		r.Logger.Error("category::FindCategoryByID - failed to find category by id: ", err)
		return nil, err
	}

	return res, nil
}

func (r *CategoryRepository) FindAllCategory(ctx context.Context, limit, offset int) ([]models.Category, error) {
	var (
		res      = make([]models.Category, 0)
		cacheKey = fmt.Sprintf("categories:limit:%d:offset:%d", limit, offset)
	)

	cachedData, err := r.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &res)
		if err == nil {
			r.Logger.Info("category::FindAllCategory - Data retrieved from cache")
			return res, nil
		}
		r.Logger.Warn("category::FindAllCategory - Failed to unmarshal cache data: ", err)
	}

	err = r.DB.SelectContext(ctx, &res, r.DB.Rebind(queryFindAllCategory), limit, offset)
	if err != nil {
		r.Logger.Error("category::FindAllCategory - failed to find all category: ", err)
		return nil, err
	}

	dataToCache, err := json.Marshal(res)
	if err != nil {
		r.Logger.Warn("category::FindAllCategory - Failed to marshal data for caching: ", err)
	} else {
		err = r.Redis.Set(ctx, cacheKey, dataToCache, 5*time.Minute).Err()
		if err != nil {
			r.Logger.Warn("category::FindAllCategory - Failed to cache data: ", err)
		}
	}

	return res, nil
}

func (r *CategoryRepository) UpdateNewCategory(ctx context.Context, category *models.Category) error {
	_, err := r.DB.ExecContext(ctx, r.DB.Rebind(queryUpdateNewCategory), category.Name, category.Description, category.ID)
	if err != nil {
		r.Logger.Error("category::UpdateNewCategory - failed to update new category: ", err)
		return err
	}

	return nil
}

func (r *CategoryRepository) DeleteCategoryByID(ctx context.Context, id string) error {
	_, err := r.DB.ExecContext(ctx, r.DB.Rebind(queryDeleteCategoryByID), id)
	if err != nil {
		r.Logger.Error("category::DeleteCategoryByID - failed to delete category by id: ", err)
		return err
	}

	return nil
}
