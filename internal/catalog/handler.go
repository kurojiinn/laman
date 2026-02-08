package catalog

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Handler обрабатывает HTTP запросы для каталога.
type Handler struct {
	catalogService *CatalogService
}

// NewHandler создает новый обработчик каталога.
func NewHandler(catalogService *CatalogService) *Handler {
	return &Handler{
		catalogService: catalogService,
	}
}

// RegisterRoutes регистрирует маршруты каталога.
func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	catalog := router.Group("/catalog")
	{
		catalog.GET("/categories", h.GetCategories)
		catalog.GET("/products", h.GetProducts)
		catalog.GET("/products/:id", h.GetProduct)
	}
}

// GetCategories обрабатывает GET /catalog/categories
func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := h.catalogService.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetProducts обрабатывает GET /catalog/products
func (h *Handler) GetProducts(c *gin.Context) {
	var categoryID *uuid.UUID
	if catIDStr := c.Query("category_id"); catIDStr != "" {
		if catID, err := uuid.Parse(catIDStr); err == nil {
			categoryID = &catID
		}
	}

	availableOnly := c.Query("available_only") == "true"

	products, err := h.catalogService.GetProducts(c.Request.Context(), categoryID, availableOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProduct обрабатывает GET /catalog/products/:id
func (h *Handler) GetProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID товара"})
		return
	}

	product, err := h.catalogService.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}
