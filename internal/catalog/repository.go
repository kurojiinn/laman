package catalog

import (
	"context"
	"Laman/internal/models"
	"github.com/google/uuid"
)

// CategoryRepository определяет интерфейс для доступа к данным категорий.
type CategoryRepository interface {
	// GetAll получает все категории.
	GetAll(ctx context.Context) ([]models.Category, error)
	
	// GetByID получает категорию по ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Category, error)
}

// ProductRepository определяет интерфейс для доступа к данным товаров.
type ProductRepository interface {
	// GetAll получает все товары с опциональными фильтрами.
	GetAll(ctx context.Context, categoryID *uuid.UUID, availableOnly bool) ([]models.Product, error)
	
	// GetByID получает товар по ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error)
	
	// GetByIDs получает несколько товаров по их ID.
	GetByIDs(ctx context.Context, ids []uuid.UUID) ([]models.Product, error)
}

// StoreRepository определяет интерфейс для доступа к данным магазинов.
type StoreRepository interface {
	// GetAll получает все магазины.
	GetAll(ctx context.Context) ([]models.Store, error)
	
	// GetByID получает магазин по ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Store, error)
}
