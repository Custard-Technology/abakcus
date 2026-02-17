package models

import "time"

type Menu struct {
	MenuID      string    `bson:"_id" json:"menu_id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	BusinessID  string    `bson:"business_id" json:"business_id"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
	IsActive    bool      `bson:"is_active" json:"is_active"`
}

type CreateMenuRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateMenuRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsActive    *bool  `json:"is_active,omitempty"`
}

type MenuItem struct {
	ItemID      string    `bson:"_id" json:"item_id"`
	MenuID      string    `bson:"menu_id" json:"menu_id"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	Price       float64   `bson:"price" json:"price"`
	ImageURL    string    `bson:"image_url" json:"image_url"`
	Ingredients []string  `bson:"ingredients" json:"ingredients"`
	IsActive    bool      `bson:"is_active" json:"is_active"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
