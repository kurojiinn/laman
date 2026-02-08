package models

import (
	"time"

	"github.com/google/uuid"
)

// Category представляет категорию товара.
type Category struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description *string   `db:"description" json:"description,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// Product представляет товар в каталоге.
type Product struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	CategoryID  uuid.UUID  `db:"category_id" json:"category_id"`
	StoreID     *uuid.UUID `db:"store_id" json:"store_id,omitempty"`
	Name        string     `db:"name" json:"name"`
	Description *string    `db:"description" json:"description,omitempty"`
	Price       float64    `db:"price" json:"price"`
	Weight      *float64   `db:"weight" json:"weight,omitempty"`
	IsAvailable bool       `db:"is_available" json:"is_available"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
}

// Store представляет магазин, где можно приобрести товары.
type Store struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Address     string    `db:"address" json:"address"`
	Phone       *string   `db:"phone" json:"phone,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
