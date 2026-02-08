package catalog

import (
	"context"
	"database/sql"
	"fmt"
	"Laman/internal/database"
	"Laman/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// postgresCategoryRepository реализует CategoryRepository используя PostgreSQL.
type postgresCategoryRepository struct {
	db *database.DB
}

// NewPostgresCategoryRepository создает новый PostgreSQL репозиторий категорий.
func NewPostgresCategoryRepository(db *database.DB) CategoryRepository {
	return &postgresCategoryRepository{db: db}
}

func (r *postgresCategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	query := `SELECT id, name, description, created_at, updated_at FROM categories ORDER BY name`
	err := r.db.SelectContext(ctx, &categories, query)
	return categories, err
}

func (r *postgresCategoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Category, error) {
	var category models.Category
	query := `SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1`
	err := r.db.GetContext(ctx, &category, query, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("категория не найдена")
	}
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// postgresProductRepository реализует ProductRepository используя PostgreSQL.
type postgresProductRepository struct {
	db *database.DB
}

// NewPostgresProductRepository создает новый PostgreSQL репозиторий товаров.
func NewPostgresProductRepository(db *database.DB) ProductRepository {
	return &postgresProductRepository{db: db}
}

func (r *postgresProductRepository) GetAll(ctx context.Context, categoryID *uuid.UUID, availableOnly bool) ([]models.Product, error) {
	var products []models.Product
	query := `SELECT id, category_id, store_id, name, description, price, weight, is_available, created_at, updated_at FROM products WHERE 1=1`
	args := []interface{}{}
	argPos := 1

	if categoryID != nil {
		query += fmt.Sprintf(" AND category_id = $%d", argPos)
		args = append(args, *categoryID)
		argPos++
	}

	if availableOnly {
		query += fmt.Sprintf(" AND is_available = $%d", argPos)
		args = append(args, true)
		argPos++
	}

	query += " ORDER BY name"

	err := r.db.SelectContext(ctx, &products, query, args...)
	return products, err
}

func (r *postgresProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	var product models.Product
	query := `SELECT id, category_id, store_id, name, description, price, weight, is_available, created_at, updated_at FROM products WHERE id = $1`
	err := r.db.GetContext(ctx, &product, query, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("товар не найден")
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *postgresProductRepository) GetByIDs(ctx context.Context, ids []uuid.UUID) ([]models.Product, error) {
	if len(ids) == 0 {
		return []models.Product{}, nil
	}

	var products []models.Product
	query, args, err := sqlx.In(`SELECT id, category_id, store_id, name, description, price, weight, is_available, created_at, updated_at FROM products WHERE id IN (?)`, ids)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.SelectContext(ctx, &products, query, args...)
	return products, err
}

// postgresStoreRepository реализует StoreRepository используя PostgreSQL.
type postgresStoreRepository struct {
	db *database.DB
}

// NewPostgresStoreRepository создает новый PostgreSQL репозиторий магазинов.
func NewPostgresStoreRepository(db *database.DB) StoreRepository {
	return &postgresStoreRepository{db: db}
}

func (r *postgresStoreRepository) GetAll(ctx context.Context) ([]models.Store, error) {
	var stores []models.Store
	query := `SELECT id, name, address, phone, created_at, updated_at FROM stores ORDER BY name`
	err := r.db.SelectContext(ctx, &stores, query)
	return stores, err
}

func (r *postgresStoreRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Store, error) {
	var store models.Store
	query := `SELECT id, name, address, phone, created_at, updated_at FROM stores WHERE id = $1`
	err := r.db.GetContext(ctx, &store, query, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("магазин не найден")
	}
	if err != nil {
		return nil, err
	}
	return &store, nil
}
