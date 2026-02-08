package catalog

import (
	"context"
	"fmt"
	"Laman/internal/models"
	"github.com/google/uuid"
)

// CatalogService обрабатывает бизнес-логику, связанную с каталогом,
// включая категории, товары и магазины.
type CatalogService struct {
	categoryRepo CategoryRepository
	productRepo  ProductRepository
	storeRepo    StoreRepository
}

// NewCatalogService создает новый сервис каталога.
func NewCatalogService(
	categoryRepo CategoryRepository,
	productRepo ProductRepository,
	storeRepo StoreRepository,
) *CatalogService {
	return &CatalogService{
		categoryRepo: categoryRepo,
		productRepo:  productRepo,
		storeRepo:    storeRepo,
	}
}

// GetCategories получает все категории.
func (s *CatalogService) GetCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := s.categoryRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить категории: %w", err)
	}
	return categories, nil
}

// GetProducts получает товары с опциональными фильтрами.
func (s *CatalogService) GetProducts(ctx context.Context, categoryID *uuid.UUID, availableOnly bool) ([]models.Product, error) {
	products, err := s.productRepo.GetAll(ctx, categoryID, availableOnly)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить товары: %w", err)
	}
	return products, nil
}

// GetProduct получает товар по ID.
func (s *CatalogService) GetProduct(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	product, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить товар: %w", err)
	}
	return product, nil
}

// GetStores получает все магазины.
func (s *CatalogService) GetStores(ctx context.Context) ([]models.Store, error) {
	stores, err := s.storeRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить магазины: %w", err)
	}
	return stores, nil
}
