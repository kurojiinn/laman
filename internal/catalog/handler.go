package catalog

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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
		catalog.GET("/subcategories", h.GetSubcategories)
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

	var subcategoryID *uuid.UUID
	if subIDStr := c.Query("subcategory_id"); subIDStr != "" {
		if subID, err := uuid.Parse(subIDStr); err == nil {
			subcategoryID = &subID
		}
	}

	var search *string
	if searchStr := c.Query("search"); searchStr != "" {
		search = &searchStr
	}

	availableOnly := c.Query("available_only") == "true"

	products, err := h.catalogService.GetProductsWithFilters(c.Request.Context(), categoryID, subcategoryID, search, availableOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetSubcategories обрабатывает GET /catalog/subcategories
func (h *Handler) GetSubcategories(c *gin.Context) {
	catIDStr := c.Query("category_id")
	if catIDStr == "" {
		c.JSON(http.StatusOK, []interface{}{})
		return
	}

	categoryID, err := uuid.Parse(catIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID категории"})
		return
	}

	subcategories, err := h.catalogService.GetSubcategories(c.Request.Context(), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subcategories)
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
